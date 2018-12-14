package ais

import "github.com/markustenghamn/nordeago"

// AccountDetailed is returned as part of the ListAccountsResponse when fetching account details
type AccountDetailed struct {
	AccountName                  string          `json:"accountName"`
	AccountNumber                AccountNumber   `json:"accountNumber,omitempty"`
	AccountNumbers               []AccountNumber `json:"accountNumbers,omitempty"`
	ID                           string          `json:"_id,omitempty"`
	Links                        []nordeago.Link `json:"_links,omitempty"`
	AccountType                  string          `json:"accountType"` // Always Current
	AvailableBalance             string          `json:"availableBalance"`
	Bank                         Bank            `json:"bank"`
	BookedBalance                string          `json:"bookedBalance"`
	Country                      string          `json:"country,omitempty"`
	CreditLimit                  string          `json:"creditLimit,omitempty"`
	Currency                     string          `json:"currency"`
	LatestTransactionBookingDate string          `json:"latestTransactionBookingDate,omitempty"`
	OwnerName                    string          `json:"ownerName,omitempty"`
	Product                      string          `json:"product"`
	Status                       string          `json:"status"`
	ValueDatedBalance            string          `json:"valueDatedBalance,omitempty"`
}

// Bank represents a bank entity in request and response types
type Bank struct {
	BIC     string `json:"bic"`
	Country string `json:"country"`
	Name    string `json:"name"`
}

// Account is part of the Response and represents a users accounts
type Account struct {
	AccountID     string `json:"AccountId"`
	AccountNumber string `json:"AccountNumber"`
	Currency      string `json:"Currency"`
}

// GetAccountTransactionsResponse represents returned data from GetAccountTransactions and is
// part of the GetAccountTransactionsResult object
type GetAccountTransactionsResponse struct {
	ContinuationKey string          `json:"continuationKey"`
	Links           []nordeago.Link `json:"links"`
	Transactions    []Transaction   `json:"transactions"`
}

// ListAccountsResponse contains a list of AccountDetailed types
type ListAccountsResponse struct {
	Accounts []AccountDetailed `json:"accounts"`
}
