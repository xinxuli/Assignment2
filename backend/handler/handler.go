package handler

import (
	"backend/models"
	"backend/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"strings"
)

// user login : username, password
// return token

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

func InitGin(g gin.IRouter) {
	g.POST("/login", UserLogin)
	g.GET("/user/info", GetUserInfo)

}

var jwtKey = []byte("your_secret_key")

func login(req UserLoginRequest) (string, error) {
	user, err := models.GetUserByUsername(req.Username)
	if err != nil {
		return "", errors.New("user not found")
	}
	if !utils.CheckPassword(req.Password, user.Password) {
		return "", errors.New("invalid password")
	}
	
	var jwtKey = []byte("your_secret_key")

	claims := jwt.MapClaims{
    "username": user.Username,
    "exp":      time.Now().Add(time.Hour * 24).Unix(), 
	}

	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenObj.SignedString(jwtKey)
	if err != nil {
    	return "", err
	}
	return token, nil
}

func UserLogin(c *gin.Context) {
	var req UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := login(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}
	c.JSON(http.StatusOK, models.UserLoginResponse{
		Token: token})
}



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

    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil || !token.Valid {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        return
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
        return
    }

    username, ok := claims["username"].(string)
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token payload"})
        return
    }

    user, err := models.GetUserByUsername(username)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"nick_name": user.NickName})
}