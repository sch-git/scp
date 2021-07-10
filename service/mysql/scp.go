package mysql

import (
	"context"
	"service/models"
)

type SCPDAO struct {
	db *DB
}

func (d *SCPDAO) Add(ctx context.Context, record *models.SCP) (err error) {
	return nil
}
