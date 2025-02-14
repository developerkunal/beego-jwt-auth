package main

import (
	"log"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/developerkunal/beego-jwt-auth/auth"
	"github.com/developerkunal/beego-jwt-auth/middleware" // Adjust to your GitHub repo path
)

func main() {
	// Define routes
	beego.Router("/", &HomeController{}, "get:Home")
	beego.Router("/login", &AuthController{}, "get:Login") // Changed POST to GET
	beego.Router("/protected", &AuthController{}, "get:Protected")

	// Apply JWT middleware to protect the route
	beego.InsertFilter("/protected", beego.BeforeRouter, middleware.JWTMiddleware)

	// Start the Beego server
	log.Println("🚀 Server running on http://localhost:8080")
	beego.Run()
}

// HomeController handles public routes
type HomeController struct {
	beego.Controller
}

// Public route (no authentication needed)
func (c *HomeController) Home() {
	c.Ctx.Output.SetStatus(200)
	c.Ctx.Output.Body([]byte("🏠 Welcome to the Beego API"))
}

// AuthController handles authentication routes
type AuthController struct {
	beego.Controller
}

// 🔹 **Updated: Login route now supports GET**
func (c *AuthController) Login() {
	token, err := auth.GenerateJWT("kunal") // Generate JWT for user "kunal"
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Ctx.Output.Body([]byte("Error generating token"))
		return
	}

	// Return JSON response
	c.Ctx.Output.Header("Content-Type", "application/json")
	c.Ctx.Output.Body([]byte(`{"token": "` + token + `"}`))
}

// Protected route (requires authentication)
func (c *AuthController) Protected() {
	c.Ctx.Output.SetStatus(200)
	c.Ctx.Output.Body([]byte("🔐 Access granted to protected route!"))
}
