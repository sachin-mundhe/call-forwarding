package models

// CallForwardingBody http request body to be parsed
type CallForwardingBody struct {
	PhoneNumber string `json:"phoneNumber"`
}

// UnconditionalCallForwardingStatusHttpResp http response for unconditional call forwarding status
type UnconditionalCallForwardingStatusHttpResp struct {
	Active bool `json:"active"`
}

// UnconditionalCallForwardingStatusHttpResp http response for unconditional call forwarding status
type UnconditionalCFStatusDb struct {
	UnconditionalCFStatus bool   `json:"unconditionalCallForwardingStatus,omitempty"`
	PhoneNumber           string `json:"phoneNumber,omitempty"`
}

// ConditionalCallForwardingHttpResp http response for conditional call forward status
type ConditionalCallForwardingHttpResp struct {
	PhoneNumber             string   `json:"phoneNumber,omitempty"`
	ConditionalCFStatus     bool     `json:"conditionalCallForwardingStatus,omitempty"`
	CallForwadingConditions []string `json:"callForwardingConditions,omitempty"`
}

// HttpErrorResp http error response
type HttpErrorResp struct {
	Status  *int    `json:"status,omitempty"`
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}
