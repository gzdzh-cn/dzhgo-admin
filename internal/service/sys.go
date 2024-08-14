// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "dzhgo/internal/api/admin_v1"
	"dzhgo/internal/model"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gzdzh-cn/dzhcore"
)

type (
	IBaseOpenService interface {
		// AdminEPS 获取eps
		AdminEPS(ctx g.Ctx) (result *g.Var, err error)
		// AdminEPS 获取eps
		AppEPS(ctx g.Ctx) (result *g.Var, err error)
		// 获取全部版本
		Versions(ctx context.Context, req *v1.VersionsReq) (data interface{}, err error)
		// 站点配置
		GetSetting(ctx context.Context, req *v1.GetSettingReq) (data interface{}, err error)
	}
	IBaseSysAddonsService interface {
		// 安装卸载插件
		InstallUpdateStatus(ctx context.Context, req *v1.InstallUpdateStatusReq) (data interface{}, err error)
		// 上下架插件
		LineUpdateStatus(ctx context.Context, req *v1.LineUpdateStatusReq) (data interface{}, err error)
	}
	IBaseSysAddonsTypeService interface{}
	IBaseSysConfService       interface {
		// UpdateValue 更新配置值
		UpdateValue(cKey, cValue string) error
		// GetValue 获取配置值
		GetValue(cKey string) string
	}
	IBaseSysDepartmentService interface {
		// GetByRoleIds 获取部门
		GetByRoleIds(roleIds []string, isAdmin bool) (res []uint)
		// Order 排序部门
		Order(ctx g.Ctx) (err error)
	}
	IBaseSysLogService interface {
		// Record 记录日志
		Record(ctx g.Ctx)
		// Clear 清除日志
		Clear(isAll bool) (err error)
	}
	IBaseSysLoginService interface {
		// Login 登录
		Login(ctx context.Context, req *v1.BaseOpenLoginReq) (result *v1.TokenRes, err error)
		// Captcha 图形验证码
		Captcha(req *v1.BaseOpenCaptchaReq) (interface{}, error)
		// Logout 退出登录
		Logout(ctx context.Context) (err error)
		// RefreshToken 刷新token
		RefreshToken(ctx context.Context, token string) (result *v1.TokenRes, err error)
		// 根据用户生成前端需要的Token信息
		GenerateTokenByUser(ctx context.Context, user *model.BaseSysUser) (result *v1.TokenRes, err error)
	}
	IBaseSysMemberLoginService interface {
		Login(ctx g.Ctx, req *v1.LoginReq) (result *v1.TokenRes, err error)
		// MpLoginReq 微信公众号登录
		MpLoginReq(ctx g.Ctx, req *v1.MpLoginReq) (result *v1.TokenRes, err error)
		// MiniLogin 小程序登录
		MiniLogin(ctx g.Ctx, req *v1.MiniLoginReq) (result *v1.TokenRes, err error)
		// AutoPhone 手机授权登录
		AutoPhone(ctx g.Ctx, req *v1.AutoPhoneReq) (result *v1.TokenRes, err error)
		// VerifyCount 验证游客次数
		VerifyCount(ctx g.Ctx, req *v1.VerifyCountReq) (data interface{}, err error)
		// AccountRegister 账号注册
		AccountRegister(ctx g.Ctx, req *v1.AccountRegisterReq) (result *v1.TokenRes, err error)
	}
	IBaseSysMenuService interface {
		// GetPerms 获取菜单的权限
		GetPerms(roleIds []string) []string
		// GetMenus 获取菜单
		GetMenus(roleIds []string, isAdmin bool) (result gdb.Result)
		// ModifyAfter 修改后
		ModifyAfter(ctx context.Context, method string, param g.MapStrAny) (err error)
		// ServiceAdd 添加
		ServiceAdd(ctx context.Context, req *dzhcore.AddReq) (data interface{}, err error)
	}
	IBaseSysParamService interface {
		// HtmlByKey 根据配置参数key获取网页内容(富文本)
		HtmlByKey(key string) string
		// ModifyAfter 修改后
		ModifyAfter(ctx context.Context, method string, param g.MapStrAny) (err error)
		// DataByKey 根据配置参数key获取数据
		DataByKey(ctx context.Context, key string) (data string, err error)
	}
	IBaseSysPermsService interface {
		// permmenu 方法
		Permmenu(ctx context.Context, roleIds []string) (res interface{})
		// RefreshPerms refreshPerms(userId)
		RefreshPerms(ctx context.Context, userId string) (err error)
	}
	IBaseSysRoleService interface {
		// ModifyAfter modify after
		ModifyAfter(ctx context.Context, method string, param g.MapStrAny) (err error)
		// GetByUser get array  roleId by userId
		GetByUser(userId string) []string
		// ServiceInfo 方法重构
		ServiceInfo(ctx context.Context, req *dzhcore.InfoReq) (data interface{}, err error)
	}
	IBaseSysSettingService interface{}
	IBaseSysUserService    interface {
		// Person 方法 返回不带密码的用户信息
		Person(userId string) (res gdb.Record, err error)
		ModifyBefore(ctx context.Context, method string, param g.MapStrAny) (err error)
		ModifyAfter(ctx context.Context, method string, param g.MapStrAny) (err error)
		// ServiceAdd 方法 添加用户
		ServiceAdd(ctx context.Context, req *dzhcore.AddReq) (data interface{}, err error)
		// ServiceInfo 方法 返回服务信息
		ServiceInfo(ctx g.Ctx, req *dzhcore.InfoReq) (data interface{}, err error)
		// ServiceUpdate 方法 更新用户信息
		ServiceUpdate(ctx context.Context, req *dzhcore.UpdateReq) (data interface{}, err error)
		// Move 移动用户部门
		Move(ctx g.Ctx) (err error)
	}
)

