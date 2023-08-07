package api

import (
	"net/http"
)

func HealCheck(w http.ResponseWriter, r *http.Request) {
	// Respond with a success message or HTTP status code.
	w.WriteHeader(http.StatusOK)
}
