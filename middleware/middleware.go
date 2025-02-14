package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/auth0/go-jwt-middleware/v2/validator"
	beegoCtx "github.com/beego/beego/v2/server/web/context"
)

// Secret key (must match auth.go)
var secretKey = []byte("supersecret")

// Expected issuer and audience (must match auth.go)
const (
	issuer   = "my-app"
	audience = "my-audience"
)

// Global JWT validator instance
var jwtValidator *validator.Validator

// Extracts the standard Go context from Beego’s context
func GetStandardContext(ctx *beegoCtx.Context) context.Context {
	if ctx != nil && ctx.Request != nil {
		return ctx.Request.Context()
	}
	return context.Background() // Fallback if no request exists
}

// Key function for JWT validation
func keyFunc(ctx context.Context) (interface{}, error) {
	return secretKey, nil
}

// Initialize JWT validator inside init()
func init() {
	var err error
	jwtValidator, err = validator.New(
		keyFunc,
		validator.HS256,
		issuer,
		[]string{audience}, // ✅ Audience is now provided
	)

	if err != nil {
		fmt.Println("❌ Failed to initialize JWT validator:", err)
		panic(err) // Exit if initialization fails
	}

	fmt.Println("✅ JWT Validator initialized successfully")
}

// Middleware function to validate JWT in Beego
func JWTMiddleware(ctx *beegoCtx.Context) {
	stdCtx := GetStandardContext(ctx) // Convert Beego context to standard Go context

	token := ctx.Input.Header("Authorization") // Get token from request header
	if token == "" {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.Body([]byte("⛔ Missing token"))
		return
	}

	// Remove "Bearer " prefix if present
	token = strings.TrimPrefix(token, "Bearer ")

	fmt.Println("🔍 Received Token:", token) // Debugging line

	// Ensure validator is properly initialized before use
	if jwtValidator == nil {
		ctx.Output.SetStatus(http.StatusInternalServerError)
		ctx.Output.Body([]byte("❌ JWT Validator not initialized"))
		return
	}

	// Validate the token
	_, err := jwtValidator.ValidateToken(stdCtx, token)
	if err != nil {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.Body([]byte("⛔ Invalid or expired token: " + err.Error()))
		return
	}

	// If validation succeeds, allow request
	ctx.Output.SetStatus(http.StatusOK)
	ctx.Output.Body([]byte("🔓 Token Validated!"))
}
