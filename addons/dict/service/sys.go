// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	definetype "dzhgo/addons/bank/defineType"

	"github.com/gogf/gf/v2/frame/g"
)

type (
	IDictInfoService interface {
		ModifyBefore(ctx context.Context, method string, param g.MapStrAny) (err error)
		// ModifyAfter 修改后
		ModifyAfter(ctx context.Context, method string, param map[string]interface{}) (err error)
		// Data 方法, 用于获取数据
		Data(ctx context.Context, types []string) (data interface{}, err error)
		// 列表
		DictList(ctx context.Context, types string) (data []*definetype.DictType, err error)
	}
)

var (
	localDictInfoService IDictInfoService
)

func DictInfoService() IDictInfoService {
	if localDictInfoService == nil {
		panic("implement not found for interface IDictInfoService, forgot register?")
	}
	return localDictInfoService
}

func RegisterDictInfoService(i IDictInfoService) {
	localDictInfoService = i
}
