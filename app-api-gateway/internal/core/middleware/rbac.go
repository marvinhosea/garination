package middleware

import "github.com/gin-gonic/gin"

func (m middlewareManager) Permission(ctx *gin.Context) {
	//TODO implement me
	ctx.Next()
}
