package middleware

import "net/http"

func CORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		res.Header().Set("Access-Control-Allow-Credentials", "true")
		res.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
		res.Header().Set("Access-Control-Max-Age", "1209600")
		res.Header().Set("Content-Type", "application/json")

		handler.ServeHTTP(res, req)
	})
}
