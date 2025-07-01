package sys

import (
	"context"
	v1 "dzhgo/addons/member/api/app_v1"
	"dzhgo/addons/member/consts"
	"dzhgo/addons/member/dao"
	memberDefineType "dzhgo/addons/member/defineType"
	"dzhgo/addons/member/model"
	"dzhgo/addons/member/model/entity"
	"dzhgo/addons/member/service"
	"dzhgo/internal/config"
	"dzhgo/internal/defineType"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gzdzh-cn/dzhcore"
	"github.com/gzdzh-cn/dzhcore/utility/util"
)

func init() {
	service.RegisterMemberManageService(&sMemberManageService{})
}

type sMemberManageService struct {
	*dzhcore.Service
}

func NewsMemberManageService() *sMemberManageService {
	return &sMemberManageService{
		&dzhcore.Service{
			Dao:   &dao.AddonsMemberManage,
			Model: model.NewMemberManage(),
			PageQueryOp: &dzhcore.QueryOp{
				FieldEQ:      []string{},
				KeyWordField: []string{},
				AddOrderby:   g.MapStrStr{},
				Select:       "addons_member_manage.*,member_attr.type,member_attr.user_id,member_attr.notify,member_attr.country,member_attr.province,member_attr.city",
				Join: []*dzhcore.JoinOp{
					{
						Model:     model.NewMemberAttr(),
						Alias:     "member_attr",
						Type:      "LeftJoin",
						Condition: "member_attr.user_id = addons_member_manage.`id`",
					},
				},
			},
		},
	}
}

// 新增|删除|修改前的操作
func (s *sMemberManageService) ModifyBefore(ctx context.Context, method string, param g.MapStrAny) (err error) {

	return
}

// 新增|删除|修改后的操作
func (s *sMemberManageService) ModifyAfter(ctx context.Context, method string, param g.MapStrAny) (err error) {
	return
}

// 新增
func (s *sMemberManageService) ServiceAdd(ctx context.Context, req *dzhcore.AddReq) (data interface{}, err error) {

	var (
		m          = dao.AddonsMemberManage.Ctx(ctx)
		r          = g.RequestFromCtx(ctx)
		rmap       = r.GetMap()
		memberAttr *entity.AddonsMemberAttr
	)

	// 如果reqmap["password"]不为空，则对密码进行md5加密
	if !r.Get("password").IsNil() {
		rmap["password"] = gmd5.MustEncryptString(r.Get("password").String())
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		rmap["id"] = dzhcore.NodeSnowflake.Generate().String()
		_, err = m.TX(tx).Data(rmap).Insert()
		if err != nil {
			g.Log().Error(ctx, err)
			return err
		}

		err = gconv.Struct(rmap, &memberAttr)
		if err != nil {
			g.Log().Error(ctx, err)
			return err
		}
		memberAttr.Id = dzhcore.NodeSnowflake.Generate().String()
		memberAttr.UserId = gconv.String(rmap["id"])
		_, err = dao.AddonsMemberAttr.Ctx(ctx).Data(memberAttr).OmitEmpty().Insert()
		if err != nil {
			g.Log().Error(ctx, err)
			return err
		}

		data = g.Map{"id": rmap["id"]}
		return
	})
	return
}

// 查询
func (s *sMemberManageService) ServiceInfo(ctx context.Context, req *dzhcore.InfoReq) (data interface{}, err error) {

	rmap := g.RequestFromCtx(ctx).GetMap()

	m := dao.AddonsMemberManage.Ctx(ctx).As("m").Fields("m.*,member_attr.type,member_attr.user_id,member_attr.notify,member_attr.country,member_attr.province,member_attr.city")
	m = m.LeftJoin("addons_member_attr member_attr", "member_attr.user_id = m.id")
	one, err := m.Where("m.id", rmap["id"]).One()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	var member *memberDefineType.MemberInfo
	err = gconv.Struct(one, &member)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	data = member
	return
}

