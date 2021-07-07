package models

type SignInModel struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
