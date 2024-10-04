package middleware


import (
	"net/http"
	"log"
	"github.com/IndySeh/go-crud-application/pkg/logging"	
)

func LogIncomingRequestMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request Method: %s, Request URL: %s", r.Method, r.URL)
		logging.RequestLogger.Info("Request Method: " + r.Method + " Requested URL: " + r.URL.Path)
		next.ServeHTTP(w, r)
	})
}