// 账号登录
func (s *sMemberManageService) AccountLogin(ctx context.Context, req *v1.AccountLoginReq) (data interface{}, err error) {

	pw, _ := gmd5.Encrypt(req.PassWord)
	one, err := dao.AddonsMemberManage.Ctx(ctx).Where("username=?", req.UserName).Where("password=?", pw).One()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if one == nil {
		err = gerror.New("账号密码错误")
		g.Log().Error(ctx, err)
		return
	}

	var member *entity.AddonsMemberManage
	err = gconv.Struct(one, &member)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	//更新登录时间
	member.LastLoginTime = gtime.Now()
	_, err = dao.AddonsMemberManage.Ctx(ctx).Where("id", member.Id).Data(member).Update()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	// 生成token
	data, err = s.GenerateTokenByUser(ctx, member)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	return
}

// 公众号登录
func (s *sMemberManageService) MpLogin(ctx context.Context, req *v1.MpLoginReq) (data interface{}, err error) {

	type Member struct {
		*entity.AddonsMemberManage
		*entity.AddonsMemberAttr
	}

	//微信配置
	var (
		manage = dao.AddonsMemberManage
		attr   = dao.AddonsMemberAttr
	)

	//获取配置
	var BaseConfig = config.NewBaseConfig()
	wxConfig := BaseConfig.WxConfig
	if wxConfig == nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	//通过code获取token
	wxMpTokenResponse, err := s.GetWxAccessToken(ctx, req.Code, wxConfig)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	// 获取解密后的数据
	wxMpUserInfoResponse, err := s.GetUserInfo(ctx, wxMpTokenResponse)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	member := &Member{
		AddonsMemberManage: &entity.AddonsMemberManage{
			Openid:    wxMpTokenResponse.Openid,
			UnionId:   wxMpTokenResponse.Unionid,
			Nickname:  wxMpUserInfoResponse.Nickname,
			Sex:       wxMpUserInfoResponse.Sex,
			AvatarUrl: wxMpUserInfoResponse.Headimgurl,
		},
		AddonsMemberAttr: &entity.AddonsMemberAttr{
			Openid:     wxMpUserInfoResponse.Openid,
			Sex:        wxMpUserInfoResponse.Sex,
			Province:   wxMpUserInfoResponse.Province,
			City:       wxMpUserInfoResponse.City,
			Country:    wxMpUserInfoResponse.Country,
			Headimgurl: wxMpUserInfoResponse.Headimgurl,
			Notify:     req.Notify,
		},
	}

	//	1、有UserId，微信登录绑定到账号，没有UserId就不绑定账号

	//不绑定账号
	if req.UserId == nil {

		// 判断openId是否存在
		//err = manage.Ctx(ctx).Where("openid", wxMpTokenResponse.Openid).Scan(member.AddonsMemberManage)
		//if err != nil {
		//	return nil, err
		//}

		memberManage, err := manage.Ctx(ctx).Where("openid", wxMpTokenResponse.Openid).One()
		if err != nil {
			g.Log().Error(ctx, err)
			return nil, err
		}
		if memberManage != nil {
			err = gconv.Struct(memberManage, member.AddonsMemberManage)
			if err != nil {
				g.Log().Error(ctx, err)
				return nil, err
			}
		}
	}

	err = manage.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		//不绑定账号并且不存在openId的会员，直接新增一个微信登录的会员
		if req.UserId == nil && member.AddonsMemberManage.Id == "" {

			//主表
			member.AddonsMemberManage.Id = dzhcore.NodeSnowflake.Generate().String()
			member.AddonsMemberManage.Password = gmd5.MustEncryptString("123456")
			_, err = tx.Model(manage.Table()).Data(member.AddonsMemberManage).OmitEmpty().Insert()
			if err != nil {
				g.Log().Error(ctx, err)
				return err
			}

			//副表
			member.AddonsMemberAttr.Id = dzhcore.NodeSnowflake.Generate().String()
			member.AddonsMemberAttr.UserId = member.AddonsMemberManage.Id
			_, err = tx.Model(attr.Table()).Data(member.AddonsMemberAttr).OmitEmpty().Insert()
			if err != nil {
				g.Log().Error(ctx, err)
				return err
			}
		} else {

			var (
				manageWhereData g.Map
				attrWhereData   g.Map
				id              string
			)
			member.AddonsMemberManage.LastLoginTime = gtime.Now()

			//不绑定账号并且会员openId已存在，更新到openId的账号
			if req.UserId == nil && member.AddonsMemberManage.Id != "" {
				id = member.AddonsMemberManage.Id
			}

			//微信数据直接更新到指定的账号
			if req.UserId != nil {
				id = gconv.String(req.UserId)
				member.AddonsMemberManage.Id = id
			}

			manageWhereData = g.Map{
				"id": id,
			}
			attrWhereData = g.Map{
				"user_id": id,
			}

			//主表
			_, err = tx.Model(manage.Table()).Where(manageWhereData).Data(member.AddonsMemberManage).OmitEmpty().Update()
			if err != nil {
				g.Log().Error(ctx, err)
				return err
			}

			count, err := tx.Model(attr.Table()).Where("user_id = ?", id).Count()
			if err != nil {
				g.Log().Error(ctx, err)
				return err
			}
			if count == 0 {
				//副表
				member.AddonsMemberAttr.Id = dzhcore.NodeSnowflake.Generate().String()
				member.AddonsMemberAttr.UserId = id
				_, err = tx.Model(attr.Table()).Where(attrWhereData).Data(member.AddonsMemberAttr).OmitEmpty().Insert()
				if err != nil {
					g.Log().Error(ctx, err)
					return err
				}
			} else {
				//副表
				_, err = tx.Model(attr.Table()).Where(attrWhereData).Data(member.AddonsMemberAttr).OmitEmpty().Update()
				if err != nil {
					g.Log().Error(ctx, err)
					return err
				}
			}

		}

		// 生成token
		data, err = s.GenerateTokenByUser(ctx, member.AddonsMemberManage)
		if err != nil {
			g.Log().Error(ctx, err)
			return err
		}

		//存access_token
		err = dzhcore.CacheManager.Set(ctx, "member:access_token:"+gconv.String(member.AddonsMemberManage.Id), wxMpTokenResponse.AccessToken, time.Duration(wxMpTokenResponse.ExpiresIn)*time.Second)
		if err != nil {
			g.Log().Error(ctx, err)
			return err
		}

		return err
	})

	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	return
}

