package sys

import (
	"context"
	"dzhgo/internal/common"
	"dzhgo/internal/dao"
	"dzhgo/internal/model"
	"dzhgo/internal/service"
	"dzhgo/internal/types"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gzdzh-cn/dzhcore"
)

func init() {
	service.RegisterBaseSysMenuService(NewsBaseSysMenuService())
}

type sBaseSysMenuService struct {
	*dzhcore.Service
}

// NewsBaseSysMenuService 创建一个sBaseSysMenuService实例
func NewsBaseSysMenuService() *sBaseSysMenuService {
	return &sBaseSysMenuService{
		&dzhcore.Service{
			Dao:   &dao.BaseSysMenu,
			Model: model.NewBaseSysMenu(),
			ListQueryOp: &dzhcore.QueryOp{

				Where: func(ctx context.Context) []g.Array {
					var (
						condition []g.Array
						admin     = common.GetAdmin(ctx)
						r         = g.RequestFromCtx(ctx)
						rmap      = r.GetMap()
					)

					//不是超管只看上架的
					if !gstr.Equal(admin.UserId, "1") {
						condition = []g.Array{{"isShow=?", true}}
						condition = []g.Array{{"isInstall=?", true}}
					}

					if rmap["addons"] != nil {

						additionalConditions := []g.Array{
							{"menuType=?", "addon"},
							{"type=?", 0},
							{"parentId IS NULL"},
						}

						condition = append(condition, additionalConditions...)
					}

					return condition
				},
			},
			PageQueryOp: &dzhcore.QueryOp{

				Where: func(ctx context.Context) []g.Array {
					var condition []g.Array
					admin := common.GetAdmin(ctx)

					//不是超管只看上架的
					if !gstr.Equal(admin.UserId, "1") {
						condition = []g.Array{{"isShow=?", true}}
						condition = []g.Array{{"isInstall=?", true}}
					}

					return condition
				},
			},
		},
	}
}

// GetPerms 获取菜单的权限
func (s *sBaseSysMenuService) GetPerms(roleIds []string) []string {
	var (
		perms  []string
		result gdb.Result
	)
	m := s.Dao.Ctx(ctx).As("a")
	// 如果roldIds 包含 1 则表示是超级管理员，则返回所有权限
	if garray.NewIntArrayFrom(gconv.Ints(roleIds)).Contains(1) {
		result, _ = m.Fields("a.perms").All()
	} else {
		result, _ = m.InnerJoin("base_sys_role_menu b", "a.id=b.menuId").InnerJoin("base_sys_role c", "b.roleId=c.id").Where("c.id IN (?)", roleIds).Fields("a.perms").All()
	}
	for _, v := range result {
		vmap := v.Map()
		if vmap["perms"] != nil {
			p := gstr.Split(vmap["perms"].(string), ",")
			perms = append(perms, p...)
		}
	}
	return perms
}

// GetMenus 获取菜单
func (s *sBaseSysMenuService) GetMenus(roleIds []string, isAdmin bool) (result gdb.Result) {
	// 屏蔽 base_sys_role_menu.id 防止部分权限的用户登录时菜单渲染错误
	m := s.Dao.Ctx(ctx).As("a").Fields("a.*")
	if isAdmin {
		result, _ = m.Group("a.id").Order("a.orderNum asc").All()
	} else {
		result, _ = m.InnerJoin("base_sys_role_menu b", "a.id=b.menuId").Where("b.roleId IN (?)", roleIds).Group("a.id").Order("a.orderNum asc").All()
	}
	return

}

// ModifyAfter 修改后
func (s *sBaseSysMenuService) ModifyAfter(ctx context.Context, method string, param g.MapStrAny) (err error) {
	if method == "Delete" {
		ids := gconv.Strings(param["ids"])
		if len(ids) > 0 {
			_, err = s.Dao.Ctx(ctx).Where("parentId IN (?)", ids).Delete()
		}
		return
	}
	return
}

// ServiceAdd 添加
func (s *sBaseSysMenuService) ServiceAdd(ctx context.Context, req *dzhcore.AddReq) (data interface{}, err error) {
	r := g.RequestFromCtx(ctx)
	rjson, err := r.GetJson()
	if err != nil {
		return
	}

	rmap := gconv.Map(rjson)
	m := s.Dao.Ctx(ctx)

	id := dzhcore.NodeSnowflake.Generate().String()
	switch gconv.Int(rmap["type"]) {
	case 0:

		insertData := gconv.Map(rmap["data"])
		insertData["id"] = id
		insertData["isInstall"] = true
		_, _ = m.Data(insertData).Insert()
		data = g.Map{"id": id}
		return

	case 1:

		insertData := gconv.Map(rmap["data"])
		insertData["id"] = id
		insertData["isInstall"] = true
		_, _ = m.Data(insertData).Insert()
		data = g.Map{"id": id}
		return

	case 2:

		var permsList []*types.Perms
		_ = gconv.Scan(rmap["data"], &permsList)
		for _, v := range permsList {
			id = dzhcore.NodeSnowflake.Generate().String()
			v.Id = gconv.String(id)
			v.IsInstall = true
		}
		_, _ = m.Data(permsList).Insert()
		return

	}
	return
}
