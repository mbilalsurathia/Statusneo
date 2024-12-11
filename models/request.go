package models

type CreateMessageRequest struct {
	UserID    string `json:"user_id"`
	Recipient string `json:"recipient"`
	Message   string `json:"message"`
}

type UpdateMessageRequest struct {
	RequestID uint64 `json:"request_id"`
	UserID    string `json:"user_id"`
	Status    string `json:"status"`
}
