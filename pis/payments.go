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

package pis

import (
	"errors"
	"github.com/markustenghamn/nordeago"
	"net/http"
)

// GetPayments returns a nordeago.Result with pis.PaymentsResponse as the response
func GetPayments(c *nordeago.Client, country string) (*PaymentsResponse, error) {
	responseType := &PaymentsResponse{}
	result := nordeago.Result{Response: responseType}

	endpoint := getEndpointFromCountry(country)

	response, err := c.GetWithAccessToken(endpoint, nil)

	if err != nil {
		return responseType, err
	}

	defer response.Body.Close()

	_, err = c.HandleResponse(response, &result)

	return responseType, err
}

// InitiatePayment sends an InitiatePaymentRequest and returns true if the API responds with a 201 created status code
func InitiatePayment(c *nordeago.Client, country string, request InitiatePaymentRequest, skipAccessControl bool) (bool, error) {
	endpoint := getEndpointFromCountry(country)

	var headers map[string]string
	// X-Response-Scenarios header can be set to AuthorizationSkipAccessControl in sandbox environments
	if skipAccessControl {
		headers = make(map[string]string)
		headers["X-Response-Scenarios"] = "AuthorizationSkipAccessControl"
	}

	response, err := c.PostWithAccessToken(endpoint, request, headers)

	if err != nil {
		return false, err
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusCreated {
		return true, nil
	}

	err = errors.New("Request returned status " + string(response.StatusCode))
	return false, err
}

// GetPayment returns a nordeago.Result with pis.Payment as the response
func GetPayment(c *nordeago.Client, country string, paymentID string, skipAccessControl bool) (*Payment, error) {
	responseType := &Payment{}
	result := nordeago.Result{Response: responseType}

	endpoint := nordeago.ReplaceVariable(getEndpointFromCountry(country)+"/{{paymentId}}", "paymentId", paymentID)

	var headers map[string]string
	// X-Response-Scenarios header can be set to AuthorizationSkipAccessControl in sandbox environments
	if skipAccessControl {
		headers = make(map[string]string)
		headers["X-Response-Scenarios"] = "AuthorizationSkipAccessControl"
	}

	response, err := c.GetWithAccessToken(endpoint, headers)

	if err != nil {
		return responseType, err
	}

	defer response.Body.Close()

	_, err = c.HandleResponse(response, &result)

	return responseType, err
}

// ConfirmPayment returns a nordeago.Result with pis.Payment as the response
func ConfirmPayment(c *nordeago.Client, country string, paymentID string, responseScenario string) (*Payment, error) {
	responseType := &Payment{}
	result := nordeago.Result{Response: responseType}

	endpoint := nordeago.ReplaceVariable(getEndpointFromCountry(country)+"/{{paymentId}}/confirm", "paymentId", paymentID)

	// X-Response-Scenarios header can be set to AuthorizationSkipAccessControl, PaymentSigningExpires, PaymentMissingFunds or PaymentOnHold in sandbox environments
	var headers map[string]string
	if len(responseScenario) > 0 {
		headers = make(map[string]string)
		headers["X-Response-Scenarios"] = responseScenario
	}

	response, err := c.PutWithAccessToken(endpoint, headers)

	if err != nil {
		return responseType, err
	}

	defer response.Body.Close()

	_, err = c.HandleResponse(response, &result)

	return responseType, err
}

func getEndpointFromCountry(country string) string {
	// Other countries will be implemented in the future
	if country == "SE" {
		return "/payments/domestic"
	}
	// Default FI
	return "/payments/sepa"
}
