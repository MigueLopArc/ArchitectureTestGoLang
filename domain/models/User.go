package models

type User struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Password  string `json:"password,omitempty"`
	Token     string `json:"token,omitempty"`
	CreatedAt string `json:"date_created,omitempty"`
}
