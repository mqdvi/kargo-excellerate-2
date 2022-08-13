package origins

import "github.com/go-playground/validator/v10"

type originController struct {
	originSvc OriginServiceInterface
	validator  *validator.Validate
}

func NewController(originSvc OriginServiceInterface) *originController {
	return &originController{originSvc: originSvc}
}
