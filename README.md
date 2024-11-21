
# Movie Festival Backend

This is the backend service for a Movie Festival application, implemented in Golang. It provides APIs for managing movies, tracking viewership, user authentication, and a voting system.

## Features

- **User Authentication**:
  - Register, login, and logout functionality using JWT tokens.
- **Admin Features**:
  - Create, update, and manage movies.
  - Get the most viewed and most voted movies.
- **User Features**:
  - Vote and unvote movies.
  - List all voted movies.
- **Public Features**:
  - View and search movies.
  - Track movie viewership.

## Project Structure

```
├── cmd
│   ├── migrate            # Migration management
│   └── server             # Application entry point
├── internal
│   ├── db                 # Database connection and migrations
│   ├── handlers           # HTTP request handlers
│   ├── middleware         # Middleware for request handling
│   ├── models             # Database models
│   ├── repositories       # Data access logic
│   └── services           # Business logic and utilities
├── migrations             # SQL migration files
├── pkg                    # Utility packages
└── routes                 # Application routing
```

## Setup Instructions

### Prerequisites

- Golang (>= 1.18)
- PostgreSQL
- `golang-migrate` for database migrations

### Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd movie-festival-backend
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Configure environment variables in a `.env` file:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=yourpassword
   DB_NAME=movie_festival
   DB_SSLMODE=disable
   SECRET_KEY=your_secret_key
   ```

4. Run database migrations:
   ```bash
   go run cmd/migrate/main.go
   ```

5. Start the server:
   ```bash
   go run cmd/server/main.go
   ```

## API Endpoints

### Public Endpoints

- **Register**: `POST /auth/register`
- **Login**: `POST /auth/login`
- **List Movies**: `GET /movies`
- **Search Movies**: `GET /movies/search?query=<query>`
- **Track Viewership**: `POST /movies/{id}/view`

### Protected User Endpoints

- **Vote Movie**: `POST /movies/{movie_id}/vote`
- **Unvote Movie**: `DELETE /movies/{movie_id}/unvote`
- **List User Votes**: `GET /user/votes`

### Protected Admin Endpoints

- **Create Movie**: `POST /admin/movies`
- **Update Movie**: `PUT /admin/movies/{id}`
- **Most Viewed Movies**: `GET /admin/movies/most-viewed`
- **Most Voted Movies**: `GET /admin/movies/most-voted`

## Testing with Postman

1. Import the provided Postman collection.
2. Use Postman to test all endpoints.
3. Set the `{{base_url}}` variable in Postman to your server's address (e.g., `http://localhost:8080`).
4. To test admin and user protected path set `token` and `admin_token` based on token that returned when use `Login` endpoint 

## Example Data

Use the following JSON structure to add movies through Postman:

```json
{
  "movie": {
    "title": "The Dark Knight",
    "description": "A gritty tale of Batman facing off against the Joker.",
    "duration": 152,
    "artists": "Christian Bale, Heath Ledger",
    "watch_url": "http://example.com/the_dark_knight"
  },
  "genres": ["Action", "Thriller"]
}
```

## Contributing

Feel free to fork this repository and submit pull requests. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License.
