package rest

import (
	"maker-checker/models"
	"maker-checker/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateMessage(ctx *gin.Context)
	UpdateMessage(ctx *gin.Context)
	GetMessages(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{userService: userService}
}

// CreateMessage Create Message API
func (u *userController) CreateMessage(ctx *gin.Context) {
	var request *models.CreateMessageRequest

	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewStandardResponse(false, models.INVALID_INPUT, models.INVALID_INPUT_MESSAGE, nil))
		return
	}

	err = request.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewStandardResponse(false, models.INVALID_INPUT, err.Error(), nil))
		return
	}

	//service method call to create Message
	message, err := u.userService.CreateMessage(request)
	if err != nil {
		if standardError, ok := err.(*models.StandardError); ok {
			ctx.JSON(http.StatusBadRequest, NewStandardResponse(false, standardError.Code, standardError.Message, nil))
			return
		} else {
			ctx.JSON(http.StatusBadRequest, NewStandardResponse(false, models.INTERNAL_SERVER_ERROR, err.Error(), nil))
			return
		}
	}

	ctx.JSON(http.StatusOK, NewStandardResponse(true, models.SUCCESS, models.SUCCESSFULLY, message))
}

// UpdateMessage Update Message API
func (u *userController) UpdateMessage(ctx *gin.Context) {
	var request *models.UpdateMessageRequest

	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewStandardResponse(false, models.INVALID_INPUT, models.INVALID_INPUT_MESSAGE, nil))
		return
	}

	err = request.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewStandardResponse(false, models.INVALID_INPUT, err.Error(), nil))
		return
	}

	//service method call to Update Message for Approve or Reject
	messages, err := u.userService.UpdateMessage(request)
	if err != nil {
		if standardError, ok := err.(*models.StandardError); ok {
			ctx.JSON(http.StatusBadRequest, NewStandardResponse(false, standardError.Code, standardError.Message, nil))
			return
		} else {
			ctx.JSON(http.StatusBadRequest, NewStandardResponse(false, models.INTERNAL_SERVER_ERROR, err.Error(), nil))
			return
		}
	}

	ctx.JSON(http.StatusOK, NewStandardResponse(true, models.SUCCESS, models.SUCCESSFULLY, messages))
}

// GetMessages Get Message API
func (u *userController) GetMessages(ctx *gin.Context) {
	messageId := uint64(0)
	requestId := ctx.Query("messageId")
	if requestId != "" {
		messageIdInt, err := strconv.Atoi(requestId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, NewStandardResponse(false, models.BAD_REQUEST, err.Error(), nil))
			return
		}
		messageId = uint64(messageIdInt)
	}

	//service method call to Get Messages by MessageId or All messages
	messages, err := u.userService.GetMessages(messageId)
	if err != nil {
		if standardError, ok := err.(*models.StandardError); ok {
			ctx.JSON(http.StatusBadRequest, NewStandardResponse(false, standardError.Code, standardError.Message, nil))
			return
		} else {
			ctx.JSON(http.StatusBadRequest, NewStandardResponse(false, models.INTERNAL_SERVER_ERROR, err.Error(), nil))
			return
		}
	}

	ctx.JSON(http.StatusOK, NewStandardResponse(true, models.SUCCESS, models.SUCCESSFULLY, messages))
}
