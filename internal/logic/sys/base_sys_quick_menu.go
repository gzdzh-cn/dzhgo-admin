package sys

import (
	"context"
	"dzhgo/internal/common"
	"dzhgo/internal/dao"
	"dzhgo/internal/model"
	"dzhgo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gzdzh-cn/dzhcore"
)

func init() {
	service.RegisterBaseSysQuickMenuService(&sBaseSysQuickMenuService{})
}

type sBaseSysQuickMenuService struct {
	*dzhcore.Service
}

func NewsBaseSysQuickMenuService() *sBaseSysQuickMenuService {
	return &sBaseSysQuickMenuService{
		&dzhcore.Service{
			Dao:   &dao.BaseSysQuickMenu,
			Model: model.NewBaseSysQuickMenu(),
			ListQueryOp: &dzhcore.QueryOp{
				FieldEQ:      []string{""},                   // 字段等于
				KeyWordField: []string{""},                   // 模糊搜索匹配的数据库字段
				AddOrderby:   g.MapStrStr{"orderNum": "ASC"}, // 按排序升序
				Where: func(ctx context.Context) []g.Array { // 自动按当前用户过滤
					admin := common.GetAdmin(ctx)
					return []g.Array{
						{"user_id = ?", admin.UserId},
					}
				},
				OrWhere: func(ctx context.Context) []g.Array {
					return []g.Array{}
				},
				Select: "",
				As:     "",
				Join:   []*dzhcore.JoinOp{},
				Extend: func(ctx g.Ctx, m *gdb.Model) *gdb.Model {
					return m
				},
				ModifyResult: func(ctx g.Ctx, data any) any {
					return data
				},
			},
			PageQueryOp: &dzhcore.QueryOp{
				FieldEQ:      []string{""},
				KeyWordField: []string{""},
				AddOrderby:   g.MapStrStr{"orderNum": "ASC"},
				Where: func(ctx context.Context) []g.Array {
					admin := common.GetAdmin(ctx)
					return []g.Array{
						{"user_id = ?", admin.UserId},
					}
				},
				OrWhere: func(ctx context.Context) []g.Array {
					return []g.Array{}
				},
				Select: "",
				As:     "",
				Join:   []*dzhcore.JoinOp{},
				Extend: func(ctx g.Ctx, m *gdb.Model) *gdb.Model {
					return m
				},
				ModifyResult: func(ctx g.Ctx, data any) any {
					return data
				},
			},
			InsertParam: func(ctx context.Context) g.MapStrAny { // Add时自动注入user_id和name
				admin := common.GetAdmin(ctx)
				r := g.RequestFromCtx(ctx)
				menuId := r.Get("menuId").String()
				insertData := g.MapStrAny{
					"userId": admin.UserId,
				}
				if menuId != "" {
					menuInfo, _ := dao.BaseSysMenu.Ctx(ctx).Where("id", menuId).One()
					if menuInfo != nil {
						insertData["name"] = gconv.String(menuInfo["name"])
					}
				}
				return insertData
			},
			Before: func(ctx context.Context) (err error) {
				return nil
			},
			InfoIgnoreProperty: "userId,deletedAt",
			UniqueKey:          g.MapStrStr{},
			NotNullKey:         g.MapStrStr{},
		},
	}
}

// ServiceList 重写父结构体的ServiceList方法，返回当前用户的快捷菜单
func (s *sBaseSysQuickMenuService) ServiceList(ctx context.Context, req *dzhcore.ListReq) (data any, err error) {
	type QuickMenuListItem struct {
		Id       string `json:"id"`
		MenuId   string `json:"menuId"`
		Name     string `json:"name"`
		Icon     string `json:"icon"`
		Router   string `json:"router"`
		OrderNum int    `json:"orderNum"`
	}

	admin := common.GetAdmin(ctx)
	var list []*QuickMenuListItem
	err = dao.BaseSysQuickMenu.Ctx(ctx).As("q").
		LeftJoin("base_sys_menu m", "q.menu_id=m.id COLLATE utf8mb4_unicode_ci").
		Fields("q.id, q.menu_id AS menuId, q.name, q.icon, IFNULL(q.router, m.router) AS router, q.order_num AS orderNum").
		Where("q.status", 1).
		Where("q.user_id", admin.UserId).
		Order("q.order_num asc").
		Scan(&list)
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = []*QuickMenuListItem{}
	}
	return list, nil
}

// QuickMenuList 获取用户有权限的菜单列表（用于前端选择添加快捷菜单）
func (s *sBaseSysQuickMenuService) QuickMenuList(ctx context.Context, roleIds []string) (data interface{}, err error) {
	type QuickMenuItem struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Icon   string `json:"icon"`
		Router string `json:"router"`
	}

	// 获取用户有权限的菜单
	menus := service.BaseSysMenuService().GetMenus(roleIds)
	var list []*QuickMenuItem
	for _, menu := range menus {
		menuRecord := menu.Map()
		if gconv.Int(menuRecord["type"]) == 1 && menuRecord["router"] != nil && gconv.String(menuRecord["router"]) != "" {
			list = append(list, &QuickMenuItem{
				Id:     gconv.String(menuRecord["id"]),
				Name:   gconv.String(menuRecord["name"]),
				Icon:   gconv.String(menuRecord["icon"]),
				Router: gconv.String(menuRecord["router"]),
			})
		}
	}
	if list == nil {
		list = []*QuickMenuItem{}
	}
	return list, nil
}

func (s *sBaseSysQuickMenuService) Test(ctx context.Context) (err error) {
	return nil
}
