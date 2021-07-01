package api

import (
	"admin/models"
	"admin/pkg/app"
	"admin/pkg/code"
	"admin/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
import "github.com/astaxie/beego/validation"

type authForm struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
	Remember bool   `valid:"Required"`
}

func Login(c *gin.Context) {
	var remember bool
	response := app.Response{C: c}
	valid := validation.Validation{}
	username := c.PostForm("username")
	password := c.PostForm("password")
	re := c.PostForm("remember")

	if re == "false" {
		remember = false
	} else {
		remember = true
	}

	a := authForm{username, password, remember}

	ok, _ := valid.Valid(&a)

	if !ok {
		app.MarkErrors(valid.Errors)
		response.JSONResponse(http.StatusBadRequest, code.INVALID_PARAMS, nil)
		return
	}

	user := models.GetUser(username, password)

	if user.ID < 0 {
		response.JSONResponse(http.StatusUnauthorized, code.ERROR_AUTH, nil)
		return
	}
	// create token
	token, err := utils.GenerateToken(username, password, time.Hour*24*10, true)
	token2, err := utils.GenerateToken(username, password, time.Hour*24*30, true)
	if err != nil {
		fmt.Println(err.Error())
		response.JSONResponse(http.StatusInternalServerError, code.ERROR_AUTH_TOKEN, nil)
		return
	}
	c.SetCookie("token", token, 0, "/", "localhost", false, true)
	c.SetCookie("token2", token2, 0, "/", "localhost", false, true)
	response.JSONResponse(http.StatusOK, code.SUCCESS, nil)
}
