package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
)

func (m middlewareManager) RequestMetrics(ctx *gin.Context) {
	if ctx.IsAborted() {
		return
	}
	// get the request path
	path := ctx.Request.URL.Path
	method := ctx.Request.Method

	label := fmt.Sprintf("%s %s", method, path)
	timer := prometheus.NewTimer(m.promMetrics.ResponseDuration.WithLabelValues(label))

	// increment the total number of requests
	m.promMetrics.RequestCount.WithLabelValues(label).Inc()

	defer func() {
		timer.ObserveDuration()
		status := ctx.Writer.Status()
		m.promMetrics.ResponseStatus.WithLabelValues(label, strconv.Itoa(status)).Inc()
	}()

	ctx.Next()
}
