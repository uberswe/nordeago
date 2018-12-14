package pis

import "github.com/markustenghamn/nordeago"

// PaymentsResponse is returned when listing multiple payments via the GetPayments method
type PaymentsResponse struct {
	Payments []Payment `json:"payments"`
}

// Payment represents the structure of a payment returned via the api
type Payment struct {
	ID            string          `json:"_id"`
	Links         []nordeago.Link `json:"_links,omitempty"`
	Amount        string          `json:"amount,omitempty"`
	Currency      string          `json:"currency"`
	Creditor      Creditor        `json:"creditor"`
	Debtor        Debtor          `json:"debtor"`
	ExternalID    string          `json:"externalId,omitempty"`
	PaymentStatus string          `json:"paymentStatus,omitempty"` // PendingConfirmation, PendingUserApproval, OnHold, Confirmed, Rejected, Paid, InsufficientFunds, LimitExceeded, UserApprovalFailed, UserApprovalTimeout, UserApprovalCancelled, Unknown
	Timestamp     string          `json:"timestamp"`
}

// Creditor is part of the payment type
type Creditor struct {
	Account   Account           `json:"account"`
	Message   string            `json:"message,omitempty"`
	Name      string            `json:"name,omitempty"`
	Reference CreditorReference `json:"reference,omitempty"`
}

// Account represents the account, containing the type, currency and account number as value. Similar to ais.Account but not the same
type Account struct {
	Type     string `json:"_type"` // IBAN, BBAN_SE
	Currency string `json:"currency"`
	Value    string `json:"value"`
}

// CreditorReference represents an invoice number or reference id
type CreditorReference struct {
	Type  string `json:"_type"` // RF, INVOICE
	Value string `json:"value,omitempty"`
}

// Debtor is part of the payment type
type Debtor struct {
	AccountID string  `json:"_accountId"`
	Account   Account `json:"account"`
	Message   string  `json:"message"`
}
