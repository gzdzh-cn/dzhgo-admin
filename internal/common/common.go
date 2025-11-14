package common

import (
	"context"
	"dzhgo/internal/dao"
	"dzhgo/internal/model/do"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gzdzh-cn/dzhcore"
)

type Pagination struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
}

type Admin struct {
	IsRefresh       bool     `json:"isRefresh"`
	RoleIds         []string `json:"roleIds"`
	Username        string   `json:"username"`
	UserId          string   `json:"userId"`
	PasswordVersion *int32   `json:"passwordVersion"`
}

// 获取传入ctx 中的 admin 对象
func GetAdmin(ctx context.Context) *Admin {
	r := g.RequestFromCtx(ctx)
	if r == nil {
		g.Log().Warning(ctx, "RequestFromCtx returned nil, cannot get admin info")
		return nil
	}

	var admin *Admin
	adminVar := r.GetCtxVar("admin")
	if adminVar.IsEmpty() {
		g.Log().Warning(ctx, "admin context variable is empty")
		return nil
	}

	err := gjson.New(adminVar.String()).Scan(&admin)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil
	}
	//g.Dump(admin)

	return admin
}

// 记录操作日志
func RecordActionLog(ctx context.Context, actionType, userId, logRemark string) (err error) {
	do := do.BaseSysActionLog{
		Id:     dzhcore.NodeSnowflake.Generate().String(),
		UserId: userId,
		Name:   actionType,
		Remark: logRemark,
	}
	_, err = dao.BaseSysActionLog.Ctx(ctx).Data(do).Insert()
	if err != nil {
		return
	}
	return
}

// 获取全部管理员ID
func getAdminUserIds(ctx context.Context) (data []string, err error) {
	ids, err := dao.BaseSysUser.Ctx(ctx).As("a").
		LeftJoin(dao.BaseSysUserRole.Table(), "user_role", "user_role.userId = a.id").
		LeftJoin(dao.BaseSysRole.Table(), "role", "role.id = user_role.roleId").
		Where("role.id = ?", 1). // 假设角色ID为1的是管理员
		Array("a.id")
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return nil, err
	}
	return gconv.Strings(ids), nil
}

// 获取全部用户ID
func getAllUserIds(ctx context.Context) (data []string, err error) {
	ids, err := dao.BaseSysUser.Ctx(ctx).As("a").
		Array("a.id")
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return nil, err
	}
	return gconv.Strings(ids), nil
}

// 是否超管
func IsSuperAdmin(ctx context.Context) bool {
	admin := GetAdmin(ctx)
	if admin == nil {
		return false
	}
	return garray.NewStrArrayFrom(admin.RoleIds).Contains("1")
}
