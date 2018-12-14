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

package main

import (
	"fmt"
	"github.com/markustenghamn/nordeago"
	"github.com/markustenghamn/nordeago/ais"
	"github.com/markustenghamn/nordeago/ina"
	"net/http"
	"time"
)

// run this example from the terminal by navigating to this directory and running 'go run _examples/AIS_create_account.go'

func main() {
	// You can get a free sandbox clientID and clientSecret via the Nordea Open Banking developer
	// portal by signing up here https://developer.nordeaopenbanking.com

	clientID := "Your-Client-Id"
	clientSecret := "Your-Client-Secret"
	redirectURI := "https://httpbin.org/get"

	// Create the client by calling InitClient with your clientID and clientSecret
	client := nordeago.InitClient(clientID, clientSecret, redirectURI)

	// Build the AuthRequestDecoupled
	// Documentation can be found here https://developer.nordeaopenbanking.com/app/documentation?api=Identity%20and%20Access%20API&version=2.1#authorize
	authRequest := ina.AuthRequestDecoupled{
		ResponseType: "nordea_code",
		PsuID:        "193805010844",
		Scope:        []string{"ACCOUNTS_BASIC", "PAYMENTS_MULTIPLE", "ACCOUNTS_TRANSACTIONS", "ACCOUNTS_DETAILS", "ACCOUNTS_BALANCES"},
		Language:     "FI",
		RedirectURI:  redirectURI,
		AccountList:  []string{"41770042136"},
		Duration:     129600,
		State:        "some id",
	}

	// Pass the request to the StartAuthDecoupled of the client to make the request
	authResponse, err := ina.StartAuthDecoupled(&client, authRequest)

	// Check for any errors
	if err != nil {
		panic(err)
	}

	fmt.Printf("Polling for auth code every 1 seconds...")
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf(".")
		// We get the order_ref and poll for access token
		_, statusCode, err := ina.PollForAuthCodeDecoupled(&client, authResponse.OrderRef)

		// Check for any errors
		if err != nil {
			panic(err)
		}

		if statusCode != http.StatusNotModified {
			fmt.Printf("\nresponse received\n")
			break
		}

	}

	retrieveAccessTokenRequest := ina.RetrieveAccessTokenRequest{
		GrantType:   "authorization_code",
		Code:        client.AuthCode,
		RedirectURI: client.RedirectURL,
	}

	_, err = ina.RetrieveAccessTokenDecoupled(&client, retrieveAccessTokenRequest)

	// Check for any errors
	if err != nil {
		panic(err)
	}

	// Make request to create an account
	createAccountRequest := ais.CreateAccountRequest{
		Bank: ais.Bank{
			BIC:     "NDEAFIHH",
			Country: "FI",
			Name:    "no example",
		},
		ID:                           "no example",
		Links:                        []nordeago.Link{},
		AccountName:                  "Test account",
		AccountNumber:                ais.AccountNumber{Type: "IBAN", Value: "FI1350001520000081"},
		AccountNumbers:               []ais.AccountNumber{{Type: "IBAN", Value: "FI1350001520000081"}},
		AccountType:                  "Current",
		AvailableBalance:             "1000.00",
		BookedBalance:                "1000.00",
		Country:                      "FI",
		CreditLimit:                  "1000.00",
		Created:                      "2017-06-01",
		Currency:                     "EUR",
		LatestTransactionBookingDate: "2017-06-01",
		OwnerName:                    "no example",
		Product:                      "Example Product",
		Status:                       "OPEN",
		ValueDatedBalance:            "123.45",
	}

	created, err := ais.CreateAccount(&client, createAccountRequest)

	// print our response
	if created {
		fmt.Printf("Account created\n")
	} else {
		fmt.Printf("Account was not created\n")
	}

	if err != nil {
		fmt.Println(err)
		return
	}
}
