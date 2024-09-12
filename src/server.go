package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/SV1Stail/test_ozon/db"
	"github.com/SV1Stail/test_ozon/graph"
)

const defaultPort = "8080"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db.Connect()
	defer db.ClosePool()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// type Resolver struct {
// 	posts    []*Post
// 	comments []*Comment
// }

// type User struct {
// 	ID string
// }
// type Post struct {
// 	ID            string
// 	Title         string
// 	Content       string
// 	AllowComments bool
// 	Comments      []*Comment
// 	Author        User
// }

// type Comment struct {
// 	ID       string
// 	Text     string
// 	PostID   string
// 	ParentID *string
// 	Children []*Comment
// 	Author   User
// }

// var postIDCounter = 1
// var commentIDCounter = 1

// func (r *Resolver) CommentsForPost(ctx context.Context, postId string) (<-chan *Comment, error) {
// 	pool := db.GetPool()
// 	conn, err := pool.Acquire(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer conn.Release()
// 	rows, err := conn.Query(ctx, "SELECT id, text, post_id, parent_id, user_id FROM comments WHERE post_id=$1", postId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	resChan := make(chan *Comment)
// 	go func() {
// 		defer close(resChan) // Закрываем канал после завершения чтения данных
// 		for rows.Next() {
// 			var comment Comment
// 			err := rows.Scan(&comment.ID, &comment.Text, &comment.PostID, &comment.ParentID, &comment.Author.ID)
// 			if err != nil {
// 				log.Println("Error scanning row:", err)
// 				continue
// 			}
// 			go func(comm Comment) {

// 				resChan <- &comm

// 			}(comment)
// 		}
// 		if rows.Err() != nil {
// 			log.Println("Error during row iteration:", rows.Err())
// 		}
// 	}()

// 	return resChan, nil
// }
