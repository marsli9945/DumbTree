package middleware

import (
	"fmt"
	"net/http"
)

type Test3 struct {
}

func (t Test3) before() {
	panic("implement me")
}

func (t Test3) next() {
	panic("implement me")
}

func Test1(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("test1 before")
		handler.ServeHTTP(writer, request)
		fmt.Println("test1 next")
	})
}

func Test2(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("test2 before")
		handler.ServeHTTP(writer, request)
		fmt.Println("test2 next")
	})
}
