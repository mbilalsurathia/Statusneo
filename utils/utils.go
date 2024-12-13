package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"maker-checker/conf"
	"maker-checker/models"
	"net/http"
	"net/smtp"
	"strconv"
)


func SendEmail(emailConfig conf.Email, sendName string, recipient string) (err error) {

	subject := fmt.Sprintf("Subject: Test Email from %v Status Neo", sendName)
	body := "Hello, this is a test email sent from Go!"

	// Combine subject and body
	message := []byte(subject + "\n" + body)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v", r)
			LogError("Recovered from panic", err)
		}
	}()

	// Authentication
	auth := smtp.PlainAuth("", emailConfig.From, emailConfig.Password, emailConfig.SmtpHost)

	// Sending the email
	err = smtp.SendMail(emailConfig.SmtpHost+":"+emailConfig.SmtpPort, auth, emailConfig.From, []string{recipient}, message)
	if err != nil {
		LogError("Failed to send email", err)
	}

	return err
}

func ParseMessageID(ctx *gin.Context) (uint64, error) {
	requestId := ctx.Query("messageId")
	if requestId == "" {
		return 0, nil
	}

	messageIdInt, err := strconv.Atoi(requestId)
	if err != nil {
		return 0, err
	}

	return uint64(messageIdInt), nil
}

func LogError(message string, err error) {
	zap.L().Error(message, zap.Error(err))
}

func HandleServiceError(ctx *gin.Context, err error) {
	if standardError, ok := err.(*models.StandardError); ok {
		ctx.JSON(http.StatusBadRequest, NewStandardResponse(false, standardError.Code, standardError.Message, nil))
	} else {
		ctx.JSON(http.StatusBadRequest, NewStandardResponse(false, models.INTERNAL_SERVER_ERROR, err.Error(), nil))
	}
}

func NewStandardResponse(result bool, code uint, msg string, data interface{}) *models.StandardResponse {

	if data == nil {
		data = ""
	}
	return &models.StandardResponse{
		Result:  result,
		Code:    code,
		Message: msg,
		Data:    data,
	}
}
