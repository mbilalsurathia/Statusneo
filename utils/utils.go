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

func SendEmail(emailConfig conf.Email, sendName string, recipient string) error {
	subject := fmt.Sprintf("Subject: Test Email from %v Status Neo", sendName)
	body := "Hello, this is a test email sent from Go!"

	// Combine subject and body
	message := []byte(subject + "\n" + body)

	// Authentication
	auth := smtp.PlainAuth("", emailConfig.From, emailConfig.Password, emailConfig.SmtpHost)

	// Sending the email
	err := smtp.SendMail(emailConfig.SmtpHost+":"+emailConfig.SmtpPort, auth, emailConfig.From, []string{recipient}, message)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
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
		ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(false, standardError.Code, standardError.Message, nil))
	} else {
		ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(false, models.INTERNAL_SERVER_ERROR, err.Error(), nil))
	}
}
