package sys

import (
	"context"
	"dzhgo/internal/common"
	"dzhgo/internal/dao"
	"dzhgo/internal/model"
	baseEntity "dzhgo/internal/model/entity"
	"dzhgo/internal/service"
	"fmt"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gzdzh-cn/dzhcore"
)

func init() {
	service.RegisterBaseSysUserService(NewsBaseSysUserService())
}

type sBaseSysUserService struct {
	*dzhcore.Service
}

// NewsBaseSysUserService 创建一个新的sBaseSysUserService实例
func NewsBaseSysUserService() *sBaseSysUserService {
	return &sBaseSysUserService{
		Service: &dzhcore.Service{
			Dao:                &dao.BaseSysUser,
			Model:              model.NewBaseSysUser(),
			InfoIgnoreProperty: "password",
			UniqueKey: map[string]string{
				"username": "用户名不能重复",
			},
			ListQueryOp: &dzhcore.QueryOp{
				Select:  "user.*",
				FieldEQ: []string{"password"},
				As:      "user",
				Join:    []*dzhcore.JoinOp{},
				Where: func(ctx context.Context) []g.Array {

					r := g.RequestFromCtx(ctx)
					var where = []g.Array{}
					if r.GetCtxVar("admin").String() != "" {

						//  找出全部管理员id
						adminList, err := dao.BaseSysUserRole.Ctx(ctx).Where("roleId", "1").Fields("userId").Array()
						if err != nil {
							g.Log().Error(ctx, err)
							return where
						}
						// 排除 adminList 中的 id
						where = append(where, g.Array{"user.id NOT IN (?)", adminList})
					}

					return where
				},
				Extend: func(ctx g.Ctx, m *gdb.Model) *gdb.Model {
					return m
				},
				KeyWordField: []string{},
				ModifyResult: func(ctx g.Ctx, data interface{}) interface{} {

					return data
				},
			},
			PageQueryOp: &dzhcore.QueryOp{
				Select:  "user.*,dept.`name` as departmentName,GROUP_CONCAT( role.`name` ) AS `roleName`",
				FieldEQ: []string{"password"},
				As:      "user",
				Join: []*dzhcore.JoinOp{
					{
						Model:     model.NewBaseSysDepartment(),
						Alias:     "dept",
						Type:      "LeftJoin",
						Condition: "`user`.`departmentId` = `dept`.`id`",
					},
					{
						Model:     model.NewBaseSysUserRole(),
						Alias:     "user_role",
						Type:      "LeftJoin",
						Condition: "`user`.`id` = `user_role`.`userId`",
					},
					{
						Model:     model.NewBaseSysRole(),
						Alias:     "`role`",
						Type:      "LeftJoin",
						Condition: "`role`.`id` = `user_role`.`roleId`",
					},
				},
				Where: func(ctx context.Context) []g.Array {

					r := g.RequestFromCtx(ctx)
					rMap := r.GetMap()
					admin := common.GetAdmin(ctx)
					var where = []g.Array{}
					if r.GetCtxVar("admin").String() != "" {
						// where = []g.Array{{"(user.departmentId IN (?))", gconv.SliceStr(rMap["departmentIds"])}}
						where = append(where, g.Array{"(user.departmentId IN (?))", gconv.SliceStr(rMap["departmentIds"])})
						if gstr.Equal(admin.UserId, "1152921504606846975") {
							where = append(where, g.Slice{"user.id NOT IN (?)", g.Slice{"1152921504606846976"}})
						}

						if !gstr.Equal(admin.UserId, "1152921504606846975") && !gstr.Equal(admin.UserId, "1152921504606846976") {
							where = append(where, g.Slice{"user.id NOT IN (?)", g.Slice{"1152921504606846975", "1152921504606846976"}})
						}
						//  找出全部管理员id
						adminList, err := dao.BaseSysUserRole.Ctx(ctx).Where("roleId", "1").Fields("userId").Array()
						if err != nil {
							g.Log().Error(ctx, err)
							return where
						}
						// 排除 adminList 中的 id
						where = append(where, g.Array{"user.id NOT IN (?)", adminList})
					}

					return where
				},
				Extend: func(ctx g.Ctx, m *gdb.Model) *gdb.Model {
					return m.Group("`user`.`id`")
				},
				KeyWordField: []string{"user.name", "user.username", "user.nickName"},
				ModifyResult: func(ctx g.Ctx, data interface{}) interface{} {
					type List struct {
						*baseEntity.BaseSysUser
						DepartmentName string   `json:"departmentName"`
						RoleName       string   `json:"roleName"`
						RoleIdList     []string `json:"roleIdList"`
					}

					type Pagination struct {
						Page  int `json:"page"`
						Size  int `json:"size"`
						Total int `json:"total"`
					}
					type PageData struct {
						List       []*List     `json:"list"`
						Pagination *Pagination `json:"pagination"`
					}

					var (
						userRoleList []*baseEntity.BaseSysUserRole
						userMap      = make(map[string][]string)
					)
					err := dao.BaseSysUserRole.Ctx(ctx).Scan(&userRoleList)
					if err != nil {
						return err
					}

					//会员id为key，会员roleid数组为value
					for _, userRoleRow := range userRoleList {
						userMap[userRoleRow.UserId] = append(userMap[userRoleRow.UserId], userRoleRow.RoleId)
					}

					list := gconv.Map(data)["list"]
					if len(gconv.SliceAny(list)) > 0 {
						pageData := &PageData{}
						_ = gconv.Struct(data, pageData)

						if pageData != nil && len(pageData.List) > 0 {
							for _, row := range pageData.List {
								row.RoleIdList = userMap[row.Id]
							}
						}
						data = pageData
					}

					return data
				},
			},
		},
	}
}

