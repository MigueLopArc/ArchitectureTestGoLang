package models

type Error struct {
	ResponseCode      int    `json:"rc"`
	Message           string `json:"message"`
	Detail            string `json:"detail"`
	ExternalReference string `json:"ext_ref"`
}
