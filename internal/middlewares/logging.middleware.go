package middlewares

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {
		log.Default().Printf("Request received: %s | address: %v\n", req.URL.Path, req.RemoteAddr)
		next(wr, req)
	}
}
