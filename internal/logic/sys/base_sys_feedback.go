package sys

import (
	"context"
	"dzhgo/internal/common"
	"dzhgo/internal/dao"
	"dzhgo/internal/model"
	"dzhgo/internal/model/do"
	"dzhgo/internal/model/entity"
	baseEntity "dzhgo/internal/model/entity"
	"dzhgo/internal/service"
	"fmt"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gzdzh-cn/dzhcore"
)

func init() {
	service.RegisterBaseSysFeedbackService(&sBaseSysFeedbackService{})
}

type sBaseSysFeedbackService struct {
	*dzhcore.Service
}

func NewBaseSysFeedbackService() *sBaseSysFeedbackService {
	return &sBaseSysFeedbackService{
		&dzhcore.Service{
			Dao:   &dao.BaseSysFeedback,
			Model: model.NewBaseSysFeedback(),
			ListQueryOp: &dzhcore.QueryOp{
				FieldEQ:      []string{""},                     // 字段等于
				KeyWordField: []string{""},                     // 模糊搜索匹配的数据库字段
				AddOrderby:   g.MapStrStr{"createTime": "ASC"}, // 添加排序
				Where: func(ctx context.Context) []g.Array { // 自定义条件
					return []g.Array{}
				},
				OrWhere: func(ctx context.Context) []g.Array { // or 自定义条件
					return []g.Array{}
				},
				Select: "",                  // 查询字段,多个字段用逗号隔开 如: id,name  或  a.id,a.name,b.name AS bname
				As:     "",                  //主表别名
				Join:   []*dzhcore.JoinOp{}, // 关联查询
				Extend: func(ctx g.Ctx, m *gdb.Model) *gdb.Model { // 追加其他条件
					return m
				},
				ModifyResult: func(ctx g.Ctx, data interface{}) interface{} { // 修改结果
					return data
				},
			},
			PageQueryOp: &dzhcore.QueryOp{
				FieldEQ:      []string{""},                      // 字段等于
				KeyWordField: []string{"title", "remark"},       // 模糊搜索匹配的数据库字段
				AddOrderby:   g.MapStrStr{"createTime": "DESC"}, // 添加排序
				Where: func(ctx context.Context) []g.Array { // 自定义条件
					where := []g.Array{}
					admin := common.GetAdmin(ctx)
					if !garray.NewStrArrayFrom(admin.RoleIds).Contains("1") {
						where = append(where, g.Array{"user_id = ?", admin.UserId})
					}
					return where
				},
				OrWhere: func(ctx context.Context) []g.Array { // or 自定义条件
					return []g.Array{}
				},
				Select: "",                  // 查询字段,多个字段用逗号隔开 如: id,name  或  a.id,a.name,b.name AS bname
				As:     "",                  //主表别名
				Join:   []*dzhcore.JoinOp{}, // 关联查询
				Extend: func(ctx g.Ctx, m *gdb.Model) *gdb.Model { // 追加其他条件
					return m
				},
				ModifyResult: func(ctx g.Ctx, data interface{}) interface{} { // 修改结果

					type List struct {
						*baseEntity.BaseSysFeedback
						UserName string `json:"userName"`
					}
					type Page struct {
						Pagination *common.Pagination `json:"pagination"`
						List       []*List            `json:"list"`
					}
					var (
						userNameMap = g.MapStrStr{} //id为下标的会员数集
						pageData    Page
					)

					//id为下标的会员数集
					{
						var userList []*baseEntity.BaseSysUser
						err := dao.BaseSysUser.Ctx(ctx).Fields("id", "name").Scan(&userList)
						if err != nil {
							g.Log().Error(ctx, err.Error())
							return err
						}
						for _, user := range userList {
							userNameMap[user.Id] = user.Name
						}
					}

					err := gconv.Struct(data, &pageData)
					if err != nil {
						g.Log().Error(ctx, err.Error())
						return err
					}
					if len(pageData.List) > 0 {
						for _, v := range pageData.List {
							if v != nil && v.UserId != "" {
								if userName, ok := userNameMap[v.UserId]; ok {
									v.UserName = userName
								}
							}
						}
					}

					data = pageData

					return data

				},
			},
			InsertParam: func(ctx context.Context) g.MapStrAny { // Add时插入参数
				return g.MapStrAny{}
			},
			Before: func(ctx context.Context) (err error) { // CRUD前的操作
				return nil
			},
			InfoIgnoreProperty: "",            // Info时忽略的字段,多个字段用逗号隔开
			UniqueKey:          g.MapStrStr{}, // 唯一键 key:字段名 value:错误信息
			NotNullKey:         g.MapStrStr{}, // 非空键 key:字段名 value:错误信息
		},
	}
}

func (s *sBaseSysFeedbackService) ServiceAdd(ctx context.Context, req *dzhcore.AddReq) (data any, err error) {

	admin := common.GetAdmin(ctx)
	rmap := g.RequestFromCtx(ctx).GetMap()

	insertData := do.BaseSysFeedback{
		Id:       dzhcore.NodeSnowflake.Generate().String(),
		UserId:   admin.UserId,
		FeType:   gconv.String(rmap["feType"]),
		Title:    gconv.String(rmap["title"]),
		Remark:   gconv.String(rmap["remark"]),
		Priority: gconv.String(rmap["priority"]),
		Img:      gconv.String(rmap["img"]),
	}
	_, err = dao.BaseSysFeedback.Ctx(ctx).Data(insertData).Insert()
	if err != nil {
		return nil, err
	}

	var user *entity.BaseSysUser
	err = dao.BaseSysUser.Ctx(ctx).Where("id = ?", admin.UserId).Scan(&user)
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return nil, err
	}

	common.RecordActionLog(ctx, "反馈", admin.UserId, fmt.Sprintf("工单：用户:%s,反馈类型:%s,标题:%s,内容:%s", insertData.FeType, user.Name, insertData.Title, insertData.Remark))

	// 发送通知
	service.BaseSysNoticeService().NoticeAdd(ctx, &entity.BaseSysNotice{
		Title:  "工单提醒",
		Remark: fmt.Sprintf("工单：用户:%s,反馈类型:%s,标题:%s,内容:%s", insertData.FeType, user.Name, insertData.Title, insertData.Remark),
		NoType: "safe",
	}, nil)

	return nil, nil
}

// ServiceInfo 获取反馈信息
func (s *sBaseSysFeedbackService) ServiceInfo(ctx context.Context, req *dzhcore.InfoReq) (data any, err error) {
	feedback := &baseEntity.BaseSysFeedback{}
	err = dao.BaseSysFeedback.Ctx(ctx).Where("id = ?", req.Id).Scan(feedback)
	if err != nil {
		return nil, err
	}

	return feedback, nil
}
