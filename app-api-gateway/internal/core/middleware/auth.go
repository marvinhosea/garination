package middleware

import "github.com/gin-gonic/gin"

func (m middlewareManager) Auth(ctx *gin.Context) {
	//TODO implement me
	ctx.Next()
}
