package ina

// AuthRequestDecoupled represents the data needed for the StartAuthDecoupled function
type AuthRequestDecoupled struct {
	ResponseType string   `json:"response_type"` // 'nordea_code' or 'nordea_token' both seem to work
	PsuID        string   `json:"psu_id"`
	Scope        []string `json:"scope"`
	Language     string   `json:"language,omitempty"`
	RedirectURI  string   `json:"redirect_uri,omitempty"`
	AccountList  []string `json:"account_list"`
	Duration     int64    `json:"duration"`
	State        string   `json:"state,omitempty"`
}

// AuthRequest represents the data needed for the StartAuth function
type AuthRequest struct {
	Scope        string   `url:"scope,omitempty"`
	Language     string   `url:"language,omitempty"`
	RedirectURI  string   `url:"redirect_uri,omitempty"`
	Accounts     []string `url:"accounts,omitempty"`
	Duration     int64    `url:"duration,omitempty"`
	State        string   `url:"state,omitempty"`
	ClientID     string   `url:"client_id,omitempty"`
	MaxTxHistory string   `url:"max_tx_history,omitempty"`
	UID          string   `url:"uid,omitempty"` // For sandbox only
}

// RetrieveAccessTokenRequest requires a valid code which is returned from PollForAuthCodeDecoupled
type RetrieveAccessTokenRequest struct {
	GrantType   string `json:"grant_type"` // TODO seems undocumented but exists in postman example
	Code        string `json:"code"`
	RedirectURI string `json:"redirect_uri"`
}
