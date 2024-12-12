package models

import "time"

type InitUserRes struct {
	Token string `json:"token"`
}

type MessageResp struct {
	RequestID  uint64    `json:"request_id"`
	Sender     string    `json:"sender"`
	Recipient  string    `json:"recipient"`
	Message    string    `json:"message"`
	Status     string    `json:"status"`
	ApprovedBy string    `json:"approved_by,omitempty"`
	RejectedBy string    `json:"rejected_by,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}


type StandardResponse struct {
	Result  bool        `json:"result"`
	Code    uint        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
