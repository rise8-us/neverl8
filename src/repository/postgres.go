package hello

import (
	"context"
	"errors"

	"github.com/drewfugate/neverl8/model"
	"github.com/jinzhu/gorm"
)

var ErrNotExist = errors.New("hello does not exist") // Define the error variable

type PostgresRepo struct {
	DB *gorm.DB
}

func (r *PostgresRepo) GetHello(ctx context.Context) (model.Hello, error) {
	return model.Hello{}, nil
}
