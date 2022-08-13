package origins

import (
	"context"

	"github.com/mqdvi/kargo-excellerate-2/backend/helper"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
)

type Data struct {
	Items interface{} `json:"items,omitempty"`
}

func (svc *originService) GetOriginsByName(ctx context.Context) (*Data, error) {
	origins, err := svc.repo.GetOriginsByName(ctx, svc.db)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"OriginRepository", "GetOrigins"}).
			Msg(err.Error())
		return nil, err
	}
	// data := helper.Data{}
	data = data.WithItems(origins)
	//TODO bawah belum
	// resp := helper.PaginationResponse{
	// 	Data: data,
	// }

	return &resp, nil
}
