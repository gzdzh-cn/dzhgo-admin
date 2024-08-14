package types

type Menu struct {
	CRM       string `json:"crm"`
	Order     string `json:"order"`
	Icon      string `json:"icon"`
	IsShow    bool   `json:"isShow"`
	KeepAlive bool   `json:"keepAlive"`
	Module    string `json:"module"`
	Name      string `json:"name"`
	OrderNum  int    `json:"orderNum"`
	ParentID  string `json:"parentId"`
	Router    string `json:"router"`
	Type      int    `json:"type"`
	ViewPath  string `json:"viewPath"`
}

type Perms struct {
	Id        string `json:"id"`
	Type      int    `json:"type"`
	ParentId  string `json:"parentId"`
	Name      string `json:"name"`
	Perms     string `json:"perms"`
	IsInstall bool   `json:"isInstall"`
}

type Member struct {
	Id          int     `json:"id"`
	AvatarUrl   string  `json:"avatarUrl"`
	Username    string  `json:"username"`
	NickName    string  `json:"nickName"`
	LevelName   string  `json:"levelName"`
	Level       int     `json:"level"`
	Sex         *int32  `json:"sex"`
	QQ          *string `json:"qq"`
	Mobile      *string `json:"mobile"`
	Wx          *string `json:"wx"`
	WxImg       *string `json:"wxImg"`
	Email       *string `json:"email"`
	Description *string `json:"description"`
}

type MemberSum struct {
	FansCount    int `json:"fansCount"`
	FollowCount  int `json:"followCount"`
	LikeCount    int `json:"likeCount"`
	CollectCount int `json:"collectCount"`
}

type MemberAttr struct {
	UserId          *string `json:"userId"`
	FollowJson      string  `json:"followJson"`
	LikeJson        string  `json:"likeJson"`
	CollectJson     string  `json:"collectJson"`
	CommentLikeJson string  `json:"commentLikeJson"`
	NoticeJson      string  `json:"noticeJson"`
}

// 公共站点配置
type CommonSetting struct {
	Logo     string `json:"logo"`
	SiteName string `json:"siteName"`
}
