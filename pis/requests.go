package pis

// InitiatePaymentRequest represents the parameters for creating a payment
type InitiatePaymentRequest struct {
	Amount     string   `json:"amount,omitempty"`
	Currency   string   `json:"currency"`
	Creditor   Creditor `json:"creditor"`
	Debtor     Debtor   `json:"debtor"`
	ExternalID string   `json:"externalId,omitempty"`
}
