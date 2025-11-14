package sys

import (
	"context"
	"dzhgo/internal/common"
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
	service.RegisterBaseSysRoleService(NewsBaseSysRoleService())
}

type sBaseSysRoleService struct {
	*dzhcore.Service
}

// NewsBaseSysRoleService create a new sBaseSysRoleService
func NewsBaseSysRoleService() *sBaseSysRoleService {
	return &sBaseSysRoleService{
		Service: &dzhcore.Service{
			Dao:   &dao.BaseSysRole,
			Model: model.NewBaseSysRole(),
			ListQueryOp: &dzhcore.QueryOp{
				Where: func(ctx context.Context) [][]interface{} {
					var (
						admin   = common.GetAdmin(ctx)
						userId  = admin.UserId
						roleIds = garray.NewStrArrayFromCopy(admin.RoleIds)
					)
					return [][]interface{}{
						// 超级管理员的角色不展示
						{"label != ?", g.Slice{"admin"}, true},
						// 如果不是超管，只能看到自己新建的或者自己有的角色
						{"(userId=? or id in (?))", g.Slice{userId, admin.RoleIds}, !roleIds.Contains("1")},
					}
				},
			},
			PageQueryOp: &dzhcore.QueryOp{
				KeyWordField: []string{"name", "label"},
				AddOrderby:   map[string]string{},
				Where: func(ctx context.Context) [][]interface{} {
					var (
						admin   = common.GetAdmin(ctx)
						userId  = admin.UserId
						roleIds = garray.NewStrArrayFromCopy(admin.RoleIds)
					)
					return [][]interface{}{
						// 超级管理员的角色不展示
						{"label != ?", g.Slice{"admin"}, true},
						// 如果不是超管，只能看到自己新建的或者自己有的角色
						{"(userId=? or id in (?))", g.Slice{userId, admin.RoleIds}, !roleIds.Contains("1")},
					}
				},
			},
			InsertParam: func(ctx context.Context) map[string]interface{} {
				return g.Map{"userId": common.GetAdmin(ctx).UserId}
			},
			UniqueKey: map[string]string{
				"name":  "角色名称不能重复",
				"label": "角色标识不能重复",
			},
		},
	}
}

// ModifyAfter modify after
func (s *sBaseSysRoleService) ModifyAfter(ctx context.Context, method string, param g.MapStrAny) (err error) {

	if param["id"] != nil {
		err = s.updatePerms(ctx, gconv.String(param["id"]), gconv.SliceStr(param["menuIdList"]), gconv.SliceStr(param["departmentIdList"]))
	}
	return
}

// updatePerms(roleId, menuIdList?, departmentIds = [])
func (s *sBaseSysRoleService) updatePerms(ctx context.Context, roleId string, menuIdList, departmentIds []string) (err error) {

	// 更新菜单权限
	_, err = dao.BaseSysRoleMenu.Ctx(ctx).Where("roleId = ?", roleId).Delete()
	if err != nil {
		return err
	}

	if len(menuIdList) > 0 {

		roleMenuList := make([]g.MapStrAny, len(menuIdList))
		for i, menuId := range menuIdList {
			roleMenuList[i] = g.MapStrAny{
				"id":     dzhcore.NodeSnowflake.Generate().String(),
				"roleId": roleId,
				"menuId": menuId,
			}
		}
		_, err = dao.BaseSysRoleMenu.Ctx(ctx).Data(roleMenuList).Insert()
		if err != nil {
			return err
		}
	}

	// 更新部门权限
	_, err = dao.BaseSysRoleDepartment.Ctx(ctx).Where("roleId = ?", roleId).Delete()
	if err != nil {
		return err
	}
	if len(departmentIds) > 0 {

		roleDepartmentList := make([]g.MapStrAny, len(departmentIds))
		for i, departmentId := range departmentIds {
			roleDepartmentList[i] = g.MapStrAny{
				"id":           dzhcore.NodeSnowflake.Generate().String(),
				"roleId":       roleId,
				"departmentId": departmentId,
			}
		}
		_, err = dao.BaseSysRoleDepartment.Ctx(ctx).Data(roleDepartmentList).Insert()
		if err != nil {
			return err
		}
	}

	// 刷新权限
	userRoles, err := dao.BaseSysUserRole.Ctx(ctx).Where("roleId = ?", roleId).All()
	if err != nil {
		return
	}
	for _, v := range userRoles {
		vMap := v.Map()
		if vMap["userId"] != nil {
			err = service.BaseSysPermsService().RefreshPerms(ctx, gconv.String(vMap["userId"]))
			if err != nil {
				return err
			}
		}
	}

	return

}

// GetByUser get array  roleId by userId
func (s *sBaseSysRoleService) GetByUser(userId string) []string {
	var (
		roles []string
	)
	res, _ := dao.BaseSysUserRole.Ctx(ctx).Where("userId = ?", userId).Array("roleId")
	for _, v := range res {
		roles = append(roles, gconv.String(v))
	}
	return roles
}

// ServiceInfo 方法重构
func (s *sBaseSysRoleService) ServiceInfo(ctx context.Context, req *dzhcore.InfoReq) (data interface{}, err error) {
	info, err := s.Dao.Ctx(ctx).Where("id = ?", req.Id).One()
	if err != nil {
		return nil, err
	}
	if !info.IsEmpty() {
		var menus gdb.Result
		if req.Id == 1 {
			menus, err = dao.BaseSysMenu.Ctx(ctx).All()
			if err != nil {
				return nil, err
			}
		} else {
			menus, err = dao.BaseSysRoleMenu.Ctx(ctx).Where("roleId = ?", req.Id).All()
			if err != nil {
				return nil, err
			}
		}
		menuIdList := garray.NewStrArray()
		for _, v := range menus {
			menuIdList.Append(gconv.String(v["menuId"]))
		}
		var departments gdb.Result
		if req.Id == 1 {
			departments, err = dao.BaseSysRoleDepartment.Ctx(ctx).All()
			if err != nil {
				return nil, err
			}
		} else {
			departments, err = dao.BaseSysRoleDepartment.Ctx(ctx).Where("roleId = ?", req.Id).All()
			if err != nil {
				return nil, err
			}
		}
		departmentIdList := garray.NewStrArray()
		for _, v := range departments {
			departmentIdList.Append(gconv.String(v["departmentId"]))
		}
		result := gconv.Map(info)
		result["menuIdList"] = menuIdList.Slice()
		result["departmentIdList"] = departmentIdList.Slice()
		data = result
		return
	}
	data = g.Map{}
	return
}
