package graph

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/SV1Stail/test_ozon/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	posts    []*Post
	comments []*Comment
	mu       sync.Mutex
}

type User struct {
	ID string
}
type Post struct {
	ID            string
	Title         string
	Content       string
	AllowComments bool
	Comments      []*Comment
	Author        User
}

type Comment struct {
	ID       string
	Text     string
	PostID   string
	ParentID *string
	Children []*Comment
	Author   User
}

var postIDCounter = 1
var commentIDCounter = 1

// получить все посты
// можно докинуть пагинацию и асинхронность
func (r *Resolver) Posts(ctx context.Context) ([]*Post, error) {
	pool := db.GetPool()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, "SELECT id, title, content, allow_comments, user_id FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []*Post

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content,
			&post.AllowComments, &post.Author)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

// пост по id
func (r *Resolver) Post(ctx context.Context, id string) (*Post, error) {
	pool := db.GetPool()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	row := conn.QueryRow(ctx, "SELECT id, title, content, allow_comments, user_id FROM posts WHERE id=$1", id)

	post := &Post{}
	err = row.Scan(&post.ID, &post.Title, &post.Content,
		&post.AllowComments, &post.Author)
	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("post not found")
	} else if err != nil {
		return nil, err
	}
	return post, nil
}

// создать пост
func (r *Resolver) CreatePost(ctx context.Context, title string, content string, allowComments bool, authorId string) (*Post, error) {
	pool := db.GetPool()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	postID := uuid.New().String()
	_, err = conn.Exec(ctx, "INSERT INTO posts (id, title, content, allow_comments, user_id) VALUES ($1, $2, $3, $4, $5)",
		postID, title, content, allowComments, authorId)
	if err != nil {
		return nil, err
	}
	post := &Post{
		ID:            postID,
		Title:         title,
		Content:       content,
		AllowComments: allowComments,
		Author:        User{ID: authorId},
	}
	return post, nil
}

// создать коммент
func (r *Resolver) CreateComment(ctx context.Context, postId string, parentId *string, text string, authorId string) (*Comment, error) {
	if len(text) > 2000 {
		return nil, fmt.Errorf("comment text exceeds 2000 characters")
	}
	pool := db.GetPool()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	commentID := uuid.New().String()
	_, err = conn.Exec(ctx,
		"INSERT INTO comments (id, text, post_id, parent_id, user_id)VALUES ($1, $2, $3, $4, $5)",
		commentID, text, postId, parentId, authorId)
	if err != nil {
		return nil, err
	}
	comment := &Comment{
		ID:       commentID,
		Text:     text,
		PostID:   postId,
		ParentID: parentId,
		Children: []*Comment{},
		Author:   User{ID: authorId},
	}
	return comment, nil
}
func (r *Resolver) CommentsForPost(ctx context.Context, postId string) (<-chan *Comment, error) {
	pool := db.GetPool()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(ctx, "SELECT id, text, post_id, parent_id, user_id FROM comments WHERE post_id=$1", postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	resChan := make(chan *Comment)
	go func() {
		defer close(resChan) // Закрываем канал после завершения чтения данных
		for rows.Next() {
			var comment Comment
			err := rows.Scan(&comment.ID, &comment.Text, &comment.PostID, &comment.ParentID, &comment.Author.ID)
			if err != nil {
				log.Println("Error scanning row:", err)
				continue
			}
			go func(comm Comment) {

				resChan <- &comm

			}(comment)
		}
		if rows.Err() != nil {
			log.Println("Error during row iteration:", rows.Err())
		}
	}()

	return resChan, nil
}
