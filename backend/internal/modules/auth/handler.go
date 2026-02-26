package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/codeApe-7/go-user-center/internal/modules/user"
	"github.com/codeApe-7/go-user-center/internal/utils"
)

type Handler struct {
	DB        *gorm.DB
	JWTSecret string
	Issuer    string
	AccessTTL time.Duration
}

type registerReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=64"`
	Nickname string `json:"nickname" binding:"required,min=2,max=32"`
}

type loginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) Register(c *gin.Context) {
	var req registerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existing user.User
	if err := h.DB.Where("email = ?", req.Email).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "email already registered"})
		return
	}

	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "hash password failed"})
		return
	}

	u := user.User{
		UUID:         uuid.NewString(),
		Email:        req.Email,
		PasswordHash: hash,
		Nickname:     req.Nickname,
	}
	if err := h.DB.Create(&u).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "create user failed"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": u})
}

func (h *Handler) Login(c *gin.Context) {
	var req loginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var u user.User
	if err := h.DB.Where("email = ?", req.Email).First(&u).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}
	if !utils.CheckPassword(u.PasswordHash, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	tok, err := utils.SignAccessToken(h.JWTSecret, h.Issuer, u.UUID, u.Email, h.AccessTTL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "sign token failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"accessToken": tok, "user": u})
}
