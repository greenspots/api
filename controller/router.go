package controller

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Router struct {
	*httprouter.Router
}

func (r *Router) Get(path string, handler http.Handler) {
	r.GET(path, wrapHandler(handler))
}

func (r *Router) Post(path string, handler http.Handler) {
	r.POST(path, wrapHandler(handler))
}

func NewRouter() *Router {
	return &Router{httprouter.New()}
}

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		r = r.WithContext(context.WithValue(r.Context(), "params", ps))
		h.ServeHTTP(w, r)
	}
}

func RecoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic:q %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
