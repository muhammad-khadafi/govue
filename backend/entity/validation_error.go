package entity

type ValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}
