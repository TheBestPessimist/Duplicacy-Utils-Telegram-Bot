package entity

// BackupNotification stores a notification received from the backup script.
// It will be sent to the ChatId telegram chat.
type BackupNotification struct {
	ChatId  int64  `json:"chat_id"`
	Content string `json:"content"`
}
