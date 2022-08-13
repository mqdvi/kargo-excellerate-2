package driver

import (
	"context"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
	"github.com/rs/zerolog/log"
)

func (svc *driverService) Update(ctx context.Context, payload *models.RequestCreateDriver, driverID int64) (*models.ResponseCreateDriver, error) {
	log.Debug().Interface("Req", payload).Msg("Start to run driver service create driver ")
	svc.sanitizePayload(payload)

	result, err := svc.repo.UpdateDriver(ctx, svc.db, payload, driverID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

