package controllers

import (
	"book-api/internal/middleware"
	"book-api/internal/models"
	"book-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{ svc *services.UserService }

func NewUserController(s *services.UserService) *UserController { return &UserController{s} }

type loginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (uc *UserController) Register(c *gin.Context) {
	var req loginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := uc.svc.FindByUsername(req.Username)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "username already exists"})
		return
	}

	hashed, err := middleware.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	user := models.User{
		Username: req.Username,
		Password: hashed,
	}

	id, err := uc.svc.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully",
		"user_id": id,
	})
}


func (uc *UserController) Login(c *gin.Context) {
	var req loginReq
	if err := c.ShouldBindJSON(&req); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }

	user, err := uc.svc.FindByUsername(req.Username)
	if err != nil { c.JSON(http.StatusUnauthorized, gin.H{"error":"invalid credentials"}); return }

	if !middleware.CheckPassword(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"invalid credentials"})
		return
	}

	token, err := middleware.GenerateJWT(user.Username)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error":"token error"}); return }
	c.JSON(http.StatusOK, gin.H{"token": token})
}
