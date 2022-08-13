package driver

import (
	"context"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"

	"github.com/microcosm-cc/bluemonday"
)

func (svc *driverService) Create(ctx context.Context, payload *models.RequestCreateDriver) (*models.ResponseCreateDriver, error) {

	svc.sanitizePayload(payload)

	result, err := svc.repo.CreateDriver(ctx, svc.db, payload)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (svc *driverService) sanitizePayload(payload *models.RequestCreateDriver) {
	p := bluemonday.StrictPolicy()

	payload.Name = p.Sanitize(payload.Name)
	payload.PhoneNumber = p.Sanitize(payload.PhoneNumber)
	payload.IDCard = p.Sanitize(payload.IDCard)
	payload.DriverLicense = p.Sanitize(payload.DriverLicense)
}
