package middleware

import (
	"errors"
	"garination.com/gateway/internal/core/common"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func (m middlewareManager) RecoverHandler(ctx *gin.Context, err any) {
	// Log the error
	log.Println(err, "recovering from panic")
	// Set the response status code

	ctx.Writer.WriteHeader(http.StatusInternalServerError)

	// Send the "something went wrong" response
	ctx.JSON(http.StatusInternalServerError, common.HttpReponse{
		Message: "Something went wrong, and it's not your fault.",
	})
}

func (m middlewareManager) Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			// Check for a broken connection, as it is not really a
			// condition that warrants a panic stack trace.
			var brokenPipe bool
			if ne, ok := err.(*net.OpError); ok {
				var se *os.SyscallError
				if errors.As(ne, &se) {
					seStr := strings.ToLower(se.Error())
					if strings.Contains(seStr, "broken pipe") ||
						strings.Contains(seStr, "connection reset by peer") {
						brokenPipe = true
					}
				}
			}
			if brokenPipe {
				// If the connection is dead, we can't write a status to it.
				c.Error(err.(error)) //nolint: errcheck
				c.Abort()
			} else {
				c.Writer.WriteHeader(http.StatusInternalServerError)

				// Send the "something went wrong" response
				c.JSON(http.StatusInternalServerError, common.HttpReponse{
					Message: "Something went wrong, and it's not your fault.",
				})
			}
		}
	}()
	c.Next()
}
