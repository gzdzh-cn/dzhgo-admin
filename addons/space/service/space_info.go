package service

import (
	"dzhgo/addons/space/dao"
	"dzhgo/addons/space/model"
	"github.com/gzdzh-cn/dzhcore"
)

type SpaceInfoService struct {
	*dzhcore.Service
}

func NewSpaceInfoService() *SpaceInfoService {
	return &SpaceInfoService{
		&dzhcore.Service{
			Dao:   &dao.AddonsSpaceInfo,
			Model: model.NewSpaceInfo(),
		},

		// Service: dzhcore.NewService(model.NewSpaceInfo()),
	}
}
