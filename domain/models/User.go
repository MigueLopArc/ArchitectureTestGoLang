package models

import (
	"time"
)

type User struct {
	Id           string
	Email        string
	Firstname    string
	Lastname     string
	Password     string
	CreationDate time.Time
}
