package handler

import (
	"encoding/json"
	"net/http"

	appCtx "taudience.com/number-service/appcontext"
	"taudience.com/number-service/filter"
)

/*
Middleware ...
*/
func Middleware(app appCtx.ResponseBuilder, next func(app appCtx.ResponseBuilder, w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next(app, w, r)
	})
}

/*
HandleNumbers ...
*/
func HandleNumbers(app appCtx.ResponseBuilder, w http.ResponseWriter, r *http.Request) {
	urls := filter.GetValidURLs(r.URL.Query()["u"])
	numbers := app.Query(urls)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"numbers": numbers})
}