// 通过code获取token
func (s *sMemberManageService) GetWxAccessToken(ctx context.Context, code string, wxConfig *defineType.WxConfig) (data *memberDefineType.WxMpTokenResponse, err error) {

	//通过code换取网页授权access_token
	//https://api.weixin.qq.com/sns/oauth2/access_token?appid=APPID&secret=SECRET&code=CODE&grant_type=authorization_code
	url := consts.WxMpTokenUrl
	postData := fmt.Sprintf(`appid=%v&secret=%v&code=%v&grant_type=%v`, wxConfig.Appid, wxConfig.Secret, code, wxConfig.GrantType)
	if wxConfig.Appid == "" || wxConfig.Secret == "" || code == "" {
		err = gerror.New("缺少配置参数")
		if err != nil {
			g.Log().Error(ctx, err)
			return nil, err
		}
		g.Log().Error(ctx, err)
		return nil, err
	}

	wxMpTokenResponse := &memberDefineType.WxMpTokenResponse{}
	header := g.MapStrStr{"content-type": "application/x-www-form-urlencoded"}
	err = util.NewHttpClient().Post(ctx, url, header, postData, wxMpTokenResponse)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	data = wxMpTokenResponse
	return
}

// 获取解密后的数据
func (s *sMemberManageService) GetUserInfo(ctx context.Context, wxMpTokenResponse *memberDefineType.WxMpTokenResponse) (data *memberDefineType.WxMpUserInfoResponse, err error) {

	//拉取用户信息(需scope为 snsapi_userinfo)
	//https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
	url := consts.WxMpUserInfoUrl
	postData := fmt.Sprintf(`access_token=%v&openid=%v&lang=zh_CN`, wxMpTokenResponse.AccessToken, wxMpTokenResponse.Openid)
	if wxMpTokenResponse.AccessToken == "" || wxMpTokenResponse.Openid == "" {
		err = gerror.New("缺少配置参数")
		if err != nil {
			g.Log().Error(ctx, err)
			return nil, err
		}
		g.Log().Error(ctx, err)
		return nil, err
	}

	wxMpUserInfoResponse := &memberDefineType.WxMpUserInfoResponse{}
	err = util.NewHttpClient().Post(ctx, url, nil, postData, wxMpUserInfoResponse)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	data = wxMpUserInfoResponse
	return
}

