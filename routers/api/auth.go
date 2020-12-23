package api

import (
	"fmt"
	"learngo/models"
	"learngo/pkg/e"
	"learngo/pkg/util"
	"learngo/pkg/logging"
	"net/http"

	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"username"`
	Password string `valid:"password"`
}

// GetAuth 获取用户
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	// a := auth{Username: username, Password: password}

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	if username != "" && password != "" {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				fmt.Println(err)
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	}else{
		logging.Info("auth", "用户名密码必填")
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
