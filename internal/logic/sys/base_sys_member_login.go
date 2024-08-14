package sys

import (
	v1 "dzhgo/internal/api/admin_v1"
	"dzhgo/internal/config"
	"dzhgo/internal/model"
	"dzhgo/internal/service"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gzdzh-cn/dzhcore"
	"time"
)

func init() {
	service.RegisterBaseSysMemberLoginService(NewsBaseSysMemberLoginService())
}

type sBaseSysMemberLoginService struct {
	*dzhcore.Service
}

func NewsBaseSysMemberLoginService() *sBaseSysMemberLoginService {
	return &sBaseSysMemberLoginService{
		&dzhcore.Service{},
	}
}

type LoginData struct {
	SessionKey string `json:"session_key"`
	Openid     string `json:"openid"`
	UnionId    string `json:"unionId"`
}

func (s *sBaseSysMemberLoginService) Login(ctx g.Ctx, req *v1.LoginReq) (result *v1.TokenRes, err error) {

	url := "https://api.weixin.qq.com/sns/jscode2session"
	params := g.Map{
		"appid":      "wx44b19e138b892f99",
		"secret":     "1e483dec348f5749f5dc30e3ec6d0afa",
		"js_code":    req.JsCode,
		"grant_type": "authorization_code",
	}

	// 请求腾讯开发者服务器,code换取openid
	codeRes, err := g.Client().Post(ctx, url, params)
	if err != nil {
		panic(err)
	}
	defer func(codeRes *gclient.Response) {
		err := codeRes.Close()
		if err != nil {

		}
	}(codeRes)

	// 读取响应体并且json转换到结构体
	loginData := &LoginData{}
	err = json.NewDecoder(codeRes.Body).Decode(loginData)
	if err != nil {
		g.Log().Error(ctx, err.Error(), err)
		err = gerror.New("读取响应体失败")
		return
	}
	var (
		member     *model.BaseSysMember
		memberId   interface{}
		memberData g.Map
	)

	err = s.GetDao().Ctx(ctx).Where("openid=?", loginData.Openid).Where("status=?", 1).Scan(&member)
	if err != nil {
		return nil, err
	}
	if member == nil {
		md5password, _ := gmd5.Encrypt("123456")
		memberData = g.Map{
			"username":   req.NickName,
			"password":   md5password,
			"nickName":   req.NickName,
			"levelName":  "普通会员",
			"avatarUrl":  req.AvatarUrl,
			"openid":     loginData.Openid,
			"unionId":    loginData.UnionId,
			"sessionKey": loginData.SessionKey,
		}
		memberId, err = s.Dao.Ctx(ctx).Data(memberData).InsertAndGetId()
		if err != nil {
			g.Log().Error(ctx, "func Login insert :"+err.Error(), err)
			err = gerror.New("读取响应体失败")
			return
		}
		memberData["id"] = memberId
		err := gconv.Struct(memberData, &member)
		if err != nil {
			return nil, err
		}
	} else {
		memberData = g.Map{
			"username":  req.NickName,
			"avatarUrl": req.AvatarUrl,
		}
		_, err = s.Dao.Ctx(ctx).Where("openid=?", loginData.Openid).Data(memberData).Update()
		err = s.GetDao().Ctx(ctx).Where("openid=?", loginData.Openid).Scan(&member)
		if err != nil {
			err = gerror.New("查询会员失败")
			return
		}
	}

	result, err = s.generateTokenByUser(ctx, member)
	if err != nil {
		return
	}

	return
}

//// MemberInfo 会员信息
//func (s *sBaseSysMemberLoginService) MemberInfo(ctx g.Ctx) (data interface{}, err error) {
//
//	type MemberData struct {
//		*types.Member
//		*types.MemberSum
//		*types.MemberAttr
//	}
//	var (
//		memberData *MemberData
//	)
//
//	member := common.GetMember(ctx)
//	m := dao.BaseSysMember.Ctx(ctx).Where("id", member.UserId).Fields("addons_cms_member_user.*,attr.followJson,attr.likeJson,attr.collectJson,attr.commentLikeJson,attr.noticeJson")
//	m.LeftJoin(dao.BaseSysMember.Table(), "addons_cms_member_user.id = attr.userId").As("attr")
//
//	err = m.Scan(&memberData)
//	if err != nil {
//		g.Log().Debugf(ctx, "PersonData find member error %v", err.Error())
//		err := gerror.New("查询失败")
//		return nil, err
//	}
//
//	data = memberData
//	return
//}

