package controller

import (
	"cleanarch/entity"
	"cleanarch/service"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	postService service.PostService
)

type controller struct{}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

func (ctrl *controller) GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	defer func() {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.New("error getting all posts data"))
	}()

	posts, err := postService.FindAll()
	if err != nil {
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func (ctrl *controller) AddPosts(resp http.ResponseWriter, req *http.Request) {
	var post entity.Post
	resp.Header().Set("Content-type", "application/json")
	defer func() {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.New("error getting all posts data"))
	}()

	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		return
	}

	// Validate the post object.
	err = postService.Validate(&post)
	if err != nil {
		return
	}

	result, err := postService.Create(&post)
	if err != nil {
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)
}
