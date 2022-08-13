package origins

import (
	"context"

	"github.com/mqdvi/kargo-excellerate-2/backend/helper"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
	// "github.com/mqdvi/kargo-excellerate-2/backend/models"
)

func (svc *originService) GetOrigins(ctx context.Context, filter *models.OriginFilter) (*[]models.Origin, error) {
	origins, err := svc.repo.GetOrigins(ctx, svc.db, filter)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"OriginRepository", "GetOrigins"}).
			Msg(err.Error())
		return nil, err
	}

	return &origins, nil
}
