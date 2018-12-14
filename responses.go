package nordeago

// Result represents returned data
type Result struct {
	GroupHeader GroupHeader   `json:"groupHeader,omitempty"`
	Response    interface{}   `json:"response,omitempty"` // Response can have many formats, string or object
	Error       ErrorResponse `json:"error,omitempty"`
}

// GroupHeader is a general response object returned after a request and gives the HTTP status code along with
// a creation datetime and message ID that can be used for debugging purposes
type GroupHeader struct {
	MessageIdentification string              `json:"messageIdentification"`
	CreationDateTime      string              `json:"creationDateTime"`
	HTTPCode              int64               `json:"httpCode"`
	MessagePagination     []MessagePagination `json:"messagePagination,omitempty"`
}

// MessagePagination Resource listing may return a continuationKey if there's more results available.
// Request may be retried with the continuationKey, but otherwise same parameters, in order to get more results.
type MessagePagination struct {
	ContinuationKey string `json:"continuationKey"`
}

// ErrorResponse is returned if a request fails
type ErrorResponse struct {
	HTTPCode        int          `json:"httpCode,omitempty"`
	HTTPMessage     string       `json:"httpMessage,omitempty"`
	MoreInformation string       `json:"moreInformation,omitempty"`
	Request         RequestError `json:"request,omitempty"`
	Failures        []Failure    `json:"failures,omitempty"`
}

// RequestError contains the request message identifier and the url that was called in the original request
type RequestError struct {
	MessageIdentifier string `json:"messageIdentifier"`
	URL               string `json:"url"`
}

// Failure contains a failure code along with a description of the failure
type Failure struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

// Link represents a hyperlink with parameters, usually part of a list of other actions that can be performed on the entity
type Link struct {
	Rel         string `json:"rel"`
	Href        string `json:"href"`
	Deprecation string `json:"deprecation,omitempty"`
	Hreflang    string `json:"hreflang,omitempty"`
	Media       string `json:"media,omitempty"`
	Templated   bool   `json:"templated,omitempty"`
	Title       string `json:"title,omitempty"`
	Type        string `json:"type,omitempty"`
}
