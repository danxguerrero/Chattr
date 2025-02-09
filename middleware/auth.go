package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/labstack/echo/v4"
)

// Create a new Clerk Client
var clerkClient, err = clerk.NewClient(os.Getenv("CLERK_SECRET_KEY"))

// Validates the Clerk session
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionToken := c.Request().Header.Get("Authorization")
		if sessionToken == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "No session token"})
		}

		session, err := clerkClient.Sessions().Verify(sessionToken, "")
		if err != nil || session.Status != "active" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid session"})
		}

		c.Set("userID", session.UserID)
		return next(c)
	}
}

func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}

		if parts[1] == "" {
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}

		return next(c)
	}
}