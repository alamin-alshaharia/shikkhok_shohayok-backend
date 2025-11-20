package controllers

import (
	"net/http"
	"shikkhok_shohayok/db/sqlc"
	"shikkhok_shohayok/util"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	store  *db.Queries
	config util.Config
}

func NewAuthController(store *db.Queries, config util.Config) *AuthController {
	return &AuthController{store: store, config: config}
}

type registerRequest struct {
	Phone      string `json:"phone" binding:"required"`
	Password   string `json:"password" binding:"required,min=6"`
	FullName   string `json:"full_name" binding:"required"`
	SchoolName string `json:"school_name" binding:"required"`
}

func (ctrl *AuthController) Register(ctx *gin.Context) {
	var req registerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	arg := db.CreateUserParams{
		Phone:      req.Phone,
		Password:   hashedPassword,
		FullName:   req.FullName,
		SchoolName: req.SchoolName,
	}

	user, err := ctrl.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type loginRequest struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ctrl *AuthController) Login(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.store.GetUserByPhone(ctx, req.Phone)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := util.CreateToken(user.Phone, 24*time.Hour, ctrl.config.TokenSymmetricKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}
