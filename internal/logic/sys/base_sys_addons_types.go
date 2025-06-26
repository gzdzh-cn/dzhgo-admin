package sys

import (
	"context"
	"dzhgo/internal/dao"
	"dzhgo/internal/model"
	"dzhgo/internal/service"

	"github.com/gzdzh-cn/dzhcore"

	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	service.RegisterBaseSysAddonsTypesService(NewsBaseSysAddonsTypesService())
}

type sBaseSysAddonsTypesService struct {
	*dzhcore.Service
}

func NewsBaseSysAddonsTypesService() *sBaseSysAddonsTypesService {
	return &sBaseSysAddonsTypesService{
		&dzhcore.Service{
			Dao:   &dao.BaseSysAddonsTypes,
			Model: model.NewBaseSysAddonsTypes(),
			PageQueryOp: &dzhcore.QueryOp{
				KeyWordField: []string{"name", "remark"},
				AddOrderby:   g.MapStrStr{"`base_sys_addons_types`.`orderNum`": "ASC", "`base_sys_addons_types`.`createTime`": "DESC"},
			},
		},
	}
}

func (s *sBaseSysAddonsTypesService) Show(ctx context.Context) (data interface{}, err error) {

	return
}
