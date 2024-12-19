package service

import (
	"dzhgo/addons/space/dao"
	"dzhgo/addons/space/model"
	"github.com/gzdzh-cn/dzhcore"
)

type SpaceTypeService struct {
	*dzhcore.Service
}

func NewSpaceTypeService() *SpaceTypeService {
	return &SpaceTypeService{
		&dzhcore.Service{
			Dao:   &dao.AddonsSpaceType,
			Model: model.NewSpaceType(),
		},

		// Service: dzhcore.NewService(model.NewSpaceType()),
	}
}
