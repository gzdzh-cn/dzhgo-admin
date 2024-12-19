package sys

import (
	"dzhgo/internal/dao"
	"dzhgo/internal/model"
	"dzhgo/internal/service"
	"github.com/gzdzh-cn/dzhcore"

	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	service.RegisterBaseSysAddonsTypeService(NewsBaseSysAddonsTypeService())
}

type sBaseSysAddonsTypeService struct {
	*dzhcore.Service
}

func NewsBaseSysAddonsTypeService() *sBaseSysAddonsTypeService {
	return &sBaseSysAddonsTypeService{
		&dzhcore.Service{
			Dao:   &dao.BaseSysAddonsType,
			Model: model.NewBaseSysAddonsType(),
			PageQueryOp: &dzhcore.QueryOp{
				KeyWordField: []string{"name", "remark"},
				AddOrderby:   g.MapStrStr{"`base_sys_addonsType`.`orderNum`": "ASC", "`base_sys_addonsType`.`createTime`": "DESC"},
			},
		},
	}
}
