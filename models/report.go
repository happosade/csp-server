package models

type Report struct {
	Document_uri       string `json:"document-uri"`
	Referrer           string `json:"referrer"`
	Blocked_uri        string `json:"blocked-uri"`
	Violated_directive string `json:"violated-directive"`
	Original_policy    string `json:"original-policy"`
}
