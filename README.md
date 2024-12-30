# go-feed

A command-line RSS feed aggregator written in Go that allows users to manage and follow RSS feeds.

## Features

- User management (register, login)
- Feed management (add, list, follow, unfollow)
- Post aggregation from RSS feeds
- PostgreSQL database for persistent storage

## Prerequisites

- Go 1.23.4 or higher
- PostgreSQL database
- Database URL in the format: `postgresql://username:password@localhost:5432/dbname?sslmode=disable`

## Installation

1. Clone the repository:
```
git clone https://github.com/KrisQ/go-feed.git
```
2. Navigate to the project directory:
```
cd go-feed
```
3. Install dependencies:
```
go mod download
```
4. Create a `.gatorconfig.json` file in your home directory with your database configuration:

```json
{
"db_url": "postgresql://username:password@localhost:5432/dbname?sslmode=disable"
}
```

## Available Commands

- `register <username>` - Register a new user
- `login <username>` - Login as an existing user
- `reset` - Reset (delete) all users from the database
- `users` - List all registered users
- `addfeed <name> <url>` - Add a new RSS feed (requires login)
- `feeds` - List all available feeds
- `follow <url>` - Follow a feed (requires login)
- `following` - List feeds you're following (requires login)
- `unfollow <url>` - Unfollow a feed (requires login)
- `agg <duration>` - Start the feed aggregator (e.g., `agg 1m` for 1-minute intervals)
- `browse <limit>` - Browse your feed posts with a limit

### Usage

From the root of the project:

```
go run . register user1
go run . login user1
go run . addfeed myfeed https://example.com/rss
go run . follow https://example.com/rss
go run . agg 1m
go run . browse 10
```
