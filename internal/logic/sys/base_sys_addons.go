package sys

import (
	"context"
	v1 "dzhgo/internal/api/admin_v1"
	"dzhgo/internal/common"
	"dzhgo/internal/dao"
	"dzhgo/internal/model"
	"dzhgo/internal/service"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gzdzh-cn/dzhcore"
)

func init() {
	service.RegisterBaseSysAddonsService(NewsBaseSysAddonsService())
}

type sBaseSysAddonsService struct {
	*dzhcore.Service
}

func NewsBaseSysAddonsService() *sBaseSysAddonsService {
	return &sBaseSysAddonsService{
		&dzhcore.Service{
			Dao:   &dao.BaseSysAddons,
			Model: model.NewBaseSysAddons(),
			PageQueryOp: &dzhcore.QueryOp{
				KeyWordField: []string{"name", "remark"},
				AddOrderby:   g.MapStrStr{"`base_sys_addons`.`orderNum`": "ASC", "`base_sys_addons`.`createTime`": "DESC"},
				Where: func(ctx context.Context) []g.Array {
					var (
						admin = common.GetAdmin(ctx)
						r     = g.RequestFromCtx(ctx)
						rmap  = r.GetMap()
					)
					condition := []g.Array{{"typeId = ?", rmap["typeId"], rmap["typeId"]}, {"c.IsInstall = ?", true, rmap["type"]}}

					roleIds := garray.NewStrArrayFrom(admin.RoleIds)
					if !roleIds.Contains("1") {
						condition = append(condition, g.Array{"c.isShow", 1})
					}

					return condition
				},
				Select: "base_sys_addons.*,b.name as typeName,c.isInstall,c.isShow",
				Join: []*dzhcore.JoinOp{
					{
						Model:     model.NewBaseSysAddonsTypes(),
						Alias:     "b",
						Type:      "LeftJoin",
						Condition: "`base_sys_addons`.`typeId` = `b`.`id`",
					},
					{
						Model:     model.NewBaseSysMenu(),
						Alias:     "c",
						Type:      "LeftJoin",
						Condition: "`base_sys_addons`.`menuId` = `c`.`id`",
					},
				},
			},
		},
	}
}

// 安装卸载插件
func (s *sBaseSysAddonsService) InstallUpdateStatus(ctx context.Context, req *v1.InstallUpdateStatusReq) (data interface{}, err error) {

	condition := g.Map{
		"id = ?": g.Slice{req.Id},
	}
	updateData := g.Map{
		"isInstall": req.Active,
	}

	_, err = dao.BaseSysMenu.Ctx(ctx).Where(condition).Data(updateData).Update()
	if err != nil {
		g.Log().Error(ctx, err.Error())
		err = gerror.New("操作失败")
		return
	}
	return
}

// 上下架插件
func (s *sBaseSysAddonsService) LineUpdateStatus(ctx context.Context, req *v1.LineUpdateStatusReq) (data interface{}, err error) {

	condition := g.Map{
		"id = ?": g.Slice{req.Id},
	}
	updateData := g.Map{
		"isShow": req.Active,
	}
	_, err = dao.BaseSysMenu.Ctx(ctx).Where(condition).Data(updateData).Update()
	if err != nil {
		g.Log().Error(ctx, err.Error())
		err = gerror.New("操作失败")
		return
	}

	return
}
