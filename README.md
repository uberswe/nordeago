# NordeaGo - Golang wrapper for Nordea Open Banking

[![GoDoc](https://godoc.org/github.com/markustenghamn/nordeago?status.svg)](https://godoc.org/github.com/markustenghamn/nordeago)
[![Go Report Card](https://goreportcard.com/badge/github.com/markustenghamn/nordeago)](https://goreportcard.com/report/github.com/markustenghamn/nordeago)
[![Build Status](https://api.travis-ci.org/markustenghamn/nordeago.svg?branch=master)](https://travis-ci.org/markustenghamn/nordeago)

NordeaGo is a wrapper for the [Nordea Open Banking](https://developer.nordeaopenbanking.com/app/docs) API version 2.3 written in Go.

This was written and tested with the sandbox API provided by Nordea, I am not able to test this in a production setting as I do not have the required licenses.

Please note that this is a work in progress and I may have made mistakes here and there.

### Sample

I tried to keep things simple, here is an example of how to authenticate. More examples are available in the _examples folder.

```go
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
		Language:     "SE",
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

	// print our response
	fmt.Printf("%+v\n", authResponse)
}
```

## Bugs and Errors

If you find a bug or error in the wrapper please open an issue here directly. If you find an issue or have a question about the Nordea API please use their [support page](https://support.nordeaopenbanking.com/hc/en-us).

## Contributing

I am always open to contributions. You are allowed to keep your copyright of your contributions, the only thing I require is that it follows the same license as this repository. If you would like to contribute please open a pull request.

## Example code

Please see the examples folder for examples of how to use this library

## License

This library is distributed under the MIT license found in the LICENSE file.