package sys

import (
	"dzhgo/addons/dict/dao"
	"dzhgo/addons/dict/model"
	"dzhgo/addons/dict/service"
	"github.com/gzdzh-cn/dzhcore"
)

func init() {
	service.RegisterDictTypeService(&sDictTypeService{})
}

type sDictTypeService struct {
	*dzhcore.Service
}

func NewsDictTypeService() *sDictTypeService {
	return &sDictTypeService{
		Service: &dzhcore.Service{
			Dao:   &dao.AddonsDictType,
			Model: model.NewDictType(),
			ListQueryOp: &dzhcore.QueryOp{
				KeyWordField: []string{"name"},
			},
		},
	}
}
