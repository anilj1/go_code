package repository

import (
	"cleanarch/entity"
	"database/sql"
	"log"
	"os"
	"testing"
)

func setupTestDB() {
	_ = os.RemoveAll("./posts_test.db")
	db, err := sql.Open("sqlite3", "./posts_test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
        create table posts (id integer not null primary key, title text, txt text);
        delete from posts;
        `
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}
}

func TestSave(t *testing.T) {
	setupTestDB()
	repo := NewSQLiteRepository()

	post := &entity.Post{Id: 1, Title: "Test Title", Text: "Test Text"}
	savedPost, err := repo.Save(post)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if savedPost.Id != post.Id || savedPost.Title != post.Title || savedPost.Text != post.Text {
		t.Errorf("Expected post %+v, got %+v", post, savedPost)
	}
}

func TestFindAll(t *testing.T) {
	setupTestDB()
	repo := NewSQLiteRepository()

	post1 := &entity.Post{Id: 1, Title: "Title 1", Text: "Text 1"}
	post2 := &entity.Post{Id: 2, Title: "Title 2", Text: "Text 2"}
	_, _ = repo.Save(post1)
	_, _ = repo.Save(post2)

	posts, err := repo.FindAll()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(posts) != 2 {
		t.Errorf("Expected 2 posts, got %d", len(posts))
	}
}

func TestDelete(t *testing.T) {
	setupTestDB()
	repo := NewSQLiteRepository()

	post := &entity.Post{Id: 1, Title: "Title", Text: "Text"}
	_, _ = repo.Save(post)

	err := repo.Delete(post)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	posts, err := repo.FindAll()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(posts) != 0 {
		t.Errorf("Expected 0 posts, got %d", len(posts))
	}
}