// MpLoginReq 微信公众号登录
func (s *sBaseSysMemberLoginService) MpLoginReq(ctx g.Ctx, req *v1.MpLoginReq) (result *v1.TokenRes, err error) {

	var (
		r    = g.RequestFromCtx(ctx)
		rmap = r.GetMap()
	)

	_, err = s.Dao.Ctx(ctx).Data(rmap).Save()
	if err != nil {
		g.Log().Error(ctx, err.Error(), err)
		err = gerror.New("录入失败")
		return
	}

	fmt.Printf("rmap: %v\n", rmap)

	return
}

// MiniLogin 小程序登录
func (s *sBaseSysMemberLoginService) MiniLogin(ctx g.Ctx, req *v1.MiniLoginReq) (result *v1.TokenRes, err error) {

	wxCon := gconv.Map(config.GetWxCon())
	url := gconv.String(wxCon["requestUrl"]) + "?appid=" + gconv.String(wxCon["appid"]) + "&secret=" + gconv.String(wxCon["secret"]) + "&js_code=" + req.JsCode + "&grant_type=authorization_code"

	// 请求腾讯开发者服务器,code换取openid
	client := g.Client()
	codeRes, err := client.ContentJson().Get(ctx, url)
	if err != nil {
		g.Log().Errorf(ctx, "request wx service err:%v", err.Error())
		panic(err)
	}

	defer func(codeRes *gclient.Response) {
		err = codeRes.Close()
		if err != nil {
			g.Log().Error(ctx, err.Error(), err)
		}
	}(codeRes)

	// 读取响应体并且json转换到结构体
	loginData := &LoginData{}
	err = json.NewDecoder(codeRes.Body).Decode(loginData)
	if err != nil {
		g.Log().Error(ctx, err.Error(), err)
		err = gerror.New("读取响应体失败")
		return
	}
	if gconv.String(loginData) == "{}" {
		g.Log().Errorf(ctx, "request wx server codeRes.Body:%v", codeRes.Body)
		err = gerror.New("请求微信服务器错误")
		return
	}
	var (
		member     *model.BaseSysMember
		memberId   interface{}
		memberData g.Map
	)

	err = s.Dao.Ctx(ctx).Where("openid=?", loginData.Openid).Where("status=?", 1).Scan(&member)
	if err != nil {
		g.Log().Error(ctx, " Scan :"+err.Error(), err)
		return
	}
	// 不存在
	if member == nil {

		db := g.DB()
		if tx, err := db.Begin(ctx); err == nil {
			md5password, _ := gmd5.Encrypt("123456")
			characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
			username := "ID_" + grand.Str(characters, 5)
			memberData = g.Map{
				"password":   md5password,
				"levelName":  "普通会员",
				"openid":     loginData.Openid,
				"unionid":    loginData.UnionId,
				"sessionKey": loginData.SessionKey,
				"username":   username,
			}
			memberId, err = tx.Model(s.Dao).Data(memberData).InsertAndGetId()
			if err != nil {
				g.Log().Error(ctx, "func Login insert :"+err.Error(), err)
				err = gerror.New("读取响应体失败")
				tx.Rollback()
				return nil, err
			}
			memberData["id"] = memberId
			err = gconv.Struct(memberData, &member)
			if err != nil {
				g.Log().Error(ctx, "memberData Struct :"+err.Error(), err)
				return nil, err
			}

			_, err = tx.Model(model.NewBaseMemberAttr()).Data(g.Map{"userId": memberId}).Insert()
			if err != nil {
				g.Log().Error(ctx, "NewBaseMemberAttr insert :"+err.Error(), err)
				err = gerror.New("写入失败")
				tx.Rollback()
				return nil, err
			}

			if err == nil {
				tx.Commit()
			}
		}
	}

	result, err = s.generateTokenByUser(ctx, member)
	if err != nil {
		return
	}

	return
}

