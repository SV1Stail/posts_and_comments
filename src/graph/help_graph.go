package graph

import (
	"context"
	"fmt"

	"github.com/SV1Stail/test_ozon/graph/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

func commentsForPost(conn *pgxpool.Conn, ctx context.Context, postID string) ([]*model.Comment, error) {
	rows, err := conn.Query(ctx, "SELECT id, text, post_id, parent_id, user_id FROM comments WHERE post_id=$1", postID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch comments: %w", err)
	}
	defer rows.Close()

	var comments []*model.Comment
	commentsMap := make(map[string]*model.Comment)
	for rows.Next() {
		var commentAuthor model.User
		var comment model.Comment
		err := rows.Scan(&comment.ID, &comment.Text, &comment.PostID, &comment.ParentID, &commentAuthor.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan comment: %w", err)
		}
		comment.Author = &commentAuthor
		comment.Children = []*model.Comment{}
		commentsMap[comment.ID] = &comment
		if comment.ParentID == nil {
			comments = append(comments, &comment)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	for _, comment := range commentsMap {
		if comment.ParentID != nil {
			if parentComment, ok := commentsMap[*comment.ParentID]; ok {
				parentComment.Children = append(parentComment.Children, comment)
			}
		}
	}
	return comments, nil
}

// вывести комменты к постam
func commentsForPosts(conn *pgxpool.Conn, ctx context.Context) ([]*model.Comment, error) {
	rows, err := conn.Query(ctx, "SELECT id, text, post_id, parent_id, user_id FROM comments")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch comments: %w", err)
	}
	defer rows.Close()

	var comments []*model.Comment
	commentMap := make(map[string]*model.Comment)
	for rows.Next() {
		var comment model.Comment
		var commentAuthor model.User

		err := rows.Scan(&comment.ID, &comment.Text, &comment.PostID, &comment.ParentID, &commentAuthor.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan comment: %w", err)
		}

		comment.Author = &commentAuthor
		comment.Children = []*model.Comment{}
		commentMap[comment.ID] = &comment
		if comment.ParentID == nil {
			comments = append(comments, &comment)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	for _, comment := range commentMap {
		if comment.ParentID != nil {
			if parentComment, ok := commentMap[*comment.ParentID]; ok {
				parentComment.Children = append(parentComment.Children, comment)
			}
		}
	}
	return comments, nil

}

func commentsGetChildrenForComment(conn *pgxpool.Conn, ctx context.Context, postID string, limit int, offset int) ([]*model.Comment, error) {
	rows, err := conn.Query(ctx, `SELECT id, text, post_id, parent_id, user_id
	FROM comments
	WHERE post_id=$1 
	ORDER BY parent_id DESC
	LIMIT $2 OFFSET $3`, postID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch comments: %w", err)
	}
	defer rows.Close()
	commentsMap := make(map[string]*model.Comment)
	var comments []*model.Comment
	for rows.Next() {
		var child model.Comment
		var author model.User
		err = rows.Scan(&child.ID, &child.Text, &child.PostID, &child.ParentID, &author.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan comment: %w", err)
		}
		child.Author = &author
		child.Children = []*model.Comment{}
		commentsMap[child.ID] = &child
		if child.ParentID == nil {
			comments = append(comments, &child)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	for _, comment := range commentsMap {
		if comment.ParentID != nil {
			if parentComment, ok := commentsMap[*comment.ParentID]; ok {
				parentComment.Children = append(parentComment.Children, comment)
			}
		}
	}
	return comments, nil
}
