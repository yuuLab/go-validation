package response

type InvalidParams struct {
	Name   string `json:"name"`
	Reason string `json:"reason"`
}

type ValidationError struct {
	Type   string          `json:"type"`
	Title  string          `json:"title"`
	Pramas []InvalidParams `json:"invalid-params"`
}
