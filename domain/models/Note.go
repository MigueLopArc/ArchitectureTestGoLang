package models

import "time"

type Note struct {
	Id               string
	UserId           string
	Title            string
	Content          string
	CreationDate     time.Time
	ModificationDate time.Time
}
