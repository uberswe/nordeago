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

package ina

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/markustenghamn/nordeago"
	"net/http"
)

// StartAuthDecoupled initiates authentication with the Nordea API which will allow us to make requests to the Accounts and Payments API.
// This initiates the user authentication.
// Warning: Decoupled Authorisation flow is a mock version, and it is only intended to show how the production version will work.
//
// API Documentation: https://developer.nordeaopenbanking.com/app/documentation?api=Identity%20and%20Access%20API&version=2.1#authorize
func StartAuthDecoupled(c *nordeago.Client, request AuthRequestDecoupled) (*Response, error) {
	responseType := &Response{}
	result := nordeago.Result{Response: responseType}
	endpoint := "/authorize-decoupled"

	headers := make(map[string]string)
	headers["X-IBM-Client-Id"] = c.ClientID
	headers["X-IBM-Client-Secret"] = c.ClientSecret

	response, err := c.Post(endpoint, request, headers)

	if err != nil {
		return responseType, err
	}

	defer response.Body.Close()

	_, err = c.HandleResponse(response, &result)

	if err == nil {
		c.TppToken = responseType.TppToken
	}

	return responseType, err
}

// PollForAuthCodeDecoupled polls for an auth code which will be returned when the user has accepted access to
// their accounts/payments based on scope.
// Warning: Decoupled Authorisation flow is a mock version, and it is only intended to show how the production version will work.
//
// API Documentation: https://developer.nordeaopenbanking.com/app/documentation?api=Identity%20and%20Access%20API&version=2.1#getToken
func PollForAuthCodeDecoupled(c *nordeago.Client, orderRef string) (*Response, int, error) {
	responseType := &Response{}
	result := nordeago.Result{Response: responseType}

	endpoint := "/authorize-decoupled/{{order_ref}}"
	endpoint = nordeago.ReplaceVariable(endpoint, "order_ref", orderRef)

	headers := make(map[string]string)
	headers["Authorization"] = nordeago.BearerAuthHeader(c.TppToken)
	headers["X-IBM-Client-Id"] = c.ClientID
	headers["X-IBM-Client-Secret"] = c.ClientSecret

	response, err := c.Get(endpoint, headers)

	if err != nil {
		return responseType, response.StatusCode, err
	}

	defer response.Body.Close()

	statusCode, err := c.HandleResponse(response, &result)

	if err == nil {
		c.AuthCode = responseType.Code
	}

	return responseType, statusCode, err
}

// RetrieveAccessTokenDecoupled returns a bearer token to use for the Accounts and Payments API requests.
// Warning: Decoupled Authorisation flow is a mock version, and it is only intended to show how the production version will work.
//
// API Documentation: https://developer.nordeaopenbanking.com/app/documentation?api=Identity%20and%20Access%20API&version=2.1#getToken
func RetrieveAccessTokenDecoupled(c *nordeago.Client, request RetrieveAccessTokenRequest) (RetrieveAccessTokenResponse, error) {
	var retrieveAccessTokenResponse RetrieveAccessTokenResponse
	result := nordeago.Result{Response: &retrieveAccessTokenResponse}

	endpoint := "/authorize-decoupled/token"

	headers := make(map[string]string)
	headers["Authorization"] = nordeago.BearerAuthHeader(c.TppToken)
	headers["X-IBM-Client-Id"] = c.ClientID
	headers["X-IBM-Client-Secret"] = c.ClientSecret

	response, err := c.Post(endpoint, request, headers)

	if err != nil {
		return retrieveAccessTokenResponse, err
	}

	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)

	if response.StatusCode == http.StatusOK {
		err = decoder.Decode(&retrieveAccessTokenResponse)

		if err != nil {
			return retrieveAccessTokenResponse, err
		}

		c.AccessToken = retrieveAccessTokenResponse.AccessToken
	} else {
		_, err = c.HandleResponse(response, &result)
		return retrieveAccessTokenResponse, err
	}

	return retrieveAccessTokenResponse, nil
}

// StartAuth returns a url to redirect the user to for Oauth flow?
//
// TODO I am not sure how this differs as I do not have access to a production environment and can't test
//
// API Documentation: https://developer.nordeaopenbanking.com/app/documentation?api=Identity%20and%20Access%20API&version=2.1#startAuthentication
func StartAuth(c *nordeago.Client, request AuthRequest) (string, error) {
	endpoint := "/authorize"

	req, err := http.NewRequest("GET", c.GetFullURL(endpoint), nil)
	if err != nil {
		return endpoint, err
	}

	v, _ := query.Values(request)

	req.URL.RawQuery = v.Encode()

	return req.URL.String(), nil
}

// RetrieveAccessToken returns a bearer token to use for the Accounts and Payments API requests.
//
// TODO I am not sure how this differs as I do not have access to a production environment and can't test
//
// API Documentation: https://developer.nordeaopenbanking.com/app/documentation?api=Identity%20and%20Access%20API&version=2.1#accessToken
func RetrieveAccessToken(c *nordeago.Client, request RetrieveAccessTokenRequest) (RetrieveAccessTokenResponse, error) {
	var retrieveAccessTokenResponse RetrieveAccessTokenResponse

	endpoint := "/authorize/access_token"

	headers := make(map[string]string)
	headers["Authorization"] = nordeago.BearerAuthHeader(c.TppToken)
	headers["X-IBM-Client-Id"] = c.ClientID
	headers["X-IBM-Client-Secret"] = c.ClientSecret

	response, err := c.Post(endpoint, request, headers)

	if err != nil {
		return retrieveAccessTokenResponse, err
	}

	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)

	if response.StatusCode == http.StatusOK {
		err = decoder.Decode(&retrieveAccessTokenResponse)

		if err != nil {
			return retrieveAccessTokenResponse, err
		}

		c.AccessToken = retrieveAccessTokenResponse.AccessToken
	} else {
		var errorResponse nordeago.ErrorResponse
		err = decoder.Decode(&errorResponse)

		if err != nil {
			return retrieveAccessTokenResponse, err
		}

		err := fmt.Errorf("%d - %s: %s", errorResponse.HTTPCode, errorResponse.HTTPMessage, errorResponse.MoreInformation)

		return retrieveAccessTokenResponse, err
	}

	return retrieveAccessTokenResponse, nil
}

// GetAssets use an access token to get the assets or accounts of the authenticated user
//
// API Documentation: https://developer.nordeaopenbanking.com/app/documentation?api=Identity%20and%20Access%20API&version=2.1#getAssets
func GetAssets(c *nordeago.Client) (*Response, error) {
	responseType := &Response{}
	result := nordeago.Result{Response: responseType}

	endpoint := "/assets"

	response, err := c.GetWithAccessToken(endpoint, nil)

	if err != nil {
		return responseType, err
	}

	defer response.Body.Close()

	_, err = c.HandleResponse(response, &result)

	return responseType, err
}
