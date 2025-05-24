package models

type Report struct {
	Document_uri       string `json:"document-uri" binding:"required"`
	Referrer           string `json:"referrer" binding:"required"`
	Blocked_uri        string `json:"blocked-uri" binding:"required"`
	Violated_directive string `json:"violated-directive" binding:"required"`
	Original_policy    string `json:"original-policy" binding:"required"`
}