var (
	localBaseOpenService           IBaseOpenService
	localBaseSysAddonsService      IBaseSysAddonsService
	localBaseSysAddonsTypeService  IBaseSysAddonsTypeService
	localBaseSysConfService        IBaseSysConfService
	localBaseSysDepartmentService  IBaseSysDepartmentService
	localBaseSysLogService         IBaseSysLogService
	localBaseSysLoginService       IBaseSysLoginService
	localBaseSysMemberLoginService IBaseSysMemberLoginService
	localBaseSysMenuService        IBaseSysMenuService
	localBaseSysParamService       IBaseSysParamService
	localBaseSysPermsService       IBaseSysPermsService
	localBaseSysRoleService        IBaseSysRoleService
	localBaseSysSettingService     IBaseSysSettingService
	localBaseSysUserService        IBaseSysUserService
)

func BaseOpenService() IBaseOpenService {
	if localBaseOpenService == nil {
		panic("implement not found for interface IBaseOpenService, forgot register?")
	}
	return localBaseOpenService
}

func RegisterBaseOpenService(i IBaseOpenService) {
	localBaseOpenService = i
}

func BaseSysAddonsService() IBaseSysAddonsService {
	if localBaseSysAddonsService == nil {
		panic("implement not found for interface IBaseSysAddonsService, forgot register?")
	}
	return localBaseSysAddonsService
}

func RegisterBaseSysAddonsService(i IBaseSysAddonsService) {
	localBaseSysAddonsService = i
}

func BaseSysAddonsTypeService() IBaseSysAddonsTypeService {
	if localBaseSysAddonsTypeService == nil {
		panic("implement not found for interface IBaseSysAddonsTypeService, forgot register?")
	}
	return localBaseSysAddonsTypeService
}

func RegisterBaseSysAddonsTypeService(i IBaseSysAddonsTypeService) {
	localBaseSysAddonsTypeService = i
}

func BaseSysConfService() IBaseSysConfService {
	if localBaseSysConfService == nil {
		panic("implement not found for interface IBaseSysConfService, forgot register?")
	}
	return localBaseSysConfService
}

func RegisterBaseSysConfService(i IBaseSysConfService) {
	localBaseSysConfService = i
}

func BaseSysDepartmentService() IBaseSysDepartmentService {
	if localBaseSysDepartmentService == nil {
		panic("implement not found for interface IBaseSysDepartmentService, forgot register?")
	}
	return localBaseSysDepartmentService
}

func RegisterBaseSysDepartmentService(i IBaseSysDepartmentService) {
	localBaseSysDepartmentService = i
}

func BaseSysLogService() IBaseSysLogService {
	if localBaseSysLogService == nil {
		panic("implement not found for interface IBaseSysLogService, forgot register?")
	}
	return localBaseSysLogService
}

func RegisterBaseSysLogService(i IBaseSysLogService) {
	localBaseSysLogService = i
}

func BaseSysLoginService() IBaseSysLoginService {
	if localBaseSysLoginService == nil {
		panic("implement not found for interface IBaseSysLoginService, forgot register?")
	}
	return localBaseSysLoginService
}

func RegisterBaseSysLoginService(i IBaseSysLoginService) {
	localBaseSysLoginService = i
}

func BaseSysMemberLoginService() IBaseSysMemberLoginService {
	if localBaseSysMemberLoginService == nil {
		panic("implement not found for interface IBaseSysMemberLoginService, forgot register?")
	}
	return localBaseSysMemberLoginService
}

func RegisterBaseSysMemberLoginService(i IBaseSysMemberLoginService) {
	localBaseSysMemberLoginService = i
}

func BaseSysMenuService() IBaseSysMenuService {
	if localBaseSysMenuService == nil {
		panic("implement not found for interface IBaseSysMenuService, forgot register?")
	}
	return localBaseSysMenuService
}

func RegisterBaseSysMenuService(i IBaseSysMenuService) {
	localBaseSysMenuService = i
}

func BaseSysParamService() IBaseSysParamService {
	if localBaseSysParamService == nil {
		panic("implement not found for interface IBaseSysParamService, forgot register?")
	}
	return localBaseSysParamService
}

func RegisterBaseSysParamService(i IBaseSysParamService) {
	localBaseSysParamService = i
}

func BaseSysPermsService() IBaseSysPermsService {
	if localBaseSysPermsService == nil {
		panic("implement not found for interface IBaseSysPermsService, forgot register?")
	}
	return localBaseSysPermsService
}

func RegisterBaseSysPermsService(i IBaseSysPermsService) {
	localBaseSysPermsService = i
}

func BaseSysRoleService() IBaseSysRoleService {
	if localBaseSysRoleService == nil {
		panic("implement not found for interface IBaseSysRoleService, forgot register?")
	}
	return localBaseSysRoleService
}

func RegisterBaseSysRoleService(i IBaseSysRoleService) {
	localBaseSysRoleService = i
}

func BaseSysSettingService() IBaseSysSettingService {
	if localBaseSysSettingService == nil {
		panic("implement not found for interface IBaseSysSettingService, forgot register?")
	}
	return localBaseSysSettingService
}

func RegisterBaseSysSettingService(i IBaseSysSettingService) {
	localBaseSysSettingService = i
}

func BaseSysUserService() IBaseSysUserService {
	if localBaseSysUserService == nil {
		panic("implement not found for interface IBaseSysUserService, forgot register?")
	}
	return localBaseSysUserService
}

func RegisterBaseSysUserService(i IBaseSysUserService) {
	localBaseSysUserService = i
}
