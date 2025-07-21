package store

import (
	"errors"
)

type User struct {
	Username string
	Password string
	NickName string
}

var Users = []User{
	{
		Username: "kevin",
		Password: "$2a$10$MsnkkPX8AXuHGW8J780Gx.VJb9.XKEPtVp0BbfEzXmO.woeaG38DS",
		NickName: "KevinZonda",
	},
	{
		Username: "long",
		Password: "$2a$10$ho6S906ZjYOQYih2dWXBBOfW062gDEj0cFJcib6SL5r5zqOeSdTNC",
		NickName: "Bobby",
	},
	{
		Username: "deng",
		Password: "$2a$10$dOTp3s2LpCyQR06nDDi7sOZZXb54zlUlQbXrgHga5dVtrPDmLWTXO",
		NickName: "ShuaiZhi",
	},
}

func GetUserByUsername(username string) (*User, error) {
    for _, user := range Users {
        if user.Username == username {
            return &user, nil
        }
    }
    return nil, errors.New("user not found")
}