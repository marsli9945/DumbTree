package middleware

import (
	"log"
	"net/http"
	"time"
)

func Times(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println("start")
		start := time.Now()

		handler.ServeHTTP(writer, request)

		end := time.Since(start)

		log.Println(end)
	})
}
