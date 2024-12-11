package models

import (
	"fmt"
)

var invalidInputError = fmt.Errorf("invalid data")

var invalidStatusError = fmt.Errorf("status should be Approve or Reject")

func (s *CreateMessageRequest) Validate() error {

	if s.UserID == "" {
		return invalidInputError
	}
	if s.Recipient == "" {
		return invalidInputError
	}
	if s.Message == "" {
		return invalidInputError
	}

	return nil
}

func (s *UpdateMessageRequest) Validate() error {

	if s.RequestID == 0 {
		return invalidInputError
	}
	if s.Status != APPROVE && s.Status != REJECT {
		return invalidStatusError
	}

	return nil
}
