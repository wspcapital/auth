package auth

import "net/http"

// Provider define Provider interface
type Provider interface {
	GetName() string
	ConfigAuth(*Auth)

	Login(*http.Request, http.ResponseWriter, *Claims)
	Logout(*http.Request, http.ResponseWriter, *Claims)
	Register(*http.Request, http.ResponseWriter, *Claims)
	Callback(*http.Request, http.ResponseWriter, *Claims)
	ServeHTTP(*http.Request, http.ResponseWriter, *Claims)
}
