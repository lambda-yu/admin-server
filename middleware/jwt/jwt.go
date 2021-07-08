package jwt

import (
	"admin/pkg/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)
import "admin/pkg/code"

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var rescode int
		var data interface{}
		var token string
		var err error

		rescode = code.SUCCESS
		token = c.Query("token")
		if token == "" {
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
