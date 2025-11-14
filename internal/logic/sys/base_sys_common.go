package sys

import (
	"github.com/gzdzh-cn/dzhcore"
)

func init() {
	// service.RegisterBaseSysCommonService(&sBaseSysCommonService{})
}

type sBaseSysCommonService struct {
	*dzhcore.Service
}

func NewBaseSysCommonService() *sBaseSysCommonService {
	return &sBaseSysCommonService{}
}
