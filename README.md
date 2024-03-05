# neverl8

**neverl8** is a streamlined Go application utilizing the Go-Chi router, GORM ORM, showcasing calendar operations within a PostgreSQL environment. Designed for simplicity and efficiency, this project will serve as the essential scheduler for Rise8.

## Prerequisites

To get the most out of **neverl8**, please ensure you have the following installed on your system:

- **Go programming language** (version 1.16 or higher) for the backend logic.
- **PostgreSQL database** for data persistence.
- **Git** for version control and collaboration.

## Installation

Follow these simple steps to get **neverl8** up and running on your machine:

1. **Clone the repository** to your local machine:
   ```bash
   git clone https://github.com/rise8-us/neverl8.git
   ```
2. **Navigate to the project directory**.
3. **Launch the application**:
   ```bash
   go run main.go
   ```
   Congratulations! The application should now be accessible at http://localhost:8080.

## Development Setup

**neverl8** leverages pre-commit for managing git hooks, aiding in maintaining high code quality and consistency across contributions.

### Setting Up Pre-commit

To integrate pre-commit into your development workflow:

1. **Install pre-commit** on your local machine. Refer to the [official installation guide](https://pre-commit.com/#install) for detailed instructions.
2. **Clone this repository** and navigate to the project root.
3. **Activate pre-commit** by running:
   ```bash
   pre-commit install
   ```

With these steps completed, pre-commit hooks will automatically execute on every git commit, enhancing your code quality checks.

### Using golangci-lint

**neverl8** also incorporates `golangci-lint` for enforcing Go best practices and code styles. To use `golangci-lint` in your development process:

1. **Install golangci-lint** on your local machine. You can follow the [official golangci-lint installation instructions](https://golangci-lint.run/usage/install/).
2. Once installed, you can run `golangci-lint run` in the project directory to analyze your code.

## Testing

**neverl8** embraces testing as a fundamental part of the development process. To run the unit tests and ensure your setup is correctly configured:
```bash
go test ./...
```
This command triggers all the unit tests within the project, verifying the integrity and functionality of your code.
