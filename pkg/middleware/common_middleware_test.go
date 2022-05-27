package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCorsMiddeware(t *testing.T) {
	whiteListedUrls := make(map[string]bool)
	whiteListedUrls["google.com"] = true
	whiteListedUrls["localhost:8080"] = true

	tests := []struct {
		name       string
		handler    http.HandlerFunc
		senderHost string
		method     string
	}{
		{
			name: "CorsMiddleware Case Successful",
			senderHost: "google.com",
			method: "OPTIONS",
			handler: func(w http.ResponseWriter, r *http.Request) {
				corsVal := r.Header.Get("Access-Control-Allow-Origin")
				if corsVal != "google.com"{
					t.Errorf("CorsMiddleware error, got %v, want %v", corsVal, "google.com")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.method, "http://testing", nil)
			r.Host = tt.senderHost

			CorsMiddleware(whiteListedUrls).Middleware(tt.handler).ServeHTTP(httptest.NewRecorder(), r)
		})
	}
}