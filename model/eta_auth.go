package model

type EtaLoginRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

type EtaLoginResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type EtaSubmitInvoiceResponse struct {
	SubmissionId      string `json:"submissionId"`
	AcceptedDocuments string `json:"acceptedDocuments"`
	RejectedDocuments string `json:"rejectedDocuments"`
}
