package models

type Report struct {
	document_uri       string
	referrer           string
	blocked_uri        string
	violated_directive string
	original_policy    string
}
