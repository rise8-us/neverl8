neverl8

This is a simple Go application built with Go-Chi router and GORM ORM, demonstrating basic CRUD operations with PostgreSQL.

Prerequisites

Before you begin, ensure you have the following installed on your machine:

Go programming language (version 1.16 or higher),
PostgreSQL database,
Git

Installation:

Clone the repository:

git clone https://github.com/drewfugate/neverl8.git

Navigate to the project directory:

cd src

go run main.go

The application should now be running on http://localhost:8080.

## Development Setup

This project uses [pre-commit](https://pre-commit.com/) to manage git hooks. Pre-commit hooks help ensure code quality and consistency.

### Installing Pre-commit

To set up pre-commit on your local machine, follow these steps:

1. Install pre-commit. See the [official installation instructions](https://pre-commit.com/#install).
2. Clone the repository and navigate into it.
3. Run `pre-commit install` to set up the git hook scripts.

Now pre-commit will run automatically on `git commit`!

Testing

To run unit tests, execute the following command:
go test
