// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IDictInfoService interface {
		// Data 方法, 用于获取数据
		Data(ctx context.Context, types []string) (data interface{}, err error)
		// ModifyAfter 修改后
		ModifyAfter(ctx context.Context, method string, param map[string]interface{}) (err error)
	}
	IDictTypeService interface{}
)

var (
	localDictInfoService IDictInfoService
	localDictTypeService IDictTypeService
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

func DictTypeService() IDictTypeService {
	if localDictTypeService == nil {
		panic("implement not found for interface IDictTypeService, forgot register?")
	}
	return localDictTypeService
}

func RegisterDictTypeService(i IDictTypeService) {
	localDictTypeService = i
}