// Person 方法 返回不带密码的用户信息
func (s *sMemberManageService) Person(ctx context.Context, userId string) (data interface{}, err error) {

	m := dao.AddonsMemberManage.Ctx(ctx).As("m").Fields("m.*,member_attr.type,member_attr.user_id,member_attr.notify,member_attr.country,member_attr.province,member_attr.city")
	m = m.LeftJoin("addons_member_attr member_attr", "member_attr.user_id = m.id")
	m = m.Where("m.id = ?", userId)
	one, err := m.One()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if one == nil {
		err = gerror.New("会员获取数据出错")
		g.Log().Error(ctx, err)
		return nil, err
	}

	var member *memberDefineType.MemberInfo
	err = gconv.Struct(one, &member)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	data = member

	return
}

// 根据用户生成前端需要的Token信息
func (s *sMemberManageService) GenerateTokenByUser(ctx g.Ctx, member *entity.AddonsMemberManage) (result *v1.TokenRes, err error) {
	// 获取用户角色
	//roleIds := make([]uint, 1)
	roleIds := g.SliceStr{"1"}

	// 生成token
	result = &v1.TokenRes{}
	result.Expire = config.Cfg.Jwt.Token.Expire
	result.RefreshExpire = config.Cfg.Jwt.Token.RefreshExpire
	result.Token = s.generateToken(ctx, member, roleIds, result.Expire, false)
	result.RefreshToken = s.generateToken(ctx, member, roleIds, result.RefreshExpire, true)
	// 将用户相关信息保存到缓存

	dzhcore.CacheManager.Set(ctx, "member:token:"+gconv.String(member.Id), result.Token, 0)
	dzhcore.CacheManager.Set(ctx, "member:token:refresh:"+gconv.String(member.Id), result.RefreshToken, 0)

	return
}

// generateToken  生成token
func (s *sMemberManageService) generateToken(ctx g.Ctx, member *entity.AddonsMemberManage, roleIds []string, expire uint, isRefresh bool) (token string) {
	err := dzhcore.CacheManager.Set(ctx, "member:passwordVersion:"+gconv.String(member.Id), gconv.String(member.PasswordV), 0)
	if err != nil {
		g.Log().Error(ctx, "生成token失败", err)
	}

	passwordVersion := gconv.Int32(member.PasswordV)
	claims := &dzhcore.Claims{
		IsRefresh:       isRefresh,
		RoleIds:         roleIds,
		Username:        member.Username,
		UserId:          member.Id,
		PasswordVersion: &passwordVersion,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expire) * time.Second)),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString([]byte(config.Cfg.Jwt.Secret))
	if err != nil {
		g.Log().Error(ctx, "生成token失败", err)
		return
	}
	return
}
