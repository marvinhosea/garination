package middleware

import "github.com/gin-gonic/gin"

func (m middlewareManager) CarHits(ctx *gin.Context) {
	//TODO implement me
	ctx.Next()
}

func (m middlewareManager) SparePartHits(ctx *gin.Context) {
	//TODO implement me
	ctx.Next()
}
