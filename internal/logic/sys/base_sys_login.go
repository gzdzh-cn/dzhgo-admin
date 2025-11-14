package sys

import (
	"context"
	"dzhgo/internal/common"
	"dzhgo/internal/dao"
	"dzhgo/internal/model/do"
	"dzhgo/internal/service"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"

	v1 "dzhgo/internal/api/admin_v1"
	"dzhgo/internal/config"
	"dzhgo/internal/model"

	"github.com/gzdzh-cn/dzhcore"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/util/guid"
)

func init() {
	service.RegisterBaseSysLoginService(NewsBaseSysLoginService())
}

type sBaseSysLoginService struct {
	*dzhcore.Service
}

// NewsBaseSysLoginService 创建一个新的sBaseSysLoginService
func NewsBaseSysLoginService() *sBaseSysLoginService {
	return &sBaseSysLoginService{}
}

// Login 登录
func (s *sBaseSysLoginService) Login(ctx context.Context, req *v1.BaseOpenLoginReq) (result *v1.TokenRes, err error) {
	var (
		captchaId  = req.CaptchaId
		verifyCode = req.VerifyCode
		password   = req.Password
		username   = req.Username
		r          = g.RequestFromCtx(ctx)
	)

	vcode, _ := dzhcore.CacheManager.Get(ctx, captchaId)
	if vcode.String() != verifyCode {
		err = gerror.New("验证码错误")
		g.Log().Error(ctx, err.Error())
		return
	}
	md5password, _ := gmd5.Encrypt(password)

	var user *model.BaseSysUser
	dao.BaseSysUser.Ctx(ctx).Where("username=?", username).Where("password=?", md5password).Scan(&user)
	if user != nil && gconv.Int(user.Status) == 0 {
		err = gerror.New("账号已被禁止~")
		g.Log().Error(ctx, err.Error())
		return
	}
	if user == nil {
		err = gerror.New("账户或密码不正确~")
		g.Log().Error(ctx, err.Error())
		return
	}

	dao.BaseSysUser.Ctx(ctx).Where("id=?", user.ID).Data(do.BaseSysUser{
		LoginIp: gconv.String(r.GetClientIp()),
	}).Update()

	result, err = s.GenerateTokenByUser(ctx, user)
	if err != nil {
		return
	}

	common.RecordActionLog(ctx, "登录", user.ID, fmt.Sprintf("登录成功,用户名:%s,IP:%s", user.Username, gconv.String(r.GetClientIp())))

	return
}

// Captcha 图形验证码
func (*sBaseSysLoginService) Captcha(req *v1.BaseOpenCaptchaReq) (interface{}, error) {
	type capchaInfo struct {
		CaptchaId string `json:"captchaId"`
		Data      string `json:"data"`
	}
	var (
		ctx    = gctx.GetInitCtx()
		err    error
		result = &capchaInfo{}
	)
	captchaText := grand.Digits(4)
	svg := `<svg width="150" height="50" xmlns="http://www.w3.org/2000/svg"><text x="75" y="25" text-anchor="middle" font-size="25" fill="#fff">` + captchaText + `</text></svg>`
	svgbase64 := gbase64.EncodeString(svg)

	result.Data = `data:image/svg+xml;base64,` + svgbase64
	result.CaptchaId = guid.S()
	dzhcore.CacheManager.Set(ctx, result.CaptchaId, captchaText, 1800*time.Second)
	return result, err
}

// Logout 退出登录
func (*sBaseSysLoginService) Logout(ctx context.Context) (err error) {
	admin := common.GetAdmin(ctx)
	if admin == nil {
		g.Log().Warning(ctx, "无法获取admin信息，退出登录失败")
		return gerror.New("无法获取用户信息")
	}

	userId := admin.UserId
	dzhcore.CacheManager.Remove(ctx, "admin:department:"+gconv.String(userId))
	dzhcore.CacheManager.Remove(ctx, "admin:perms:"+gconv.String(userId))
	dzhcore.CacheManager.Remove(ctx, "admin:token:"+gconv.String(userId))
	dzhcore.CacheManager.Remove(ctx, "admin:token:refresh:"+gconv.String(userId))
	return
}

