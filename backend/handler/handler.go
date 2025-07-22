package handler

import (
	"backend/models"
	"backend/store"
	"backend/utils"
	"net/http"

	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// user login : username, password
// return token


func InitGin(g gin.IRouter) {
	g.POST("/user/login", UserLogin)
	g.GET("/user/info", GetUserInfo)

}

func UserLogin(c *gin.Context) {
	var req models.UserLoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := store.GetUserByUsername(req.Username)
	if err != nil || !utils.CheckPassword(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := utils.RandString(32) 
	tokenToUsername[token] = req.Username

	expirationTime := time.Now().Add(24 * time.Hour)
	c.SetCookie("token", token, int(expirationTime.Unix()), "/", "localhost", false, true)
	c.JSON(http.StatusOK, models.TokenResp{
		Token: token,
	})
}

var tokenToUsername map[string]string = make(map[string]string)

func GetUserInfo(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		return
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
		return
	}
	tokenStr := parts[1]

	username, ok := tokenToUsername[tokenStr]
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	user, err := store.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, models.NickNameResp{
		NickName: user.NickName,
	})
}
