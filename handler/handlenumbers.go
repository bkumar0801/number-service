package handler

import (
	"encoding/json"
	"net/http"

	appCtx "taudience.com/number-service/appcontext"
	"taudience.com/number-service/filter"
)

/*
Middleware ...This is a wrapper over handler to pass AppContext in handler func
*/
func Middleware(app appCtx.ResponseBuilder, next func(app appCtx.ResponseBuilder, w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next(app, w, r)
	})
}

/*
HandleNumbers ...This is a route handler
*/
func HandleNumbers(app appCtx.ResponseBuilder, w http.ResponseWriter, r *http.Request) {
	// Filter valid URLs
	urls := filter.GetValidURLs(r.URL.Query()["u"])
	// Fetch numbers from URLs, and return sorted list of numbers
	numbers := app.Query(urls)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Write Json response
	json.NewEncoder(w).Encode(map[string]interface{}{"numbers": numbers})
}
