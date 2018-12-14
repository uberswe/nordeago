package ina

import (
	"github.com/markustenghamn/nordeago"
	"github.com/markustenghamn/nordeago/ais"
)

// Response is part of the Result and can contain different data depending on the method
// that is called
type Response struct {
	Code     string          `json:"code,omitempty"`
	OrderRef string          `json:"order_ref,omitempty"`
	Status   string          `json:"status,omitempty"`
	TppToken string          `json:"tpp_token,omitempty"`
	Links    []nordeago.Link `json:"links,omitempty"`
	State    string          `json:"state,omitempty"`
	Accounts []ais.Account   `json:"accounts,omitempty"`
	Scopes   []string        `json:"scopes,omitempty"`
}

// RetrieveAccessTokenResponse represents the response returned from PollForAuthCodeDecoupled
type RetrieveAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"` // Always BEARER
}
