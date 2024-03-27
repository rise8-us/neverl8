package testutil

import (
	"context"
	"fmt"
	"log"
	"time"

	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TestDB struct {
	DB        *gorm.DB
	container testcontainers.Container
}

func SetupTestDB() *TestDB {
	ctx := context.Background()
	container, port := createContainer(ctx)

	// Connect to the database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		"localhost", "user", "password", "testDb", port.Port(), "disable")
	db, err := gorm.Open(pg.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %s", err)
	}

	// Migrate db
	m, err := migrate.New(
		"file://../../../../backend/db/migrations",
		fmt.Sprintf(("postgres://%s:%s@%s:%s/%s?sslmode=%s"), "user", "password", "localhost", port.Port(), "testDb", "disable"))
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			log.Println("No migration to run")
		} else {
			log.Fatal(err)
		}
	}

	return &TestDB{DB: db, container: container}
}

func (tdb *TestDB) TearDown() {
	// remove test container
	_ = tdb.container.Terminate(context.Background())
}

func createContainer(ctx context.Context) (testcontainers.Container, nat.Port) {
	dbName := "testDb"
	dbUser := "user"
	dbPassword := "password"
	var timeout = 5 * time.Second

	postgresContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:15.2-alpine"),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(timeout)),
	)

	if err != nil {
		log.Fatalf("Failed to start container: %s", err)
	}

	// Get the container's mapped port
	mappedPort, err := postgresContainer.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatalf("Failed to get mapped port: %s", err)
	}

	return postgresContainer, mappedPort
}
