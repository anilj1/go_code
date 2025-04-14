package repository

import (
	"cleanarch/entity"
	"log"

	"cloud.google.com/go/firestore"
	"context"
)

const (
	projectId      string = "cleanarch-64ee2"
	collectionName string = "posts"
)

type repo struct{}

// NewFirestoreRepository
func NewFirestoreRepository() PostRepository {
	return &repo{}
}

func (r repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatal("Failed to create a Firestore client: %v", err)
		return nil, err
	}

	defer client.Close()
	_, _, err1 := client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"Id":    post.Id,
		"Title": post.Title,
		"Text":  post.Text,
	})
	if err1 != nil {
		log.Fatal("Failed adding a new post: %v", err)
		return nil, err
	}

	return post, nil
}

func (r repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatal("Failed to create a Firestore client: %v", err)
		return nil, err
	}
	defer client.Close()

	var posts []entity.Post
	iterator := client.Collection(collectionName).Documents(ctx)
	for {
		doc, _ := iterator.Next()
		if doc != nil {
			post := entity.Post{
				Id:    doc.Data()["Id"].(int64),
				Title: doc.Data()["Title"].(string),
				Text:  doc.Data()["Text"].(string),
			}
			posts = append(posts, post)
		} else {
			break
		}
	}
	return posts, nil
}
