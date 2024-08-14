package sys

import (
	"context"
	v1 "dzhgo/internal/api/admin_v1"
	"dzhgo/internal/dao"
	"dzhgo/internal/service"
	"dzhgo/internal/types"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gzdzh-cn/dzhcore"
	"github.com/gzdzh-cn/dzhcore/utility/util/logger"
)

func init() {
	service.RegisterBaseOpenService(NewsBaseOpenService())
}

type sBaseOpenService struct {
	*dzhcore.Service
}

func NewsBaseOpenService() *sBaseOpenService {
	return &sBaseOpenService{
		&dzhcore.Service{},
	}
}

// AdminEPS 获取eps
func (s *sBaseOpenService) AdminEPS(ctx g.Ctx) (result *g.Var, err error) {
	c := dzhcore.CacheEPS
	result, err = c.GetOrSetFunc(ctx, "adminEPS", func(ctx g.Ctx) (interface{}, error) {
		return s.creatAdminEPS(ctx)
	}, 0)

	return
}

// creatAdminEPS 创建eps
func (s *sBaseOpenService) creatAdminEPS(ctx g.Ctx) (adminEPS interface{}, err error) {

	type Api struct {
		Id      string `json:"id"`
		Module  string `json:"module"`  // 所属模块名称 例如：base
		Method  string `json:"method"`  // 请求方法 例如：GET
		Path    string `json:"path"`    // 请求路径 例如：/welcome
		Prefix  string `json:"prefix"`  // 路由前缀 例如：/admin/base/open
		Summary string `json:"summary"` // 描述 例如：欢迎页面
		Tag     string `json:"tag"`     // 标签 例如：base  好像暂时不用
		Dts     string `json:"dts"`     // 未知 例如：{} 好像暂时不用
	}
	// type Column struct {
	// }
	type Module struct {
		Api     []*Api                `json:"api"`
		Columns []*dzhcore.ColumnInfo `json:"columns"`
		Module  string                `json:"module"`
		Prefix  string                `json:"prefix"`
	}

	admineps := make(map[string][]*Module)
	// 获取所有路由并更新到数据库表 base_eps_admin
	_, err = dao.BaseEpsAdmin.Ctx(ctx).Where("1=1").Delete()
	if err != nil {
		return nil, err
	}

	routers := g.Server().GetRoutes()
	for _, router := range routers {
		if router.Type == ghttp.HandlerTypeMiddleware || router.Type == ghttp.HandlerTypeHook {
			continue
		}
		if router.Method == "ALL" {
			continue
		}
		routeSplite := gstr.Split(router.Route, "/")
		if len(routeSplite) < 5 {
			continue
		}
		if routeSplite[1] != "admin" {
			continue
		}
		module := routeSplite[2]
		method := router.Method
		// 获取最后一个元素加前缀 / 为 path
		path := "/" + routeSplite[len(routeSplite)-1]
		// 获取前面的元素为prefix
		prefix := gstr.Join(routeSplite[0:len(routeSplite)-1], "/")
		// 获取最后一个元素为summary
		summary := routeSplite[len(routeSplite)-1]
		_, err = dao.BaseEpsAdmin.Ctx(ctx).Insert(&Api{
			Id:      dzhcore.NodeSnowflake.Generate().String(),
			Module:  module,
			Method:  method,
			Path:    path,
			Prefix:  prefix,
			Summary: summary,
			Tag:     "",
			Dts:     "",
		})
		if err != nil {
			return nil, err
		}
	}
	// 读取数据库表生成eps
	// var modules []*Module
	items, _ := dao.BaseEpsAdmin.Ctx(ctx).Fields("DISTINCT module,prefix").All()
	for _, item := range items {
		module := item["module"].String()
		prefix := item["prefix"].String()
		apis, _ := dao.BaseEpsAdmin.Ctx(ctx).Where("module=? AND prefix=?", module, prefix).All()
		var apiList []*Api
		for _, api := range apis {
			apiList = append(apiList, &Api{
				Module:  api["module"].String(),
				Method:  api["method"].String(),
				Path:    api["path"].String(),
				Prefix:  api["prefix"].String(),
				Summary: api["summary"].String(),
				Tag:     api["tag"].String(),
				Dts:     api["dts"].String(),
			})
		}
		admineps[module] = append(admineps[module], &Module{
			Api:     apiList,
			Columns: dzhcore.ModelInfo[prefix],
			Module:  module,
			Prefix:  prefix,
		})

	}

	adminEPS = gjson.New(admineps)
	return

}

