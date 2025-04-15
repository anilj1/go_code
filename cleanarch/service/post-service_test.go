package service

import (
	"cleanarch/entity"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func (mock *MockRepository) Delete(post *entity.Post) error {
	return nil
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)
	err := testService.Validate(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "the post is empty", err.Error())
}

func TestValidateEmptyPostTitle(t *testing.T) {
	testService := NewPostService(nil)
	post := entity.Post{Id: rand.Int63(), Title: "", Text: ""}
	err := testService.Validate(&post)
	assert.NotNil(t, err)
	assert.Equal(t, "the post title is empty", err.Error())
}

func TestValidateValidPost(t *testing.T) {
	testService := NewPostService(nil)
	post := entity.Post{Id: rand.Int63(), Title: "Title", Text: "Text"}
	err := testService.Validate(&post)
	assert.Nil(t, err)
}

func TestFindAll(t *testing.T) {
	mockRepo := MockRepository{}

	// Setup the expectations.
	var postId int64 = 1
	post := entity.Post{Id: postId, Title: "Title", Text: "Txt"}
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(&mockRepo)
	result, _ := testService.FindAll()

	// Assert the result - behavioral.
	mockRepo.AssertExpectations(t)

	// Assert the result - data.
	assert.Equal(t, postId, result[0].Id)
	assert.Equal(t, "Title", result[0].Title)
	assert.Equal(t, "Txt", result[0].Text)
}

func TestCreate(t *testing.T) {
	mockRepo := MockRepository{}

	// Setup the expectations.
	post := entity.Post{Title: "Title", Text: "Txt"}
	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(&mockRepo)
	result, err := testService.Create(&post)

	// Assert the result - behavioral.
	mockRepo.AssertExpectations(t)

	// Assert the result - data.
	assert.Nil(t, err)
	assert.NotEmpty(t, result.Id)
	assert.Equal(t, "Title", result.Title)
	assert.Equal(t, "Txt", result.Text)
}
