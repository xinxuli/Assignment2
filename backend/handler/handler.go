package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"Assignment2/backend/models"
	"errors"
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

}

func login(UserLoginRequest) (UserLoginResponse, error) {


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