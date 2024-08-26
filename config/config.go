package config

import (
	"context"
	"fmt"
	"time"
	"net/http"

	supa "github.com/nedpals/supabase-go"

	"github.com/kerneels94/reports/functions"

	"github.com/labstack/echo/v4"
)

// 
func SetCookie(w http.ResponseWriter, accessToken string) {
	cookie := &http.Cookie{
		Name:     "user_access_token",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   30 * 60, // 30 minutes
	}
	http.SetCookie(w, cookie)
}

func CookieLogout(w http.ResponseWriter) {
	// Set the cookie with an expiration date in the past to delete it
	cookie := &http.Cookie{
		Name:     "user_access_token",
		Value:    "",
		Path:    "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(-24 * time.Hour), // 24 hours ago
		MaxAge:   -1, // Also sets the cookie to expire immediately
	}
	http.SetCookie(w, cookie)
}

func IsCookieValid(r *http.Request, c echo.Context) bool {
	supabaseClient, err := functions.CreateSupabaseClient()

	if err != nil {
		fmt.Println("Error connecting:", err)
		return false
	}

	cookie, err := r.Cookie("user_access_token")
	if err != nil {
		fmt.Println("Error retrieving cookie:", err)
		return false
	}

	ctx := context.Background()

	user, err := supabaseClient.Auth.User(ctx, cookie.Value)

	if err != nil {
		fmt.Println("Error retrieving user:", err)
		return false
	}

	if user == nil {
		fmt.Println("User not found")
		return false
	}

	return true
}

func GetUserIdFromCookie(r *http.Request, c echo.Context, spbase *supa.Client) string {
	cookie, err := r.Cookie("user_access_token")
	if err != nil {
		fmt.Println("Error retrieving cookie:", err)
		return ""
	}

	ctx := context.Background()

	user, err := spbase.Auth.User(ctx, cookie.Value)

	if err != nil {
		fmt.Println("Error retrieving user:", err)
		return ""
	}

	if user == nil {
		fmt.Println("User not found")
		return ""
	}

	return user.ID
}