// AutoPhone 手机授权登录
func (s *sBaseSysMemberLoginService) AutoPhone(ctx g.Ctx, req *v1.AutoPhoneReq) (result *v1.TokenRes, err error) {

	// 获取token
	type AccessTokenRes struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   string `json:"expires_in"`
	}
	accessTokenConfig := config.GetAccessToken()
	tokenRes, err := g.Client().Post(ctx, accessTokenConfig.RequestUrl, accessTokenConfig)
	if err != nil {
		g.Log().Error(ctx, err.Error(), err)
		return
	}
	defer tokenRes.Close()
	defer tokenRes.Body.Close()
	// 读取响应体并且json转换到结构体
	accessTokenRes := &AccessTokenRes{}
	err = json.NewDecoder(tokenRes.Body).Decode(accessTokenRes)

	autoConfig := config.GetAutoConfig()
	code := g.Map{"code": req.Code}
	tokenUrl := autoConfig.RequestUrl + "?access_token=" + accessTokenRes.AccessToken

	// 请求腾讯开发者服务器,code换取openid
	client := g.Client()
	client.SetHeader("Content-Type", "application/json")
	phoneRes, err := client.Post(ctx, tokenUrl, code)
	if err != nil {
		g.Log().Error(ctx, err.Error(), err)
		return
	}
	defer phoneRes.Close()
	defer phoneRes.Body.Close()

	body := phoneRes.ReadAllString()
	// 将 JSON 数据解析为 g.Map
	var resultMap g.Map
	err = gconv.Struct(body, &resultMap)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 获取 phone_info map
	phoneInfo := resultMap["phone_info"].(map[string]interface{})
	// 获取 phoneNumber 字段的值
	phoneNumber := phoneInfo["phoneNumber"].(string)

	var (
		member     *model.BaseSysMember
		memberId   interface{}
		memberData g.Map
	)

	s.GetDao().Ctx(ctx).Where("username=?", phoneNumber).Where("status=?", 1).Scan(&member)
	if member == nil {
		md5password, _ := gmd5.Encrypt("123456")
		memberData = g.Map{
			"username":  phoneNumber,
			"password":  md5password,
			"levelName": "普通会员",
		}
		db := g.DB()
		if tx, err := db.Begin(ctx); err == nil {
			memberId, err = tx.Model(s.GetDao()).Data(memberData).InsertAndGetId()
			if err != nil {
				g.Log().Error(ctx, "memberData insert :"+err.Error(), err)
				err = gerror.New("写入失败")
				tx.Rollback()
				return nil, err
			}

			_, err = tx.Model(model.NewBaseMemberAttr()).Data(g.Map{"userId": memberId}).Insert()
			if err != nil {
				g.Log().Error(ctx, "NewBaseMemberAttr insert :"+err.Error(), err)
				err = gerror.New("写入失败")
				tx.Rollback()
				return nil, err
			}
			memberData["id"] = memberId
			gconv.Struct(memberData, &member)
			if err == nil {
				tx.Commit()
			}
		}

	}
	//else {
	//	err = dzhcore.DBM(m.Model).Where("username=?", phoneNumber).Scan(&member)
	//	if err != nil {
	//		err = gerror.New("查询失败")
	//		return
	//	}
	//}

	result, err = s.generateTokenByUser(ctx, member)
	if err != nil {
		return
	}

	return
}

// VerifyCount 验证游客次数
func (s *sBaseSysMemberLoginService) VerifyCount(ctx g.Ctx, req *v1.VerifyCountReq) (data interface{}, err error) {

	if req.Token == "" {
		err = gerror.New("token不能为空")
		g.Log().Error(ctx, "func VerifyCount :"+err.Error(), err)
		return
	}

	token, _ := dzhcore.CacheManager.Get(ctx, "member:token:"+gconv.String(req.Token))
	if token.IsEmpty() {
		err = gerror.New("账号过期，请重新登陆")
		g.Log().Error(ctx, "func VerifyCount :"+err.Error(), err)
		return
	}
	tokenCount, _ := dzhcore.CacheManager.Get(ctx, "member:token:count:"+gconv.String(req.Token))

	fmt.Printf("tokenCount: %v\n", tokenCount)
	if tokenCount.Int() >= 2 {
		err = gerror.New("超过体验次数，请购买套餐")
		g.Log().Error(ctx, "func VerifyCount :"+err.Error(), err)
		return
	}
	count := tokenCount.Int() + 1
	dzhcore.CacheManager.Set(ctx, "member:token:count:"+gconv.String(token), count, 0)

	return
}

