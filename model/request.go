package model

// APIRequest ...
type APIRequest interface{}

// RegisterRequest ...
type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
