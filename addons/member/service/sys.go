// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "dzhgo/addons/member/api/app_v1"
	memberDefineType "dzhgo/addons/member/defineType"
	"dzhgo/addons/member/model/entity"
	"dzhgo/internal/defineType"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gzdzh-cn/dzhcore"
)

type (
	IMemberManageService interface {
		// 新增|删除|修改前的操作
		ModifyBefore(ctx context.Context, method string, param g.MapStrAny) (err error)
		// 新增|删除|修改后的操作
		ModifyAfter(ctx context.Context, method string, param g.MapStrAny) (err error)
		// 新增
		ServiceAdd(ctx context.Context, req *dzhcore.AddReq) (data interface{}, err error)
		// 查询
		ServiceInfo(ctx context.Context, req *dzhcore.InfoReq) (data interface{}, err error)
		// 账号登录
		AccountLogin(ctx context.Context, req *v1.AccountLoginReq) (data interface{}, err error)
		// 公众号登录
		MpLogin(ctx context.Context, req *v1.MpLoginReq) (data interface{}, err error)
		// 通过code获取token
		GetWxAccessToken(ctx context.Context, code string, wxConfig *defineType.WxConfig) (data *memberDefineType.WxMpTokenResponse, err error)
		// 获取解密后的数据
		GetUserInfo(ctx context.Context, wxMpTokenResponse *memberDefineType.WxMpTokenResponse) (data *memberDefineType.WxMpUserInfoResponse, err error)
		// Person 方法 返回不带密码的用户信息
		Person(ctx context.Context, userId string) (data interface{}, err error)
		// 根据用户生成前端需要的Token信息
		GenerateTokenByUser(ctx g.Ctx, member *entity.AddonsMemberManage) (result *v1.TokenRes, err error)
	}
)

var (
	localMemberManageService IMemberManageService
)

func MemberManageService() IMemberManageService {
	if localMemberManageService == nil {
		panic("implement not found for interface IMemberManageService, forgot register?")
	}
	return localMemberManageService
}

func RegisterMemberManageService(i IMemberManageService) {
	localMemberManageService = i
}
