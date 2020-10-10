package model

// APIResponse ...
type APIResponse struct {
	Code    int         `json:"code"`
	Value   interface{} `json:"value"`
	Message string      `json:"message"`
}