// RefreshToken 刷新token
func (s *sBaseSysLoginService) RefreshToken(ctx context.Context, token string) (result *v1.TokenRes, err error) {

	g.Log().Debugf(ctx, "RefreshToken")
	tokenClaims, err := jwt.ParseWithClaims(token, &dzhcore.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.Jwt.Secret), nil
	})
	if err != nil {
		return
	}
	claims, ok := tokenClaims.Claims.(*dzhcore.Claims)
	if !ok {
		err = gerror.New("tokenClaims.Claims.(*Claims) error")
		return
	}
	if !tokenClaims.Valid {
		err = gerror.New("tokenClaims.Valid error")
		return
	}
	if !claims.IsRefresh {
		err = gerror.New("claims.IsRefresh error")
		return
	}

	if !(gconv.Uint64(claims.UserId) > 0) {
		err = gerror.New("claims.UserId error")
		return
	}

	var (
		user *model.BaseSysUser
	)
	err = dao.BaseSysUser.Ctx(ctx).Where("id=?", claims.UserId).Where("status=?", 1).Scan(&user)
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return nil, err
	}
	if user == nil {
		err = gerror.New("用户不存在")
		return
	}

	result, err = s.GenerateTokenByUser(ctx, user)
	return
}

// generateToken  生成token
func (*sBaseSysLoginService) generateToken(ctx context.Context, user *model.BaseSysUser, roleIds []string, exprire uint, isRefresh bool) (token string) {
	err := dzhcore.CacheManager.Set(ctx, "admin:passwordVersion:"+gconv.String(user.ID), gconv.String(user.PasswordV), 0)
	if err != nil {
		g.Log().Error(ctx, "生成token失败", err)
	}

	claims := &dzhcore.Claims{
		IsRefresh:       isRefresh,
		RoleIds:         roleIds,
		Username:        user.Username,
		UserId:          user.ID,
		PasswordVersion: user.PasswordV,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(exprire) * time.Second)),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = tokenClaims.SignedString([]byte(config.Cfg.Jwt.Secret))
	if err != nil {
		g.Log().Error(ctx, "生成token失败", err)
	}
	return
}

// 根据用户生成前端需要的Token信息
func (s *sBaseSysLoginService) GenerateTokenByUser(ctx context.Context, user *model.BaseSysUser) (result *v1.TokenRes, err error) {

	// 获取用户角色
	roleIds := service.BaseSysRoleService().GetByUser(user.ID)
	// 如果没有角色，则报错
	if len(roleIds) == 0 {
		err = gerror.New("该用户未设置任何角色，无法登录~")
		return
	}

	// 生成token
	result = &v1.TokenRes{}
	result.Expire = config.Cfg.Jwt.Token.Expire
	result.RefreshExpire = config.Cfg.Jwt.Token.RefreshExpire
	result.Token = s.generateToken(ctx, user, roleIds, result.Expire, false)
	result.RefreshToken = s.generateToken(ctx, user, roleIds, result.RefreshExpire, true)
	// 将用户相关信息保存到缓存
	perms := service.BaseSysMenuService().GetPerms(roleIds)
	departments := service.BaseSysDepartmentService().GetByRoleIds(roleIds)
	dzhcore.CacheManager.Set(ctx, "admin:department:"+gconv.String(user.ID), departments, 0)
	dzhcore.CacheManager.Set(ctx, "admin:perms:"+gconv.String(user.ID), perms, 0)
	dzhcore.CacheManager.Set(ctx, "admin:token:"+gconv.String(user.ID), result.Token, 0)
	dzhcore.CacheManager.Set(ctx, "admin:token:refresh:"+gconv.String(user.ID), result.RefreshToken, 0)

	return
}
