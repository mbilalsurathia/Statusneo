package models

import "time"

type Message struct {
	MessageId  uint64    `gorm:"primaryKey;column:message_id;autoIncrement;comment:message_id" json:"message_id"`
	Sender     string    `json:"sender"`
	Recipient  string    `json:"recipient"`
	Message    string    `json:"message"`
	Status     string    `json:"status"` // "Pending", "Approved", or "Rejected"
	ApprovedBy string    `json:"approved_by,omitempty"`
	RejectedBy string    `json:"rejected_by,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
