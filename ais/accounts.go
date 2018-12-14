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

package ais

import (
	"errors"
	"github.com/google/go-querystring/query"
	"github.com/markustenghamn/nordeago"
	"net/http"
)

// Note: Different countries can have different variables for different methods.

// ListAccounts lists the accounts for the user
//
// API Documentation: https://developer.nordeaopenbanking.com/app/documentation?api=Accounts%20API&version=2.3#accountList
func ListAccounts(c *nordeago.Client) (ListAccountsResponse, error) {
	responseType := ListAccountsResponse{}
	result := nordeago.Result{Response: &responseType}
	endpoint := "/accounts"

	response, err := c.GetWithAccessToken(endpoint, nil)

	if err != nil {
		return responseType, err
	}

	defer response.Body.Close()

	_, err = c.HandleResponse(response, &result)

	return responseType, err
}

// CreateAccount for sandbox environment
//
// API Documentation: https://developer.nordeaopenbanking.com/app/documentation?api=Accounts%20API&version=2.3#createAccountV2
func CreateAccount(c *nordeago.Client, request CreateAccountRequest) (bool, error) {
	result := nordeago.Result{}

	endpoint := "/accounts"
	// Returns a 201 status code if created

	response, err := c.PostWithAccessToken(endpoint, request, nil)

	if err != nil {
		return false, err
	}

	defer response.Body.Close()

	status, err := c.HandleResponse(response, &result)

	if status == http.StatusCreated {
		return true, err
	}

	return false, err

}

// GetAccountDetails gets account details for the specified account
//
// API Documentation: https://developer.nordeaopenbanking.com/app/documentation?api=Accounts%20API&version=2.3#accountDetails
func GetAccountDetails(c *nordeago.Client, accountID string) (*AccountDetailed, error) {
	responseType := &AccountDetailed{}
	result := nordeago.Result{Response: responseType}

	endpoint := "/accounts/{{accountId}}"
	endpoint = nordeago.ReplaceVariable(endpoint, "accountId", accountID)

	response, err := c.GetWithAccessToken(endpoint, nil)

	if err != nil {
		return responseType, err
	}

	defer response.Body.Close()

	_, err = c.HandleResponse(response, &result)

	return responseType, err
}

// DeleteAccount for sandbox environment
//
// API Documentation: https://developer.nordeaopenbanking.com/app/documentation?api=Accounts%20API&version=2.3#deleteUserDefinedAccount
func DeleteAccount(c *nordeago.Client, accountID string) (string, error) {
	responseType := ""
	result := nordeago.Result{Response: &responseType}

	endpoint := "/accounts/{{accountId}}"
	endpoint = nordeago.ReplaceVariable(endpoint, "accountId", accountID)

	response, err := c.DeleteWithAccessToken(endpoint, nil)

	if err != nil {
		return responseType, err
	}

	defer response.Body.Close()

	_, err = c.HandleResponse(response, &result)

	return responseType, err
}

// GetAccountTransactions gets the transactions for the specified account
//
// API Documentation: https://developer.nordeaopenbanking.com/app/documentation?api=Accounts%20API&version=2.3#transactionsList
func GetAccountTransactions(c *nordeago.Client, accountID string, request GetAccountTransactionsRequest) (*GetAccountTransactionsResponse, error) {
	responseType := &GetAccountTransactionsResponse{}
	result := nordeago.Result{Response: responseType}

	endpoint := "/accounts/{{accountId}}/transactions"
	endpoint = nordeago.ReplaceVariable(endpoint, "accountId", accountID)

	req, err := http.NewRequest("GET", c.GetFullURL(endpoint), nil)
	if err != nil {
		return responseType, err
	}

	v, _ := query.Values(request)

	req.URL.RawQuery = v.Encode()

	response, err := c.GetWithAccessToken(req.URL.String(), nil)

	if err != nil {
		return responseType, err
	}

	defer response.Body.Close()

	_, err = c.HandleResponse(response, &result)

	return responseType, err
}

// CreateAccountTransaction creates a transaction
//
// API Documentation: https://developer.nordeaopenbanking.com/app/documentation?api=Accounts%20API&version=2.3#createTransaction
func CreateAccountTransaction(c *nordeago.Client, accountID string, request Transaction) (bool, error) {
	endpoint := "/accounts/{{accountId}}/transactions"
	endpoint = nordeago.ReplaceVariable(endpoint, "accountId", accountID)
	// Returns a 201 status code if created

	response, err := c.PostWithAccessToken(endpoint, request, nil)

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
