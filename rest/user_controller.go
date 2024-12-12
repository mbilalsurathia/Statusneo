package rest

import (
	"github.com/gin-gonic/gin"
	"maker-checker/models"
	"maker-checker/service"
	"maker-checker/utils"
	"net/http"
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
		ctx.JSON(http.StatusBadRequest, utils.NewStandardResponse(false, models.INVALID_INPUT, models.INVALID_INPUT_MESSAGE, nil))
		return
	}

	err = request.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewStandardResponse(false, models.INVALID_INPUT, err.Error(), nil))
		return
	}

	//service method call to create Message
	message, err := u.userService.CreateMessage(request)
	if err != nil {
		utils.HandleServiceError(ctx, err)
	}

	ctx.JSON(http.StatusOK, utils.NewStandardResponse(true, models.SUCCESS, models.SUCCESSFULLY, message))
}

// UpdateMessage Update Message API
func (u *userController) UpdateMessage(ctx *gin.Context) {
	var request *models.UpdateMessageRequest

	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewStandardResponse(false, models.INVALID_INPUT, models.INVALID_INPUT_MESSAGE, nil))
		return
	}

	err = request.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewStandardResponse(false, models.INVALID_INPUT, err.Error(), nil))
		return
	}

	//service method call to Update Message for Approve or Reject
	messages, err := u.userService.UpdateMessage(request)
	if err != nil {
		utils.HandleServiceError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, utils.NewStandardResponse(true, models.SUCCESS, models.SUCCESSFULLY, messages))
}

// GetMessages Get Message API
func (u *userController) GetMessages(ctx *gin.Context) {
	messageId, err := utils.ParseMessageID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewStandardResponse(false, models.BAD_REQUEST, err.Error(), nil))
		return
	}

	//service method call to Get Messages by MessageId or All messages
	messages, err := u.userService.GetMessages(messageId)
	if err != nil {
		utils.HandleServiceError(ctx, err)
	}

	ctx.JSON(http.StatusOK, utils.NewStandardResponse(true, models.SUCCESS, models.SUCCESSFULLY, messages))
}
