package handlers

import (
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	auth0ManagedEndpoint(w)
}

func Login(w http.ResponseWriter, r *http.Request) {
	auth0ManagedEndpoint(w)
}

func auth0ManagedEndpoint(w http.ResponseWriter) {
	http.Error(w, "Authentication and registration are managed by Auth0. Obtain an access token from your Auth0 application.", http.StatusGone)
}
