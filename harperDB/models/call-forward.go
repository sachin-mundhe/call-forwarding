package models

// CallForwardingBody http request body to be parsed
type CallForwardingBody struct {
	PhoneNumber string `json:"phoneNumber"`
}

// CallForwardStatusResp http response for call forward status
type CallForwardStatusResp struct {
	Active      *bool   `json:"active,omitempty"`
	Status      *int    `json:"status,omitempty"`
	Code        *string `json:"code,omitempty"`
	Message     *string `json:"message,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}