// AccountLogin 账号登录
//func (s *sBaseSysMemberLoginService) AccountLogin(ctx g.Ctx, req *v1.AccountLoginReq) (result *v1.TokenRes, err error) {
//
//	var member *model.BaseSysMember
//	pw, _ := gmd5.Encrypt(req.PassWord)
//	s.GetDao().Ctx(ctx).Where("userName=?", req.UserName).Where("passWord=?", pw).Scan(&member)
//	if member == nil {
//		err = gerror.New("账号密码错误")
//		g.Log().Error(ctx, "func AccountLogin find :"+err.Error(), err)
//		return
//	}
//
//	// 生成token
//	result, err = s.generateTokenByUser(ctx, member)
//
//	if err != nil {
//		return
//	}
//
//	return
//}

// AccountRegister 账号注册
func (s *sBaseSysMemberLoginService) AccountRegister(ctx g.Ctx, req *v1.AccountRegisterReq) (result *v1.TokenRes, err error) {

	var (
		r    = g.RequestFromCtx(ctx)
		rmap = r.GetMap()
	)

	var member *model.BaseSysMember
	s.GetDao().Ctx(ctx).Where("userName=?", req.UserName).Scan(&member)
	if member != nil {
		err = gerror.New("账号已存在")
		g.Log().Error(ctx, "func AccountRegister exit :"+err.Error(), err)
		return
	}

	if rmap["passWord"] != rmap["rePassWord"] {
		err = gerror.New("两次密码不一致！")
		g.Log().Error(ctx, "AccountRegister:"+err.Error(), err)
		return
	}

	pwMd5, _ := gmd5.Encrypt(req.PassWord)
	insertData := g.Map{
		"userName": req.UserName,
		"passWord": pwMd5,
	}

	db := g.DB()
	if tx, err := db.Begin(ctx); err == nil {
		memberId, err := tx.Model(s.Dao).Data(insertData).InsertAndGetId()
		if err != nil {
			g.Log().Error(ctx, "AccountRegister insert:"+err.Error(), err)
			err = gerror.New("写入失败")
			tx.Rollback()
			return nil, err
		}
		_, err = tx.Model(model.NewBaseMemberAttr()).Data(g.Map{"userId": memberId}).Insert()
		if err != nil {
			g.Log().Error(ctx, "NewBaseMemberAttr insert :"+err.Error(), err)
			err = gerror.New("写入失败")
			tx.Rollback()
			return nil, err
		}

		if err == nil {
			tx.Commit()
		}
	}

	return
}

// 根据用户生成前端需要的Token信息
func (s *sBaseSysMemberLoginService) generateTokenByUser(ctx g.Ctx, member *model.BaseSysMember) (result *v1.TokenRes, err error) {
	// 获取用户角色
	roleIds := make([]uint, 1)
	roleIds[0] = 1

	result.Expire = config.Config.Jwt.Token.Expire
	result.RefreshExpire = config.Config.Jwt.Token.RefreshExpire
	result.Token = s.generateToken(ctx, member, roleIds, result.Expire, false)
	result.RefreshToken = s.generateToken(ctx, member, roleIds, result.RefreshExpire, true)
	// 将用户相关信息保存到缓存

	dzhcore.CacheManager.Set(ctx, "member:token:"+gconv.String(member.ID), result.Token, 0)
	dzhcore.CacheManager.Set(ctx, "member:token:refresh:"+gconv.String(member.ID), result.RefreshToken, 0)

	return
}

// generateToken  生成token
func (s *sBaseSysMemberLoginService) generateToken(ctx g.Ctx, member *model.BaseSysMember, roleIds []uint, exprire uint, isRefresh bool) (token string) {
	err := dzhcore.CacheManager.Set(ctx, "member:passwordVersion:"+gconv.String(member.ID), gconv.String(member.PasswordV), 0)
	if err != nil {
		g.Log().Error(ctx, "生成token失败", err)
	}
	type Claims struct {
		IsRefresh       bool   `json:"isRefresh"`
		RoleIds         []uint `json:"roleIds"`
		UserName        string `json:"username"`
		NickName        string `json:"nickname"`
		UserId          string `json:"userId"`
		PasswordVersion *int32 `json:"passwordVersion"`
		jwt.RegisteredClaims
	}
	claims := &Claims{
		IsRefresh:       isRefresh,
		RoleIds:         roleIds,
		UserName:        member.UserName,
		NickName:        member.Nickname,
		UserId:          member.ID,
		PasswordVersion: member.PasswordV,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(exprire) * time.Second)),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString([]byte(config.Config.Jwt.Secret))
	if err != nil {
		g.Log().Error(ctx, "生成token失败", err)
	}
	return
}
