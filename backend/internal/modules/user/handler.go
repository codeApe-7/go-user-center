package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/codeApe-7/go-user-center/internal/utils"
)

type Handler struct {
	DB *gorm.DB
}

type updateMeReq struct {
	Nickname  *string `json:"nickname" binding:"omitempty,min=2,max=32"`
	AvatarURL *string `json:"avatarUrl" binding:"omitempty,url"`
	Bio       *string `json:"bio" binding:"omitempty,max=255"`
}

type changePasswordReq struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required,min=6,max=64"`
}

func (h *Handler) GetMe(c *gin.Context) {
	uuidAny, _ := c.Get("userUuid")
	uuid, _ := uuidAny.(string)

	var u User
	if err := h.DB.Where("uuid = ?", uuid).First(&u).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": u})
}

func (h *Handler) UpdateMe(c *gin.Context) {
	uuidAny, _ := c.Get("userUuid")
	uuid, _ := uuidAny.(string)

	var req updateMeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := map[string]any{}
	if req.Nickname != nil {
		updates["nickname"] = *req.Nickname
	}
	if req.AvatarURL != nil {
		updates["avatar_url"] = *req.AvatarURL
	}
	if req.Bio != nil {
		updates["bio"] = *req.Bio
	}

	if len(updates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no fields to update"})
		return
	}

	if err := h.DB.Model(&User{}).Where("uuid = ?", uuid).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}

	h.GetMe(c)
}

func (h *Handler) ChangePassword(c *gin.Context) {
	uuidAny, _ := c.Get("userUuid")
	uuid, _ := uuidAny.(string)

	var req changePasswordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var u User
	if err := h.DB.Where("uuid = ?", uuid).First(&u).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	if !utils.CheckPassword(u.PasswordHash, req.OldPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "old password incorrect"})
		return
	}

	hash, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "hash password failed"})
		return
	}

	if err := h.DB.Model(&User{}).Where("uuid = ?", uuid).Update("password_hash", hash).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update password failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}
