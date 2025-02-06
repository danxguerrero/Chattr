package main

import (
	"github.com/danxguerrero/chattr/handlers"
	"github.com/danxguerrero/chattr/templates"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Route to serve static files
	e.Static("/static", "static")

	// Route for chat interface
	e.GET("/", func(c echo.Context) error {
		return templates.Index().Render(c.Request().Context(), c.Response().Writer)
	})

	// WebSocket Route
	e.GET("/ws", handlers.HandleWebSocket)

	// Start WebSocker broadcaster in a goroutine
	go handlers.BroadcastMessages()

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}