package jwt

import (
	"admin/pkg/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
import "admin/pkg/code"

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var rescode int
		var data interface{}
		var token, token2 string
		var err error

		rescode = code.SUCCESS
		token = c.GetHeader("token")
		token2 = c.GetHeader("token2")
		if token == "" && token2 == "" {
			rescode = code.ERROR_AUTH
		} else {
			_, err = utils.ParseToken(token)

			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					rescode = code.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					rescode = code.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}

			if rescode == code.ERROR_AUTH_CHECK_TOKEN_TIMEOUT {
				claims, err := utils.ParseToken(token2)

				if err != nil {
					rescode = code.ERROR_AUTH_CHECK_TOKEN_FAIL
				} else {
					token, err = utils.GenerateToken(claims.Username, claims.Password, time.Hour*24*10, false)
					if err != nil {
						rescode = code.ERROR_AUTH_CHECK_TOKEN_FAIL
					} else {
						c.SetCookie("token", token, 0, "/", "localhost", false, true)
						rescode = code.SUCCESS
					}

				}
			}
		}

		if rescode != code.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": rescode,
				"msg":  code.GetMsg(rescode),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
