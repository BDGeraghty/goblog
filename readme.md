# GoBlog (Gator)

GoBlog (aka **Gator**) is a command-line RSS aggregation and blog management tool written in Go. It allows you to manage RSS feeds, follow feeds from other users, and aggregate blog posts from multiple sources.

---

## Prerequisites

Before running GoBlog, ensure you have the following installed on your system:

- **Go** (version 1.18 or newer recommended)  
  [Install Go](https://golang.org/doc/install)
- **PostgreSQL** (version 12 or newer recommended)  
  [Install PostgreSQL](https://www.postgresql.org/download/)

---

## Setup

### 1. Clone the Repository

```sh
git clone https://github.com/bdgeraghty/GoBlog.git
cd GoBlog
```

### 2. Install Dependencies

```sh
go mod download
```

### 3. Set Up PostgreSQL Database

Create a PostgreSQL database for the application:

```sql
CREATE DATABASE gator;
```

### 4. Configure the Application

Create a `.gatorconfig.json` file in your project root directory:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

**Configuration Fields:**
- `db_url`: Your PostgreSQL connection string
  - Replace `username` and `password` with your PostgreSQL credentials
  - Replace `gator` with your database name if different
- `current_user_name`: Will be set automatically when you log in (leave empty initially)

### 5. Run Database Migrations

Navigate to the SQL schema directory and apply migrations:

```sh
cd sql/schema
goose postgres "your-connection-string-here" up
cd ../..
```

Replace `"your-connection-string-here"` with the same connection string from your `.gatorconfig.json` file.

---

## Running GoBlog

Build and run the application:

```sh
go build -o gator
./goblog <command> [arguments]
```

Or run directly:

```sh
go run . <command> [arguments]
```

---

## Available Commands

### User Management

- **`register <username>`** - Register a new user
  ```sh
  ./goblog register john_doe
  ```

- **`login <username>`** - Log in as an existing user
  ```sh
  ./goblog login john_doe
  ```

- **`users`** - List all registered users
  ```sh
  ./goblog users
  ```

- **`reset`** - Reset user data (removes current user from config)
  ```sh
  ./goblog reset
  ```

### Feed Management

- **`addfeed <name> <url>`** - Add a new RSS feed (requires login)
  ```sh
  ./goblog addfeed "Tech Blog" "https://example.com/feed.xml"
  ```

- **`feeds`** - List all available feeds
  ```sh
  ./goblog feeds
  ```

- **`follow <feed-id>`** - Follow an existing feed (requires login)
  ```sh
  ./goblog follow <feed-uuid>
  ```

- **`unfollow <feed-id>`** - Unfollow a feed (requires login)
  ```sh
  ./goblog unfollow <feed-uuid>
  ```

- **`following <username>`** - List feeds followed by a user (requires login)
  ```sh
  ./goblog following john_doe
  ```

### Content Aggregation

- **`agg`** - Aggregate and fetch new posts from all feeds
  ```sh
  ./goblog agg
  ```

---

## Example Workflow

1. **Register a new user:**
   ```sh
   ./goblog register alice
   ```

2. **Add an RSS feed:**
   ```sh
   ./goblog addfeed "My Blog" "https://myblog.com/feed.xml"
   ```

3. **List all feeds:**
   ```sh
   ./goblog feeds
   ```

4. **Follow a feed (using feed ID from the feeds list):**
   ```sh
   ./goblog follow <feed-id>
   ```

5. **Aggregate new posts:**
   ```sh
   ./goblog agg
   ```

6. **See what feeds you're following:**
   ```sh
   ./goblog following alice
   ```

---

## Notes

- Commands marked with **(requires login)** need you to be logged in first using the `login` command
- Feed IDs are UUIDs that you can get from the `feeds` command
- The `.gatorconfig.json` file stores your current login state
- Make sure PostgreSQL is running before using the application

---

## Troubleshooting

**Database Connection Issues:**
- Verify PostgreSQL is running
- Check your connection string in `.gatorconfig.json`
- Ensure the database exists and migrations have been applied

**Command Not Found:**
- Make sure you've built the application: `go build -o gator`
- Or use `go run . <command>` instead

**Permission Errors:**
- Ensure your PostgreSQL user has the necessary permissions
- Check that the `.gatorconfig.json` file is readable/writable

---

## Development

To contribute to this project:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests (if available)
5. Submit a pull request

---

## License

[Add your license information here]