// Person 方法 返回不带密码的用户信息
func (s *sBaseSysUserService) Person(userId string) (res interface{}, err error) {

	type Person struct {
		Id           string      `json:"id"           orm:"id"           ` //
		CreateTime   *gtime.Time `json:"createTime"   orm:"createTime"   ` // 创建时间
		UpdateTime   *gtime.Time `json:"updateTime"   orm:"updateTime"   ` // 更新时间
		DepartmentId string      `json:"departmentId" orm:"departmentId" ` //
		Name         string      `json:"name"         orm:"name"         ` //
		Username     string      `json:"username"     orm:"username"     ` //
		NickName     string      `json:"nickName"     orm:"nickName"     ` //
		HeadImg      string      `json:"headImg"      orm:"headImg"      ` //
		Phone        string      `json:"phone"        orm:"phone"        ` //
		Email        string      `json:"email"        orm:"email"        ` //
		Status       int         `json:"status"       orm:"status"       ` //
		Remark       string      `json:"remark"       orm:"remark"       ` //
		SocketId     string      `json:"socketId"     orm:"socketId"     ` //
		RoleIds      string      `json:"roleIds"`
	}
	var personData *Person
	m := s.Dao.Ctx(ctx).As("user").Fields("user.*,GROUP_CONCAT(r.id) AS roleIds")
	m = m.LeftJoin("base_sys_user_role u_r", "user.id = u_r.userId")
	m = m.LeftJoin("base_sys_role r", "r.id = u_r.roleId")
	m = m.Where("user.id = ?", userId).Group("user.id")
	err = m.Scan(&personData)
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return nil, err
	}

	//if personData != nil {
	//	roleIds := baseCommon.GetAdmin(ctx).RoleIds
	//	roleIdsGarray := garray.NewStrArrayFrom(roleIds)
	//	if roleIdsGarray.Contains("1") {
	//		personData.RoleIds = "1"
	//	}
	//}

	return personData, nil
}

func (s *sBaseSysUserService) ModifyBefore(ctx context.Context, method string, param g.MapStrAny) (err error) {

	if method == "Update" || method == "Add" {
		count, err := dao.BaseSysUser.Ctx(ctx).WhereNot("id", gconv.String(param["id"])).Where("username", gconv.String(param["username"])).Count()
		if err != nil {
			g.Log().Error(ctx, err.Error())
			return err
		}
		if count > 0 {
			err = gerror.New("用户名不能重复，请重新填写")
			g.Log().Error(ctx, err)
			return err
		}
		return nil
	}

	if method == "Delete" {
		// 禁止删除超级管理员
		userIds := garray.NewStrArrayFrom(gconv.Strings(param["ids"]))

		if userIds.Len() == 1 {
			currentId, found := userIds.Get(0)
			if found == false {
				err = gerror.New("没有找到会员")
				return err
			}
			roleIdArray, err := dao.BaseSysUserRole.Ctx(ctx).Where("userId", currentId).Fields("roleId").Array()
			if err != nil {
				return err
			}

			if len(roleIdArray) > 0 {
				roleIdGarray := garray.NewStrArrayFrom(gconv.SliceStr(roleIdArray))
				if roleIdGarray.Contains("1") {
					err = gerror.New("超级管理员不能删除")
					return err
				}
			}
		}
		if userIds.Len() > 0 {
			var (
				userRoleList []*baseEntity.BaseSysUserRole
				userMap      = make(map[string][]string)
			)
			err := dao.BaseSysUserRole.Ctx(ctx).Scan(&userRoleList)
			if err != nil {
				return err
			}

			//会员id为key，会员roleid数组为value
			for _, userRoleRow := range userRoleList {
				userMap[userRoleRow.UserId] = append(userMap[userRoleRow.UserId], userRoleRow.RoleId)
			}

			for _, userId := range userIds.Slice() {
				if len(userMap[userId]) == 1 {
					if garray.NewStrArrayFrom(userMap[userId]).Contains("1") {
						err = gerror.New("超级管理员不能删除")
						return err
					}
				}
			}
		}

		userId, err := dao.BaseSysUserRole.Ctx(ctx).Where("roleId", "1").Value("userId")
		if err != nil {
			return err
		}

		// 排除掉超级管理员
		userIds.RemoveValue(userId.String())
		g.RequestFromCtx(ctx).SetParam("ids", userIds.Slice())

	}
	return
}

