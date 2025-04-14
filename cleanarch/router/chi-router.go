package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type chiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() Router {
	return &chiRouter{}
}

func (mr chiRouter) GET(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (mr chiRouter) POST(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	chiDispatcher.Post(uri, f)
}

func (mr chiRouter) SERVE(port string) {
	fmt.Printf("Chi http server running on port: %v", port)
	http.ListenAndServe(port, chiDispatcher)
}
