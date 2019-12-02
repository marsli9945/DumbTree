package router

import "net/http"

type middleware func(handler http.Handler) http.Handler

type Middle interface {
	before()
	next()
}

type Router struct {
	middlewareChain []middleware
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Use(m middleware) {
	r.middlewareChain = append(r.middlewareChain, m)
}

func (r *Router) Add(router string, f func(writer http.ResponseWriter, request *http.Request)) {
	var mergedHandle = (func(f func(writer http.ResponseWriter, request *http.Request)) http.Handler {
		return http.HandlerFunc(f)
	})(f)

	for i := len(r.middlewareChain) - 1; i >= 0; i-- {
		mergedHandle = r.middlewareChain[i](mergedHandle)
	}

	http.Handle(router, mergedHandle)
}
