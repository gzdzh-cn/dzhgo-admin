package sys

import (
	"context"
	"dzhgo/internal/common"
	"dzhgo/internal/dao"
	"dzhgo/internal/model"
	"dzhgo/internal/model/do"
	baseEntity "dzhgo/internal/model/entity"
	"dzhgo/internal/service"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gzdzh-cn/dzhcore"
)

func init() {
	service.RegisterBaseSysActionLogService(&sBaseSysActionLogService{})
}

type sBaseSysActionLogService struct {
	*dzhcore.Service
}

func NewBaseSysActionLogService() *sBaseSysActionLogService {
	return &sBaseSysActionLogService{
		&dzhcore.Service{
			Dao:   &dao.BaseSysActionLog,
			Model: model.NewBaseSysActionLog(),
			ListQueryOp: &dzhcore.QueryOp{
				FieldEQ:      []string{""},                      // 字段等于
				KeyWordField: []string{""},                      // 模糊搜索匹配的数据库字段
				AddOrderby:   g.MapStrStr{"createTime": "DESC"}, // 添加排序
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
				ModifyResult: func(ctx g.Ctx, data any) any { // 修改结果
					return data
				},
			},
			PageQueryOp: &dzhcore.QueryOp{
				FieldEQ:      []string{""},                      // 字段等于
				KeyWordField: []string{"name", "remark"},        // 模糊搜索匹配的数据库字段
				AddOrderby:   g.MapStrStr{"createTime": "DESC"}, // 添加排序
				Where: func(ctx context.Context) []g.Array { // 自定义条件
					admin := common.GetAdmin(ctx)
					whereData := []g.Array{}

					// 超管查看全部数据；非超管仅看自己的数据
					roleIdsGarray := garray.NewStrArrayFrom(admin.RoleIds)
					if roleIdsGarray.Contains("1") {
						return whereData
					}
					whereData = append(whereData, []g.Array{{"user_id", admin.UserId, true}}...)
					return whereData
				},
				OrWhere: func(ctx context.Context) []g.Array { // or 自定义条件
					rmap := g.RequestFromCtx(ctx).GetMap()
					kw := gconv.String(rmap["keyWord"])
					if kw == "" {
						return []g.Array{}
					}

					// 关键字存在时，去用户表模糊匹配 name/nickName/username，拿到 userId 列表
					var userList []*baseEntity.BaseSysUser
					err := dao.BaseSysUser.Ctx(ctx).
						Where("name LIKE ? OR username LIKE ? OR nickName LIKE ?", "%"+kw+"%", "%"+kw+"%", "%"+kw+"%").
						Scan(&userList)
					if err != nil {
						g.Log().Error(ctx, err.Error())
						return []g.Array{}
					}
					userIds := []string{}
					for _, u := range userList {
						userIds = append(userIds, u.Id)
					}
					if len(userIds) == 0 {
						return []g.Array{}
					}

					// 仅对超管放宽到按 userId 列表匹配；非超管仍受 Where 的本人限制
					admin := common.GetAdmin(ctx)
					roleIdsGarray := garray.NewStrArrayFrom(admin.RoleIds)
					if roleIdsGarray.Contains("1") {
						return []g.Array{{"user_id IN (?)", userIds}}
					}
					return []g.Array{}
				},
				Select: "",                  // 查询字段,多个字段用逗号隔开 如: id,name  或  a.id,a.name,b.name AS bname
				As:     "",                  //主表别名
				Join:   []*dzhcore.JoinOp{}, // 关联查询
				Extend: func(ctx g.Ctx, m *gdb.Model) *gdb.Model { // 追加其他条件
					return m
				},
				ModifyResult: func(ctx g.Ctx, data any) any { // 修改结果

					type List struct {
						*baseEntity.BaseSysActionLog
						UserName string `json:"userName"`
					}
					type Page struct {
						Pagination *common.Pagination `json:"pagination"`
						List       []*List            `json:"list"`
					}
					var (
						userNameMap = g.MapStrStr{} //id为下标的会员数集
						pageData    *Page
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

// 记录操作日志
func (s *sBaseSysActionLogService) Record(ctx context.Context, userId, name, remark string) (data any, err error) {

	do := do.BaseSysActionLog{
		Id:     dzhcore.NodeSnowflake.Generate().String(),
		UserId: userId,
		Name:   name,
		Remark: remark,
	}
	_, err = dao.BaseSysActionLog.Ctx(ctx).Data(do).Insert()
	if err != nil {
		return
	}

	return
}
