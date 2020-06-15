package utils

import (
	"net/http"
	"time"
)

const (
	cookieName    = "go_seccion"
	cookieExpires = 24 * 2 * time.Hour
)

func SetSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:    cookieName,
		Value:   "aa",
		Path:    "/",
		Expires: time.Now().Add(cookieExpires),
	}

	http.SetCookie(w, cookie)
}

func DeleteSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   cookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func getValCookie(r *http.Request) string {
	if cookie, err := r.Cookie(cookieName); err == nil {
		return cookie.Value
	}
	return ""
}

func IsAuthenticated(r *http.Request) bool {
	return getValCookie(r) != ""
}
