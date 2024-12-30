-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    title VARCHAR(255) NOT NULL,                   
    url VARCHAR(255) NOT NULL UNIQUE,                
    description TEXT,                               
    published_at TIMESTAMP,                        
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);
 

