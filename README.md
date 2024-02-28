# Go PostgreSQL CRUD API

This is a simple CRUD (Create, Read, Update, Delete) API example in Go for interacting with a PostgreSQL database. The application provides HTTP endpoints for managing users in the database.

## Functionality

- Create a new user
- Retrieve user information by ID
- Update user information
- Delete a user by ID

## Technologies

- Programming Language: Go
- Web Framework: Standard net/http library
- Database: PostgreSQL
- Docker: For containerizing the application and database
- Docker Compose: For managing the multi-container application

## How to Use

1. Install Docker and Docker Compose if not already installed on your machine.
2. Clone the repository to your machine: `git clone https://github.com/say8hi/go-crud-api.git`
3. Navigate to the project directory: `cd go-postgresql-crud`
4. Create a .env file and specify the PostgreSQL database connection parameters (see .env.example for an example).
5. Start the application using the command `docker-compose up`.
6. The application will be accessible at `http://localhost:8080`.

## Project Structure

```bash
go-postgresql-crud/
|- main.go
|- handlers/
| |- users.go
|- database/
| |- postgres.go
|- models/
| |- user.go
|- Dockerfile
|- docker-compose.yml
|- .env
```

## Contribution

If you find a bug or want to suggest an improvement, please create an Issue or Pull Request in this repository.