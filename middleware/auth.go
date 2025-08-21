package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/labstack/echo/v4"
)

func InitClerk() {
	clerk.SetKey("sk_test_zj0PpSZ82t3zfgtWZSzwnugdsnu1DnkPVK0mz8cVqs")
}

func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Printf("Processing request to: %s\n", c.Request().URL.Path)

		// Get the Authorization header
		authHeader := c.Request().Header.Get("Authorization")
		fmt.Printf("Authorization header: %s\n", authHeader)

		if authHeader == "" {
			fmt.Println("No Authorization header found")
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}

		// Parse the Bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			fmt.Printf("Invalid Authorization header format: %v\n", authHeader)
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}
		token := parts[1]
		fmt.Printf("Token received: %s...\n", token[:10])

		// Verify the session claims
		claims, ok := clerk.SessionClaimsFromContext(c.Request().Context())
		if !ok {
			fmt.Println("No session claims found")
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}

		// Verify the user
		usr, err := user.Get(c.Request().Context(), claims.Subject)
		if err != nil {
			fmt.Printf("User verification failed: %v\n", err)
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}

		if usr.Banned {
			fmt.Println("User is banned")
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}

		// Store user in context
		c.Set("user", usr)
		fmt.Println("Auth successful")

		return next(c)
	}
}
