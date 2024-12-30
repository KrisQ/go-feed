-- name: AddPost :one
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
) RETURNING *;

-- name: GetPostsForUser :many
SELECT posts.title
FROM posts
INNER JOIN feeds ON posts.feed_id = feeds.id
INNER JOIN feed_follows ON feeds.id = feed_follows.feed_id
INNER JOIN users ON feed_follows.user_id = users.id
WHERE users.id = $1
ORDER BY posts.published_at DESC
LIMIT $2;
