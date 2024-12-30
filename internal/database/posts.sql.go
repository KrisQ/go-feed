// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: posts.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const addPost = `-- name: AddPost :one
INSERT INTO posts (
    id, 
    created_at, 
    updated_at, 
    title, 
    url, 
    description, 
    published_at, 
    feed_id
) VALUES (
    $1,                         
    CURRENT_TIMESTAMP,          
    CURRENT_TIMESTAMP,         
    $2,                         
    $3,                        
    $4,                       
    $5,                      
    $6                
) RETURNING id, created_at, updated_at, title, url, description, published_at, feed_id
`

type AddPostParams struct {
	ID          uuid.UUID
	Title       string
	Url         string
	Description sql.NullString
	PublishedAt sql.NullTime
	FeedID      uuid.UUID
}

func (q *Queries) AddPost(ctx context.Context, arg AddPostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, addPost,
		arg.ID,
		arg.Title,
		arg.Url,
		arg.Description,
		arg.PublishedAt,
		arg.FeedID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Url,
		&i.Description,
		&i.PublishedAt,
		&i.FeedID,
	)
	return i, err
}

const getPostsForUser = `-- name: GetPostsForUser :many
SELECT posts.title
FROM posts
INNER JOIN feeds ON posts.feed_id = feeds.id
INNER JOIN feed_follows ON feeds.id = feed_follows.feed_id
INNER JOIN users ON feed_follows.user_id = users.id
WHERE users.id = $1
ORDER BY posts.published_at DESC
LIMIT $2
`

type GetPostsForUserParams struct {
	ID    uuid.UUID
	Limit int32
}

func (q *Queries) GetPostsForUser(ctx context.Context, arg GetPostsForUserParams) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getPostsForUser, arg.ID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var title string
		if err := rows.Scan(&title); err != nil {
			return nil, err
		}
		items = append(items, title)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
