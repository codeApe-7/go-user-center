package main

import (
	"log"
	"github.com/gin-gonic/gin"

	"github.com/codeApe-7/go-user-center/internal/config"
	"github.com/codeApe-7/go-user-center/internal/db"
	"github.com/codeApe-7/go-user-center/internal/middleware"
	"github.com/codeApe-7/go-user-center/internal/modules/auth"
	"github.com/codeApe-7/go-user-center/internal/modules/user"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	gdb, err := db.Connect(cfg.DB.DSN)
	if err != nil {
		log.Fatal(err)
	}

	if err := gdb.AutoMigrate(&user.User{}); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	// very simple CORS for learning
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	authH := &auth.Handler{DB: gdb, JWTSecret: cfg.JWT.Secret, Issuer: cfg.JWT.Issuer, AccessTTL: cfg.JWT.AccessTTL}
	userH := &user.Handler{DB: gdb}

	api := r.Group("/api")
	{
		api.POST("/auth/register", authH.Register)
		api.POST("/auth/login", authH.Login)

		me := api.Group("/me")
		me.Use(middleware.AuthRequired(middleware.AuthConfig{JWTSecret: cfg.JWT.Secret}))
		me.GET("", userH.GetMe)
		me.PUT("", userH.UpdateMe)
		me.PUT("/password", userH.ChangePassword)
	}

	log.Printf("server listening on %s", cfg.Server.Addr)
	if err := r.Run(cfg.Server.Addr); err != nil {
		log.Fatal(err)
	}
}
