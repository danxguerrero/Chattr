package main

import (
	"github.com/danxguerrero/chattr/handlers"
	"github.com/danxguerrero/chattr/middleware"
	"github.com/danxguerrero/chattr/templates"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize Clerk
	middleware.InitClerk()

	e := echo.New()

	// Serve static files
	e.Static("/static", "static")

	// Public routes
	e.GET("/", func(c echo.Context) error {
		return templates.Index().Render(c.Request().Context(), c.Response().Writer)
	})

	// SSO callback route
	e.GET("/sso-callback", func(c echo.Context) error {
		return c.Redirect(302, "/chat")
	})

	// Protected routes
	e.GET("/chat", func(c echo.Context) error {
		return templates.Chat().Render(c.Request().Context(), c.Response().Writer)
	}, middleware.RequireAuth)

	// WebSocket route (protected)
	e.GET("/ws", handlers.HandleWebSocket, middleware.RequireAuth)

	// Start WebSocker broadcaster in a goroutine
	go handlers.BroadcastMessages()

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
