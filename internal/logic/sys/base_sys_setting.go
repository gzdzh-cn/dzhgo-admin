package sys

import (
	"dzhgo/internal/dao"
	"dzhgo/internal/model"
	"github.com/gzdzh-cn/dzhcore"
)

//func init() {
//	service.RegisterBaseSysSettingService(NewsBaseSysSettingService())
//}

type sBaseSysSettingService struct {
	*dzhcore.Service
}

func NewsBaseSysSettingService() *sBaseSysSettingService {
	return &sBaseSysSettingService{
		&dzhcore.Service{
			Dao:   &dao.BaseSysSetting,
			Model: model.NewBaseSysSetting(),
		},
	}
}