func (s *sBaseSysUserService) ModifyAfter(ctx context.Context, method string, param g.MapStrAny) (err error) {
	if method == "Delete" {
		userIds := garray.NewIntArrayFrom(gconv.Ints(param["ids"]))
		userIds.RemoveValue(1)
		// 删除用户时删除相关数据
		_, err = dao.BaseSysUserRole.Ctx(ctx).WhereIn("userId", userIds.Slice()).Delete()
		if err != nil {
			return err
		}

	}
	return
}

// ServiceAdd 方法 添加用户
func (s *sBaseSysUserService) ServiceAdd(ctx context.Context, req *dzhcore.AddReq) (data interface{}, err error) {
	var (
		m      = s.Dao.Ctx(ctx)
		r      = g.RequestFromCtx(ctx)
		reqmap = r.GetMap()
	)

	// 如果reqmap["password"]不为空，则对密码进行md5加密
	if !r.Get("password").IsNil() {
		reqmap["password"] = gmd5.MustEncryptString(r.Get("password").String())
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		reqmap["id"] = dzhcore.NodeSnowflake.Generate().String()
		_, err = m.TX(tx).Data(reqmap).Insert()
		if err != nil {
			return err
		}

		// 如果请求参数中不包含roleIdList说明不修改角色信息
		if !r.Get("roleIdList").IsNil() {
			roleModel := dao.BaseSysUserRole.Ctx(ctx).TX(tx)
			roleArray := garray.NewArray()
			inRoleIdSet := gset.NewFrom(r.Get("roleIdList").Ints())
			inRoleIdSet.Iterator(func(v interface{}) bool {
				roleArray.PushRight(g.Map{
					"id":     dzhcore.NodeSnowflake.Generate().String(),
					"userId": gconv.Uint(reqmap["id"]),
					"roleId": gconv.Uint(v),
				})
				return true
			})

			_, err = roleModel.Fields("id,userId,roleId").Insert(roleArray)
			if err != nil {
				return err
			}

		}

		data = g.Map{"id": reqmap["id"]}
		return
	})

	return
}

// ServiceInfo 方法 返回服务信息
func (s *sBaseSysUserService) ServiceInfo(ctx g.Ctx, req *dzhcore.InfoReq) (data interface{}, err error) {
	result, err := s.Service.ServiceInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	if result.(gdb.Record).IsEmpty() {
		return nil, nil
	}

	resultMap := result.(gdb.Record).Map()

	// 获取角色
	userRoles := dao.BaseSysUserRole.Ctx(ctx)
	roleIds, err := userRoles.Where("userId = ?", resultMap["id"]).Fields("roleId").Array()
	if err != nil {
		return nil, err
	}

	resultMap["roleIdList"] = roleIds

	type UserInfo struct {
		Id         string   `json:"id"           orm:"id"           ` //
		Name       string   `json:"name"         orm:"name"         ` //
		Username   string   `json:"username"     orm:"username"     ` //
		NickName   string   `json:"nickName"     orm:"nickName"     ` //
		HeadImg    string   `json:"headImg"      orm:"headImg"      ` //
		Phone      string   `json:"phone"        orm:"phone"        ` //
		Email      string   `json:"email"        orm:"email"        ` //
		Status     int      `json:"status"       orm:"status"       ` //
		Remark     string   `json:"remark"       orm:"remark"       ` //
		RoleIdList []string `json:"roleIdList"`
	}

	var userInfo *UserInfo
	err = gconv.Struct(resultMap, &userInfo)
	if err != nil {
		return nil, err
	}

	data = userInfo

	return
}

