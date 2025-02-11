package middleware

import (
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/labstack/echo/v4"
)

var clerkClient clerk.Client

func InitClerk() {
	// Initialize Clerk client with your secret key
	clerk.SetKey("sk_test_zj0PpSZ82t3zfgtWZSzwnugdsnu1DnkPVK0mz8cVqs")
}

// RequireAuth adapts Clerk's middleware for Echo
func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.WrapHandler(clerkhttp.WithHeaderAuthorization()(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c := r.Context()
			claims, ok := clerk.SessionClaimsFromContext(c)
			if !ok {
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				return
			}

			// Get the Echo context
			ec := r.Context().Value("echo").(echo.Context)

			// Verify the user
			usr, err := user.Get(r.Context(), claims.Subject)
			if err != nil {
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				return
			}

			if usr.Banned {
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				return
			}

			// Store user in context
			ec.Set("user", usr)

			next(ec)
		}),
	))
}
