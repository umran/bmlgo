package bml

// AuthResponse ...
type AuthResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Payload int    `json:"payload"`
}
