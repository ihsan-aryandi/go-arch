package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var (
			start  = time.Now()
			method = request.Method
			url    = request.URL.String()
		)

		handler.ServeHTTP(writer, request)

		duration := time.Since(start).Milliseconds()
		log.Printf("[%s] %s -> %dms", method, url, duration)
	}
}
