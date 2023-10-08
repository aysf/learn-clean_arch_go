package repository

import (
	"context"
	"fmt"
	"gofb/entity"
	"log"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type repo struct{}

// NewPostRepository
func NewPostRepository() PostRepository {
	return &repo{}
}

const (
	projectId      string = "belajar-go-5bfaa"
	collectionName string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()

	// client, err := firestore.NewClient(ctx, projectId)

	opt := option.WithCredentialsFile("/Users/ananto/Sandbox/goplay/learn_go-crash-course-1/cred/belajar-go-5bfaa-firebase-adminsdk-27esc-24224124e2.json")
	cfg := &firebase.Config{
		ProjectID: projectId,
	}
	app, _ := firebase.NewApp(context.Background(), cfg, opt)

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatal("failed to create a firestore client:", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.Id,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatal("failed adding a new post:", err)
		return nil, err
	}

	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()

	// client, err := firestore.NewClient(ctx, projectId)
	opt := option.WithCredentialsFile("/Users/ananto/Sandbox/goplay/learn_go-crash-course-1/cred/belajar-go-5bfaa-firebase-adminsdk-27esc-24224124e2.json")
	cfg := &firebase.Config{
		ProjectID: projectId,
	}
	app, _ := firebase.NewApp(context.Background(), cfg, opt)
	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatal("failed to create a firestore client:", err)
		return nil, err
	}
	defer client.Close()

	var posts []entity.Post
	col := client.Collection(collectionName)

	fmt.Println("col:", col)

	iterator := col.DocumentRefs(ctx)

	for {
		doc, err := iterator.Next()
		if err != nil && err.Error() == "no more items in iterator" {
			// log.Fatal("failed to iterate the list of posts:", err)
			// return nil, err
			break
		}
		d, _ := doc.Get(ctx)
		post := entity.Post{
			Id:    d.Data()["ID"].(int64),
			Title: "test title",
			Text:  "test text",
		}

		fmt.Println("post:", post)

		posts = append(posts, post)
	}

	return posts, nil

}
