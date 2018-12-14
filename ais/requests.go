package ais

import "github.com/markustenghamn/nordeago"

// CreateAccountRequest is used with the CreateAccount method to create an account
type CreateAccountRequest struct {
	ID                           string          `json:"_id"`
	Links                        []nordeago.Link `json:"_links"`
	AccountName                  string          `json:"accountName"`
	AccountNumber                AccountNumber   `json:"accountNumber"`
	AccountNumbers               []AccountNumber `json:"accountNumbers"`
	AccountType                  string          `json:"accountType"` // Always 'Current'
	AvailableBalance             string          `json:"availableBalance"`
	Bank                         Bank            `json:"bank"`
	BookedBalance                string          `json:"bookedBalance"`
	Country                      string          `json:"country,omitempty"`
	Created                      string          `json:"created"`
	CreditLimit                  string          `json:"creditLimit,omitempty"`
	Currency                     string          `json:"currency"` // Currency code according to ISO 4217
	LatestTransactionBookingDate string          `json:"latestTransactionBookingDate,omitempty"`
	OwnerName                    string          `json:"ownerName"`
	Product                      string          `json:"product"`
	Status                       string          `json:"status"` // OPEN or CLOSED
	ValueDatedBalance            string          `json:"valueDatedBalance,omitempty"`
}

// AccountNumber represents an account number
type AccountNumber struct {
	Type  string `json:"_type"` //IBAN or BBAN_SE
	Value string `json:"value,omitempty"`
}

// GetAccountTransactionsRequest is used with the GetAccountTransactions method to list transactions for the specified account id
type GetAccountTransactionsRequest struct {
	FromDate        string `json:"fromDate"`
	ToDate          string `json:"toDate"`
	Language        string `json:"language"`
	ContinuationKey string `json:"continuationKey"`
}

// Transaction is used to create or return a transaction
type Transaction struct {
	Type                    string `json:"_type"` // CreditTransaction or DebitTransaction
	Amount                  string `json:"amount,omitempty"`
	BalanceAfterTransaction string `json:"balanceAfterTransaction,omitempty"`
	BookingDate             string `json:"bookingDate"`
	CardNumber              string `json:"cardNumber,omitempty"`
	CounterpartyName        string `json:"counterpartName,omitempty"`
	Currency                string `json:"currency"`
	CurrencyRate            string `json:"currencyRate,omitempty"`
	Message                 string `json:"message,omitempty"`
	Narrative               string `json:"narrative,omitempty"`
	OriginalCurrency        string `json:"originalCurrency,omitempty"`
	OriginalCurrencyAmount  string `json:"originalCurrencyAmount,omitempty"`
	OwnMessage              string `json:"ownMessage,omitempty"`
	PaymentDate             string `json:"paymentDate,omitempty"`
	Reference               string `json:"reference,omitempty"`
	Status                  string `json:"status"`
	TransactionDate         string `json:"transactionDate,omitempty"`
	TransactionID           string `json:"transactionId"`
	TypeDescription         string `json:"typeDescription,omitempty"`
	ValueDate               string `json:"valueDate,omitempty"`
}
