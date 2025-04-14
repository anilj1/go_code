package controller

import "net/http"

type PostController interface {
	GetPosts(resp http.ResponseWriter, req *http.Request)
	AddPosts(resp http.ResponseWriter, req *http.Request)
}
