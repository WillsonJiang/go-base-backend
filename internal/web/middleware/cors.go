package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Authorization, authorization, Referer, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
		ctx.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, PATCH, DELETE")

		method := ctx.Request.Method
		if method == http.MethodOptions {
			ctx.JSON(http.StatusOK, "Options Request!")
		}
		ctx.Next()
	}
}
