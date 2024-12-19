package sys

import (
	"context"
	"dzhgo/internal/dao"
	"dzhgo/internal/model"
	"dzhgo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gzdzh-cn/dzhcore"
)

func init() {
	service.RegisterBaseSysParamService(NewsBaseSysParamService())
}

type sBaseSysParamService struct {
	*dzhcore.Service
}

func NewsBaseSysParamService() *sBaseSysParamService {
	return &sBaseSysParamService{
		&dzhcore.Service{
			Dao:   &dao.BaseSysParam,
			Model: model.NewBaseSysParam(),
		},

		// Service: dzhcore.NewService(model.NewBaseSysParam()),
	}
}

// HtmlByKey 根据配置参数key获取网页内容(富文本)
func (s *sBaseSysParamService) HtmlByKey(key string) string {
	var (
		html = "<html><body>@content</body></html>"
	)
	m := s.Dao.Ctx(ctx)
	record, err := m.Where("keyName = ?", key).One()
	if err != nil {
		html = gstr.Replace(html, "@content", err.Error())
		return html
	}
	if record.IsEmpty() {
		html = gstr.Replace(html, "@content", "keyName notfound")
		return html
	}
	html = gstr.Replace(html, "@content", record["data"].String())

	return html
}

// ModifyAfter 修改后
func (s *sBaseSysParamService) ModifyAfter(ctx context.Context, method string, param g.MapStrAny) (err error) {
	var (
		m = s.Dao.Ctx(ctx)
	)
	result, err := m.All()
	if err != nil {
		return
	}
	for _, v := range result {
		key := "param:" + v["keyName"].String()
		value := v["data"].String()
		err = dzhcore.CacheManager.Set(ctx, key, value, 0)
		if err != nil {
			return
		}
	}
	return
}

// DataByKey 根据配置参数key获取数据
func (s *sBaseSysParamService) DataByKey(ctx context.Context, key string) (data string, err error) {
	var (
		m = s.Dao.Ctx(ctx)
	)
	rKey := "param:" + key
	dataCache, err := dzhcore.CacheManager.Get(ctx, rKey)
	if err != nil {
		return
	}
	if !dataCache.IsEmpty() {
		data = dataCache.String()
		return
	}
	record, err := m.Where("keyName = ?", key).One()
	if err != nil {
		return
	}
	if record.IsEmpty() {
		return
	}
	data = record["data"].String()
	err = dzhcore.CacheManager.Set(ctx, rKey, data, 0)
	return
}
