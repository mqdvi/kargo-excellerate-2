package driver

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
)

func (repo *driverRepository) CreateDriver(ctx context.Context, db *sqlx.DB, payload *models.RequestCreateDriver) (*models.ResponseCreateDriver, error) {
	query, args, _ := sq.Insert(repo.GetTableName()).
		SetMap(sq.Eq{
			"id":             ,
			"name":           payload.Name,
			"phone_number":   payload.PhoneNumber,
			"created_at":     sq.Expr("NOW()"),
			"status":         true,
			"id_card":        payload.IDCard,
			"driver_license": payload.DriverLicense,
		}).ToSql()

	_, err := db.ExecContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	response := &models.ResponseCreateDriver{
		Data: "ok",
	}
	return response, nil
}
