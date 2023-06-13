package middleware

import (
	"garination.com/gateway/internal/core/common"
	"github.com/gin-gonic/gin"
)

var unauthorised = common.HttpReponse{
	Success: false,
	Message: "Unauthorised",
	Data:    nil,
	Errors:  []string{"Access to resource is unauthorised"},
}

func (m middlewareManager) Auth(ctx *gin.Context) {
	// get token from header (Authorization: Bearer <token>)
	token := ctx.GetHeader("Authorization")
	if token == "" {
		unauthorised.Errors = append(unauthorised.Errors, "Token is required")
		ctx.AbortWithStatusJSON(401, unauthorised)
		return
	}

	// validate token type
	if len(token) < 7 || token[:7] != "Bearer " {
		unauthorised.Errors = append(unauthorised.Errors, "Invalid token type")
		ctx.AbortWithStatusJSON(401, unauthorised)
		return
	}

	// validate token
	claims, err := m.casdoorClient.ParseJwtToken(token[7:])
	if err != nil {
		unauthorised.Errors = append(unauthorised.Errors, "Invalid token value")
		ctx.AbortWithStatusJSON(401, unauthorised)
		return
	}

	ctx.Set("user-name", claims.Name)
	ctx.Set("user-email", claims.Email)
	ctx.Set("user-uuid", claims.Id)

	ctx.Next()
}
