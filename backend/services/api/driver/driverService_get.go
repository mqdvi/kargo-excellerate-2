package driver

import (
	"context"
	"github.com/mqdvi/kargo-excellerate-2/backend/helper"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
)

func (svc *driverService) GetDriver(ctx context.Context, filter *models.GetDriverFilter) (*models.ResponseGetDrivers, error) {

	drivers, err := svc.repo.GetDriver(ctx, svc.db, filter)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"DriverRepository", "GetDriver"}).
			Msg(err.Error())
		return nil, err
	}
	maskingDriver := make([]models.Driver, 0)
	for _, driver := range drivers {
		if driver.Status == "1" {
			driver.Status = "Active"
		} else {
			driver.Status = "Inactive"
		}
		maskingDriver = append(maskingDriver, driver)
	}
	resp := models.ResponseGetDrivers{
		Data: models.DriverPayloadData{
			Items: maskingDriver,
		},
	}

	return &resp, nil
}

func (svc *driverService) GetDriverByID(ctx context.Context, driverID int64) (*models.Driver, error) {

	driver, err := svc.repo.GetDriverByID(ctx, svc.db, driverID)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"DriverRepository", "GetDriver"}).
			Msg(err.Error())
		return nil, err
	}

	return driver, nil
}