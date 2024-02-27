package hello

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/drewfugate/neverl8/model"
)

var ErrNotExist = errors.New("hello does not exist") // Define the error variable

type PostgresRepo struct {
	DB *sql.DB
}

func (r *PostgresRepo) GetHello(ctx context.Context) (model.Hello, error) {
	var value string
	err := r.DB.QueryRowContext(ctx, "SELECT value FROM hello_table LIMIT 1").Scan(&value)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Hello{}, ErrNotExist
		}
		return model.Hello{}, fmt.Errorf("get hello: %w", err)
	}

	return model.Hello{value}, nil
}
