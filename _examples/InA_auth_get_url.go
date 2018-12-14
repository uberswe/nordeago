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
	"github.com/markustenghamn/nordeago/ina"
)

// run this example from the terminal by navigating to this directory and running 'go run _examples/InA_auth_get_url.go'

func main() {
	// You can get a free sandbox clientID and clientSecret via the Nordea Open Banking developer
	// portal by signing up here https://developer.nordeaopenbanking.com
	clientID := "Your-Client-Id"
	clientSecret := "Your-Client-Secret"
	redirectURI := "https://httpbin.org/get"

	// Create the client by calling InitClient with your clientID and clientSecret
	client := nordeago.InitClient(clientID, clientSecret, redirectURI)

	// Build the AuthRequest
	// I am unable to test this properly as I do not have access to a production environment
	authRequest := ina.AuthRequest{
		ClientID: clientID,
		Scope:    "ACCOUNTS_BASIC",
		Language: "SE",
		Duration: 129600,
		State:    "some id",
	}

	fmt.Println(ina.StartAuth(&client, authRequest))
}
