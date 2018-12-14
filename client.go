// MIT License
//
// Copyright (c) 2018 Markus Tenghamn
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package nordeago is an abstraction for the Nordea Open Banking API in Go (https://developer.nordeaopenbanking.com/app/docs).
package nordeago

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"
)

// Client holds all the needed information to communicate with the nordea API. Use InitClient to create a new Client.
type Client struct {
	BaseURL      string
	Protocol     string
	Version      string
	ClientID     string // X-IBM-Client-Id sent as header
	ClientSecret string // X-IBM-Client-Secret sent as header
	TppToken     string
	RedirectURL  string
	AuthCode     string
	AccessToken  string
}

// InitClient creates a new client from a clientID and clientSecret. You can find this information by signing up for free
// at https://developer.nordeaopenbanking.com
func InitClient(clientID string, clientSecret string, redirectURL string) Client {
	c := Client{}
	c.Protocol = "https://"
	c.BaseURL = "api.nordeaopenbanking.com"
	c.Version = "v2"
	c.ClientID = clientID
	c.ClientSecret = clientSecret
	c.RedirectURL = redirectURL
	return c
}

// request handles requests to the Nordea API
func (c *Client) request(requestType string, endpoint string, body io.Reader, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(requestType, c.GetFullURL(endpoint), body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	httpClient := &http.Client{}

	return httpClient.Do(req)
}

// Post handles post requests by converting request types to json and passing the data to Request
func (c *Client) Post(endpoint string, requestObject interface{}, headers map[string]string) (*http.Response, error) {
	requestByte, err := json.Marshal(requestObject)

	if err != nil {
		return nil, err
	}

	headers = initHeaders(headers)

	headers["Accept"] = "application/json"

	requestReader := bytes.NewReader(requestByte)

	return c.request("POST", endpoint, requestReader, headers)
}

// PostWithAccessToken handles post requests by converting request types to json and passing the data to Request and also setting the relevant headers to authenticate with an access token
func (c *Client) PostWithAccessToken(endpoint string, requestObject interface{}, headers map[string]string) (*http.Response, error) {
	requestByte, err := json.Marshal(requestObject)

	if err != nil {
		return nil, err
	}

	headers = c.setAccessTokenHeaders(headers)

	headers["Accept"] = "application/json"

	requestReader := bytes.NewReader(requestByte)

	return c.request("POST", endpoint, requestReader, headers)
}

// Get handles get requests by setting the type of request to GET along with a nil body and calling Request
func (c *Client) Get(endpoint string, headers map[string]string) (*http.Response, error) {
	headers = initHeaders(headers)

	headers["Accept"] = "application/json"

	return c.request("GET", endpoint, nil, headers)
}

// GetWithAccessToken handles get requests by setting the type of request to GET along with a nil body and calling Request with the needed headers
func (c *Client) GetWithAccessToken(endpoint string, headers map[string]string) (*http.Response, error) {
	headers = c.setAccessTokenHeaders(headers)
	return c.request("GET", endpoint, nil, headers)
}

// Put handles put requests by setting the type of request to PUT along with a nil body and calling Request
func (c *Client) Put(endpoint string, headers map[string]string) (*http.Response, error) {
	headers = initHeaders(headers)
	return c.request("PUT", endpoint, nil, headers)
}

// PutWithAccessToken handles put requests by setting the type of request to PUT along with a nil body and calling Request with the needed headers
func (c *Client) PutWithAccessToken(endpoint string, headers map[string]string) (*http.Response, error) {
	headers = c.setAccessTokenHeaders(headers)
	return c.request("PUT", endpoint, nil, headers)
}

// Delete handles delete requests by setting the type of request to DELETE along with a nil body and calling Request
func (c *Client) Delete(endpoint string, headers map[string]string) (*http.Response, error) {
	headers = initHeaders(headers)

	headers["Accept"] = "application/json"

	return c.request("DELETE", endpoint, nil, headers)
}

// DeleteWithAccessToken handles delete requests by setting the type of request to DELETE along with a nil body and calling Request with the needed headers
func (c *Client) DeleteWithAccessToken(endpoint string, headers map[string]string) (*http.Response, error) {
	headers = c.setAccessTokenHeaders(headers)

	headers["Accept"] = "application/json"

	return c.request("DELETE", endpoint, nil, headers)
}

// GetFullURL builds the endpoint url from the client configuration combining it with the supplied endpoint
func (c *Client) GetFullURL(endpoint string) string {
	return c.Protocol + path.Join(c.BaseURL, c.Version, endpoint)
}

func (c *Client) setAccessTokenHeaders(headers map[string]string) map[string]string {
	headers = initHeaders(headers)

	headers["Authorization"] = BearerAuthHeader(c.AccessToken)
	headers["X-IBM-Client-Id"] = c.ClientID
	headers["X-IBM-Client-Secret"] = c.ClientSecret

	return headers
}

func initHeaders(headers map[string]string) map[string]string {
	if headers == nil {
		headers = make(map[string]string)
	}
	// Assumes all requests need the Content-Type header
	headers["Content-Type"] = "application/json"
	return headers
}

// HandleResponse takes a http response and unmarshals the json content if possible or otherwise returns a status code and/or error
func (c *Client) HandleResponse(response *http.Response, result *Result) (int, error) {
	decoder := json.NewDecoder(response.Body)

	// TODO check APIm-Debug-Trans-Id, X-RateLimit-Limit, X-RateLimit-Remaining, X-RateLimit-Reset, X-Global-Transaction-ID headers

	if response.StatusCode == http.StatusOK {
		err := decoder.Decode(&result)

		if err != nil {
			return response.StatusCode, err
		}
	} else {
		return response.StatusCode, errorFromResponse(decoder, result)
	}

	return response.StatusCode, nil
}

func errorFromResponse(decoder *json.Decoder, result *Result) error {
	errorString := ""
	err := decoder.Decode(&result)
	if err == nil {
		if len(result.Error.Failures) > 0 {
			errorString = fmt.Sprintf("%d - %s: %s", result.Error.HTTPCode, result.Error.HTTPMessage, result.Error.MoreInformation)
			for _, failure := range result.Error.Failures {
				errorString += fmt.Sprintf("\n%s: %s", failure.Code, failure.Description)
			}
			return errors.New(errorString)
		}
	}

	var errorResponse ErrorResponse
	err2 := decoder.Decode(&errorResponse)

	if err2 == nil {
		if len(errorResponse.HTTPMessage) > 0 && len(errorResponse.MoreInformation) > 0 {
			return fmt.Errorf("%d - %s: %s", errorResponse.HTTPCode, errorResponse.HTTPMessage, errorResponse.MoreInformation)

		}
	}
	return nil

}
