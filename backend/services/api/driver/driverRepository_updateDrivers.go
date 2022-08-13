package driver

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
)

func (repo *driverRepository) UpdateDriver(ctx context.Context, db *sqlx.DB, payload *models.RequestCreateDriver, driverID int64) (*models.ResponseCreateDriver, error) {
	_, err := repo.GetDriverByID(ctx, db, driverID)
	
	if err != nil {
		response, err := repo.CreateDriver(ctx, db, payload)
		return response, err
	}

	query, args, _ := sq.Update(repo.GetTableName()).
		SetMap(sq.Eq{
			"name":           payload.Name,
			"phone_number":   payload.PhoneNumber,
			"id_card":        payload.IDCard,
			"driver_license": payload.DriverLicense,
		}).
		Where(sq.Eq{"id": driverID}).
		ToSql()

	_, err = db.ExecContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	response := &models.ResponseCreateDriver{
		Data: "ok",
	}
	return response, nil
}
