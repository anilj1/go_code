package service

import (
	"cleanarch/entity"
	"cleanarch/repository"
	"errors"
	"math/rand"
)

var (
	repo repository.PostRepository
)

type service struct{}

func NewPostService(repository repository.PostRepository) PostService {
	repo = repository
	return &service{}
}

func (s service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("the post is empty")
		return err
	}

	if post.Title == "" {
		err := errors.New("the post title is empty")
		return err
	}
	return nil
}

func (s service) Create(post *entity.Post) (*entity.Post, error) {
	post.Id = rand.Int63()
	return repo.Save(post)
}

func (s service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}
