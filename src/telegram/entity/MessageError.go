package entity

// MessageError is received from telegram if the Message previously sent is NOK
type MessageError struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}
