package sys

import (
	"dzhgo/internal/dao"
	"dzhgo/internal/model"
	"dzhgo/internal/service"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gzdzh-cn/dzhcore"
)

func init() {
	service.RegisterBaseSysDepartmentService(NewsBaseSysDepartmentService())
}

type sBaseSysDepartmentService struct {
	*dzhcore.Service
}

// NewsBaseSysDepartmentService 创建一个sBaseSysDepartmentService实例
func NewsBaseSysDepartmentService() *sBaseSysDepartmentService {
	return &sBaseSysDepartmentService{
		Service: &dzhcore.Service{
			Dao:         &dao.BaseSysDepartment,
			Model:       model.NewBaseSysDepartment(),
			ListQueryOp: &dzhcore.QueryOp{},
		},
	}
}

// GetByRoleIds 获取部门
func (s *sBaseSysDepartmentService) GetByRoleIds(roleIds []string) (res []uint) {
	var (
		result gdb.Result
	)
	// 如果roleIds不为空
	if len(roleIds) > 0 {
		// 如果是超级管理员，则返回所有部门
		if garray.NewStrArrayFrom(roleIds).Contains("1") {
			result, _ = s.Dao.Ctx(ctx).Fields("id").All()
			for _, v := range result {
				vmap := v.Map()
				if vmap["id"] != nil {
					res = append(res, gconv.Uint(vmap["id"]))
				}
			}
		} else {
			// 如果不是超级管理员，则返回角色所在部门
			result, _ = dao.BaseSysRoleDepartment.Ctx(ctx).Where("roleId IN (?)", roleIds).Fields("departmentId").All()
			for _, v := range result {
				vmap := v.Map()
				if vmap["departmentId"] != nil {
					res = append(res, gconv.Uint(vmap["departmentId"]))
				}
			}
		}

	}
	return
}

// Order 排序部门
func (s *sBaseSysDepartmentService) Order(ctx g.Ctx) (err error) {
	r := g.RequestFromCtx(ctx).GetMap()

	type item struct {
		Id       uint64  `json:"id"`
		ParentId *uint64 `json:"parentId,omitempty"`
		OrderNum int32   `json:"orderNum"`
	}

	var data *item

	for _, v := range r {
		err = gconv.Struct(v, &data)
		if err != nil {
			continue
		}
		s.Dao.Ctx(ctx).Where("id = ?", data.Id).Data(data).Update()
	}

	return
}