// AdminEPS 获取eps
func (s *sBaseOpenService) AppEPS(ctx g.Ctx) (result *g.Var, err error) {
	c := dzhcore.CacheEPS
	result, err = c.GetOrSetFunc(ctx, "appEPS", func(ctx g.Ctx) (interface{}, error) {
		return s.creatAppEPS(ctx)
	}, 0)

	return
}

// creatAppEPS 创建app eps
func (s *sBaseOpenService) creatAppEPS(ctx g.Ctx) (appEPS interface{}, err error) {

	type Api struct {
		Id      string `json:"id"`
		Module  string `json:"module"`  // 所属模块名称 例如：base
		Method  string `json:"method"`  // 请求方法 例如：GET
		Path    string `json:"path"`    // 请求路径 例如：/welcome
		Prefix  string `json:"prefix"`  // 路由前缀 例如：/admin/base/open
		Summary string `json:"summary"` // 描述 例如：欢迎页面
		Tag     string `json:"tag"`     // 标签 例如：base  好像暂时不用
		Dts     string `json:"dts"`     // 未知 例如：{} 好像暂时不用
	}
	// type Column struct {
	// }
	type Module struct {
		Api     []*Api                `json:"api"`
		Columns []*dzhcore.ColumnInfo `json:"columns"`
		Module  string                `json:"module"`
		Prefix  string                `json:"prefix"`
	}
	appeps := make(map[string][]*Module)
	// 获取所有路由并更新到数据库表 base_eps_admin
	dao.BaseEpsApp.Ctx(ctx).Where("1=1").Delete()

	//node := utility.CreateSnowflake(ctx)
	routers := g.Server().GetRoutes()
	for _, router := range routers {
		if router.Type == ghttp.HandlerTypeMiddleware || router.Type == ghttp.HandlerTypeHook {
			continue
		}
		if router.Method == "ALL" {
			continue
		}
		routeSplite := gstr.Split(router.Route, "/")
		if len(routeSplite) < 5 {
			continue
		}
		if routeSplite[1] != "app" {
			continue
		}
		module := routeSplite[2]
		method := router.Method
		// 获取最后一个元素加前缀 / 为 path
		path := "/" + routeSplite[len(routeSplite)-1]
		// 获取前面的元素为prefix
		prefix := gstr.Join(routeSplite[0:len(routeSplite)-1], "/")
		// 获取最后一个元素为summary
		summary := routeSplite[len(routeSplite)-1]
		_, err = dao.BaseEpsApp.Ctx(ctx).Insert(&Api{
			Id:      dzhcore.NodeSnowflake.Generate().String(),
			Module:  module,
			Method:  method,
			Path:    path,
			Prefix:  prefix,
			Summary: summary,
			Tag:     "",
			Dts:     "",
		})
		if err != nil {
			return nil, err
		}
	}
	// 读取数据库表生成eps
	// var modules []*Module
	items, _ := dao.BaseEpsApp.Ctx(ctx).Fields("DISTINCT module,prefix").All()
	for _, item := range items {
		module := item["module"].String()
		prefix := item["prefix"].String()
		apis, _ := dao.BaseEpsApp.Ctx(ctx).Where("module=? AND prefix=?", module, prefix).All()
		var apiList []*Api
		for _, api := range apis {
			apiList = append(apiList, &Api{
				Module:  api["module"].String(),
				Method:  api["method"].String(),
				Path:    api["path"].String(),
				Prefix:  api["prefix"].String(),
				Summary: api["summary"].String(),
				Tag:     api["tag"].String(),
				Dts:     api["dts"].String(),
			})
		}
		appeps[module] = append(appeps[module], &Module{
			Api:     apiList,
			Columns: dzhcore.ModelInfo[prefix],
			Module:  module,
			Prefix:  prefix,
		})

	}

	appEPS = gjson.New(appeps)
	return
}

// 获取全部版本
func (s *sBaseOpenService) Versions(ctx context.Context, req *v1.VersionsReq) (data interface{}, err error) {

	versions := dzhcore.GetVersions(req.Addons)
	data = versions

	return
}

// 站点配置
func (s *sBaseOpenService) GetSetting(ctx context.Context, req *v1.GetSettingReq) (data interface{}, err error) {

	var setting *types.CommonSetting
	err = dao.BaseSysSetting.Ctx(ctx).Where("id", 1).Fields("siteName,logo").Scan(&setting)
	if err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}
	data = setting
	return
}
