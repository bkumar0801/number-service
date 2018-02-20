package main

import (
	"flag"
	"log"
	"net/http"

	appCtx "taudience.com/number-service/appcontext"
	h "taudience.com/number-service/handler"
)

func main() {
	listenAddr := flag.String("http.addr", ":8070", "http listen address")
	flag.Parse()

	appContext := &appCtx.AppContext{}

	http.Handle("/numbers", h.Middleware(appContext, h.HandleNumbers))
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}
