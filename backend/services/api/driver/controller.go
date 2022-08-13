package driver

import "github.com/go-playground/validator/v10"

type driverController struct {
	driverSvc DriverServiceInterface
	validator *validator.Validate
}

func NewController(driverSvc DriverServiceInterface, validator *validator.Validate) *driverController {
	return &driverController{
		driverSvc: driverSvc,
		validator: validator,
	}
}
