package defineType

import "github.com/gogf/gf/v2/os/gtime"

type MemberInfo struct {
	Id            string      `json:"id"            orm:"id"            ` //
	CreateTime    *gtime.Time `json:"createTime"    orm:"createTime"    ` // 创建时间
	UpdateTime    *gtime.Time `json:"updateTime"    orm:"updateTime"    ` // 更新时间
	AvatarUrl     string      `json:"avatarUrl"     orm:"avatarUrl"     ` // 头像
	Username      string      `json:"username"      orm:"username"      ` // 会员账号
	Nickname      string      `json:"nickname"      orm:"nickname"      ` // 会员昵称
	LevelName     string      `json:"levelName"     orm:"levelName"     ` // 等级名称
	Level         int64       `json:"level"         orm:"level"         ` // 等级
	Sex           int         `json:"sex"           orm:"sex"           ` // 性别
	Qq            string      `json:"qq"            orm:"qq"            ` // qq
	Mobile        string      `json:"mobile"        orm:"mobile"        ` // 手机号
	Wx            string      `json:"wx"            orm:"wx"            ` // 微信号
	WxImg         string      `json:"wxImg"         orm:"wxImg"         ` // 微信二维码
	Email         string      `json:"email"         orm:"email"         ` // email
	Role          string      `json:"role"          orm:"role"          ` // 家庭角色
	SessionKey    string      `json:"sessionKey"    orm:"session_key"   ` // session_key
	Remark        string      `json:"remark"        orm:"remark"        ` // 备注
	Status        int         `json:"status"        orm:"status"        ` //
	Description   string      `json:"description"   orm:"description"   ` // 描述
	LastLoginTime *gtime.Time `json:"lastLoginTime" orm:"lastLoginTime" ` // 最后登录时间

	Type           int    `json:"type"           orm:"type"            ` // 类型:1=公众号,2=小程序
	UserId         string `json:"userId"         orm:"user_id"         ` // 用户ID
	Notify         int64  `json:"notify"         orm:"notify"          ` // 微信通知:1=是,0=否
	Language       string `json:"language"       orm:"language"        ` // 语言
	Country        string `json:"country"        orm:"country"         ` // 国家
	Province       string `json:"province"       orm:"province"        ` // 省份
	City           string `json:"city"           orm:"city"            ` // 城市
	Headimgurl     string `json:"headimgurl"     orm:"headimgurl"      ` // 头像
	Subscribe      int64  `json:"subscribe"      orm:"subscribe"       ` // 关注状态:0=未关注,1=已关注
	SubscribeTime  int64  `json:"subscribeTime"  orm:"subscribe_time"  ` // 关注时间
	Groupid        int64  `json:"groupid"        orm:"groupid"         ` // 分组ID
	TagidList      string `json:"tagidList"      orm:"tagid_list"      ` // 标签列表
	Privilege      string `json:"privilege"      orm:"privilege"       ` // 权限
	SubscribeScene string `json:"subscribeScene" orm:"subscribe_scene" ` // 关注来源
	QrScene        string `json:"qrScene"        orm:"qr_scene"        ` // 扫码场景
	QrSceneStr     string `json:"qrSceneStr"     orm:"qr_scene_str"    ` // 扫码场景描述
}

// 微信公众号配置
type WxConfig struct {
	Appid     string `json:"appid"`
	Secret    string `json:"secret"`
	GrantType string `json:"grant_type"`
}

// 微信公众号token返回数据
type WxMpTokenResponse struct {
	AccessToken    string `json:"access_token"`
	ExpiresIn      int    `json:"expires_in"`
	RefreshToken   string `json:"refresh_token"`
	Openid         string `json:"openid"`
	Scope          string `json:"scope"`
	IsSnapshotuser int    `json:"is_snapshotuser"`
	Unionid        string `json:"unionid"`
}

// 微信公众号用户解密返回数据
type WxMpUserInfoResponse struct {
	Openid     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}

// 微信公众号获取accessToken返回数据
type WxMpApiAccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// 微信手机号授权
type AutoPhone struct {
	RequestUrl string `json:"requestUrl"`
	Code       string `json:"code"`
}

//type AccessToken struct {
//	RequestUrl string `json:"requestUrl"`
//	GrantType  string `json:"grant_type"`
//	Appid      string `json:"appid"`
//	Secret     string `json:"secret"`
//}

//// 小程序
//type MinConfig struct {
//	RequestUrl string `json:"requestUrl"`
//	Appid      string `json:"appid"`
//	Secret     string `json:"secret"`
//	GrantType  string `json:"grant_type"`
//}
