package models

type SignUpModel struct {
	SignInModel
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}
