package users

import (
	usersDomain "ToDoList/internal/domain/users"
	"ToDoList/internal/service/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	RegisterUser(usersDomain.RegisterRequest) (string, error)
	LoginUser(usersDomain.LoginRequest) (usersDomain.User, error)
	GetUsers() ([]usersDomain.User, error)
	FindUserByID(string) (usersDomain.User, error)
	UpdateUser(usersDomain.UpdateRequest, string) (usersDomain.User, error)
	DeleteUser(string) error
}

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) Register(ctx *gin.Context) {
	var req usersDomain.RegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid, err := uh.userService.RegisterUser(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"uid": uid})
}

func (uh *UserHandler) Login(ctx *gin.Context) {
	var req usersDomain.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uh.userService.LoginUser(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := auth.GenerateAccessToken(user.UID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	refreshToken, err := auth.GenerateRefreshToken(user.UID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie(
		"refresh_token",
		refreshToken,
		60*60*24*30,
		"/",
		"",
		false,
		true,
	)
	ctx.JSON(http.StatusOK, gin.H{"access_token": accessToken})

}

func (uh *UserHandler) GetUsers(ctx *gin.Context) {
	users, err := uh.userService.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (uh *UserHandler) FindUserByID(ctx *gin.Context) {
	uid := ctx.Param("id")
	user, err := uh.userService.FindUserByID(uid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
	uid := ctx.Param("id")
	var req usersDomain.UpdateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uh.userService.UpdateUser(req, uid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
func (uh *UserHandler) DeleteUser(ctx *gin.Context) {
	uid := ctx.Param("id")
	err := uh.userService.DeleteUser(uid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
