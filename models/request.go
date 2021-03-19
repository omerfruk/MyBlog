package models

type RequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type RequestSingUp struct {
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
