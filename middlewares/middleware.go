package middleware

import (
	"github.com/dragfire/gocha/logger"
	"net/http"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info.Println(r.UserAgent())
		logger.Info.Println(r.Referer())
		logger.Info.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
