package driver

import (
	"context"
	"github.com/mqdvi/kargo-excellerate-2/backend/helper"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
	"github.com/rs/zerolog/log"

	"github.com/microcosm-cc/bluemonday"
)

func (svc *driverService) Deactivate(ctx context.Context, truckID int64) error {
	err := svc.repo.DeactivateTruck(ctx, svc.db, truckID)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"TruckRepository", "DeactivateTruck"}).
			Msg(err.Error())
		return err
	}

	return nil
}