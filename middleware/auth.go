package middleware

import (
	"MarvelousBlog-Backend/common"
	"MarvelousBlog-Backend/entity"
	"MarvelousBlog-Backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.JSON(http.StatusUnauthorized, entity.ResponseObject{
				Code:    common.UNAUTHORIZED,
				Message: common.Message[common.UNAUTHORIZED],
			})
			c.Abort()
			return
		}

		tokens := strings.Split(auth, " ")

		if len(tokens) != 2 || tokens[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, entity.ResponseObject{
				Code:    common.TOKEN_WRONG_TYPE,
				Message: common.Message[common.TOKEN_WRONG_TYPE],
			})
			c.Abort()
			return
		}

		auth = strings.Fields(auth)[1]
		claims, code := utils.ParseToken(auth)
		if code != 200 {
			c.JSON(http.StatusUnauthorized, entity.ResponseObject{
				Code:    code,
				Message: common.Message[code],
			})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
