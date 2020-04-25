package v1

import (
	"api.xinfos.com/api"
	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/service"
	"api.xinfos.com/utils/errs"
	"api.xinfos.com/utils/identity"

	"github.com/gin-gonic/gin"
)

//GetUserInfoByIDRequest - request
type GetUserInfoByIDRequest struct {
	Request string `json:"request_id"`
	UserID  uint64 `json:"user_id"`
}

//GetUserInfoByID - Get user info by user_id
func GetUserInfoByID(c *gin.Context) {
	var req GetUserInfoByIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}
	user, _ := service.NewUserService().GetUserInfoByID(req.UserID)
	api.JSON(c, errs.ErrSuccess, user)
	return
}

//GetAllUsers - request
func GetAllUsers(c *gin.Context) {

}

type CreateUserRequest struct {
	Request string `json:"request_id"`
	Name    string `json:"name"  binding:"required"`
	Phone   string `json:"phone"  binding:"required"`
	IDCard  string `json:"id_card" binding:"required"`
}

//CreateUser - create user request
func CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}

	//身份证校验
	idCard, err := identity.IsValidCitizenNo18(req.IDCard)
	if err != nil {
		api.JSON(c, errs.ErrParamInvalid, nil)
		return
	}

	userId, err := service.NewUserService().Create(&model.User{
		Name:     req.Name,
		Phone:    req.Phone,
		IDCard:   req.IDCard,
		Birthday: idCard.Birthday,
		Gender:   idCard.Gender,
		Age:      idCard.Age,
	})
	if err != nil {
		api.JSON(c, errs.ErrParamInvalid, nil)
		return
	}

	api.JSON(c, errs.ErrSuccess, map[string]uint64{"user_id": userId})
	return
}
