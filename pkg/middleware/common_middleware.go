package middleware

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CorsMiddleware(whiteListedUrls map[string]bool) mux.MiddlewareFunc{
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(rw http.ResponseWriter, r *http.Request) {

				if r.Method != http.MethodOptions{
					next.ServeHTTP(rw, r)
					return
				}

				requestOriginUrl := r.Host
				log.Printf("INFO CorsMiddleware: received request from %s %v", requestOriginUrl, whiteListedUrls[requestOriginUrl])
				if !whiteListedUrls[requestOriginUrl]{
					return
				}

				rw.Header().Set("Access-Control-Allow-Origin", requestOriginUrl)
				rw.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE, PATCH")
				rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization")
				rw.Header().Set("Access-Control-Allow-Credentials", "true")

				rw.Write([]byte("GAS"))
			})
	}
}