// ServiceUpdate 方法 更新用户信息
func (s *sBaseSysUserService) ServiceUpdate(ctx context.Context, req *dzhcore.UpdateReq) (data interface{}, err error) {
	var (
		admin = common.GetAdmin(ctx)
		m     = s.Dao.Ctx(ctx)
	)

	r := g.RequestFromCtx(ctx)
	rMap := r.GetMap()

	// 如果不传入ID代表更新当前用户
	userId := r.Get("id", admin.UserId).String()
	userInfo, err := m.Where("id = ?", userId).One()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	if userInfo.IsEmpty() {
		err = gerror.New("用户不存在")
		return
	}

	roleIds, err := dao.BaseSysUserRole.Ctx(ctx).Where("userId = ?", userId).Fields("roleId").Array()
	if err != nil {
		return
	}

	// 禁止禁用超级管理员
	userRoles := garray.NewStrArrayFrom(gconv.SliceStr(roleIds))
	if userRoles.Contains("1") && userId == "1" && !r.Get("status").IsNil() && r.Get("status").Int() == 0 {
		err = gerror.New("禁止禁用超级管理员")
		return
	}

	// 如果请求的password不为空并且密码加密后的值有变动，说明要修改密码
	var rPassword = r.Get("password", "").String()
	if rPassword != "" && rPassword != userInfo["password"].String() {
		rMap["password"], _ = gmd5.Encrypt(rPassword)
		rMap["passwordV"] = userInfo["passwordV"].Int() + 1
	} else {
		delete(rMap, "password")
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		roleModel := dao.BaseSysUserRole.Ctx(ctx).Where("userId = ?", userId)
		// 如果请求参数中不包含roleIdList说明不修改角色信息
		if !r.Get("roleIdList").IsNil() {
			inRoleIdSet := gset.NewFrom(r.Get("roleIdList").Strings())
			roleIdsSet := gset.NewFrom(gconv.Strings(roleIds))
			// 如果请求的角色信息未发生变化则跳过更新逻辑
			if roleIdsSet.Diff(inRoleIdSet).Size() != 0 || inRoleIdSet.Diff(roleIdsSet).Size() != 0 {
				roleArray := garray.NewArray()
				inRoleIdSet.Iterator(func(v interface{}) bool {
					roleArray.PushRight(g.Map{
						"id":     dzhcore.NodeSnowflake.Generate().String(),
						"userId": gconv.String(userId),
						"roleId": gconv.String(v),
					})
					return true
				})

				//删除角色数据
				_, err = roleModel.Delete()
				if err != nil {
					return err
				}
				//重新写入新的角色关系
				_, err = roleModel.Fields("id,userId,roleId").Insert(roleArray)
				if err != nil {
					return err
				}
			}
		}

		_, err = dao.BaseSysUser.Ctx(ctx).Where("id", userId).Update(rMap)

		if err != nil {
			return err
		}
		return
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	//禁止用户后 或者 修改了密码  删除用户登录缓存数据，用户需要重新登录
	if !r.Get("status").IsNil() && r.Get("status").Int() == 0 || rPassword != "" && rPassword != userInfo["password"].String() {
		err := service.BaseSysUserService().DeleteCache(ctx, userId)
		if err != nil {
			g.Log().Error(ctx, err)
			return nil, err
		}
	}

	return
}

// 删除用户缓存
func (s *sBaseSysUserService) DeleteCache(ctx context.Context, userId string) (err error) {
	_, err = dzhcore.CacheManager.Remove(ctx, fmt.Sprintf("admin:token:%v", userId))
	if err != nil {
		return err
	}
	_, err = dzhcore.CacheManager.Remove(ctx, fmt.Sprintf("admin:token:refresh:%v", userId))
	if err != nil {
		return err
	}
	_, err = dzhcore.CacheManager.Remove(ctx, fmt.Sprintf("admin:department:%v", userId))
	if err != nil {
		return err
	}
	_, err = dzhcore.CacheManager.Remove(ctx, fmt.Sprintf("admin:passwordVersion:%v", userId))
	if err != nil {
		return err
	}
	_, err = dzhcore.CacheManager.Remove(ctx, fmt.Sprintf("admin:perms:%v", userId))
	if err != nil {
		return err
	}

	return nil
}

// Move 移动用户部门
func (s *sBaseSysUserService) Move(ctx g.Ctx) (err error) {
	request := g.RequestFromCtx(ctx)
	departmentId := request.Get("departmentId").String()
	userIds := request.Get("userIds").Slice()

	_, err = s.Dao.Ctx(ctx).Where("`id` IN(?)", userIds).Data(g.Map{"departmentId": departmentId}).Update()

	return
}
