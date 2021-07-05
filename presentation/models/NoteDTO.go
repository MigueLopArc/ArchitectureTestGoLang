package request

import "time"

type NoteDTO struct {
	Id               string    `json:"id,omitempty"`
	UserId           string    `json:"userId,omitempty"`
	Title            string    `json:"title,omitempty"`
	Content          string    `json:"content,omitempty"`
	CreationDate     time.Time `json:"creationDate,omitempty"`
	ModificationDate time.Time `json:"modificationDate,omitempty"`
}
