package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type OAuthServer struct {
	fingerprints map[string]bool
}

func NewOAuthServer() *OAuthServer {
	return &OAuthServer{
		fingerprints: make(map[string]bool),
	}
}

func (s *OAuthServer) AuthorizeHandler(c *gin.Context) {
	userAgent := c.GetHeader("User-Agent")
	ip := c.ClientIP()

	// Создаем "fingerprint" устройства
	fingerprint := ip + userAgent

	if _, ok := s.fingerprints[fingerprint]; ok {
		// Устройство уже авторизовано
		c.JSON(http.StatusOK, gin.H{"status": "already authorized"})
		return
	}

	// Здесь вы бы выполнили процедуру аутентификации, но в данном случае мы просто авторизуем каждое новое устройство
	s.fingerprints[fingerprint] = true

	c.JSON(http.StatusOK, gin.H{"status": "authorized"})
}

func main() {
	oauthServer := NewOAuthServer()

	r := gin.Default()
	r.GET("/authorize", oauthServer.AuthorizeHandler)
	r.Run(":8080")
}
