package middleware

import (
	"garination.com/gateway/internal/platform/casdoor"
	"garination.com/gateway/internal/platform/prom"
	"github.com/gin-gonic/gin"
)

type MiddlewareManager interface {
	Cors(ctx *gin.Context)
	Auth(ctx *gin.Context)
	Permission(ctx *gin.Context)
	Recover(c *gin.Context)
	RecoverHandler(c *gin.Context, err any)
	RequestMetrics(ctx *gin.Context)
	CarHits(ctx *gin.Context)
	SparePartHits(ctx *gin.Context)
}

type middlewareManager struct {
	promMetrics   prom.Metrics
	casdoorClient casdoor.CasdoorClient
}

func NewMiddlewareManager(promMetrics prom.Metrics, casdoorClient casdoor.CasdoorClient) MiddlewareManager {
	return &middlewareManager{
		promMetrics:   promMetrics,
		casdoorClient: casdoorClient,
	}
}
