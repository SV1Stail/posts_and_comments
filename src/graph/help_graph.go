package graph

import (
	"context"
	"fmt"

	"github.com/SV1Stail/test_ozon/graph/model"
	pgx "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// вывести комменты к посту
func commentsForPost(pool *pgxpool.Pool, ctx context.Context, postID string) ([]*model.Comment, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	commentsRows, err := conn.Query(ctx, "SELECT id, text, parent_id, user_id FROM comments WHERE post_id = $1", postID)
	if err != nil {
		return nil, err
	}
	defer commentsRows.Close()
	var comments []*model.Comment
	for commentsRows.Next() {
		var comment model.Comment
		var commentAuthor model.User
		err := commentsRows.Scan(&comment.ID, &comment.Text, &comment.ParentID, &commentAuthor.ID)
		if err != nil {
			return nil, err
		}
		if comment.ParentID != nil {
			continue
		}
		comment.Author = &commentAuthor
		comment.Children = make([]*model.Comment, 0)
		err = commentsGetChildrenForComment(pool, ctx, comment.ID, &comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	return comments, nil
}

// коммент по ID
func commentsGetCommentByID(pool *pgxpool.Pool, ctx context.Context, commentID string) (*model.Comment, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	row := conn.QueryRow(ctx, "SELECT id, text, post_id, parent_id, user_id FROM comments WHERE id=$1", commentID)
	comment := &model.Comment{}
	err = row.Scan(&comment.ID, &comment.Text, &comment.PostID, &comment.ParentID, &comment.Author.ID)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("comment not found")
	} else if err != nil {
		return nil, err
	}

	return comment, nil

}

func commentsGetChildrenForComment(pool *pgxpool.Pool, ctx context.Context, commentID string, comment *model.Comment) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	rows, err := conn.Query(ctx, "SELECT id, text, post_id, parent_id, user_id FROM comments WHERE parent_id = $1", commentID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var child model.Comment
		var author model.User
		err = rows.Scan(&child.ID, &child.Text, &child.PostID, &child.ParentID, &author.ID)
		if err != nil {
			return err
		}
		child.Author = &author
		comment.Children = append(comment.Children, &child)
		child.Children = make([]*model.Comment, 0)
		err = commentsGetChildrenForComment(pool, ctx, child.ID, &child)
		if err != nil {
			return err
		}

	}
	return nil
}
