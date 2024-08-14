package test

import (
	"dzhgo/addons/customer_pro/model/entity"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/util/gconv"

	_ "dzhgo/internal/model"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"context"
	customerDao "dzhgo/addons/customer_pro/dao"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gzdzh-cn/dzhcore/utility/util/logger"
	"testing"
)

var (
	ctx context.Context
)

func Test(t *testing.T) {

	testReg()
}

func testReg() {

	path := "/app/customer_pro/config/receive53Post"
	path2 := "/app/customer_pro/config/page"
	ignorePathSlice := g.Cfg().MustGet(ctx, "modules.base.middleware.authority.ignorePath").String()
	ignoreReg := g.Cfg().MustGet(ctx, "modules.base.middleware.authority.ignoreReg").String()
	logger.Infof(ctx, "ignorePathSlice:%v", ignorePathSlice)
	logger.Infof(ctx, "ignoreReg:%v", ignoreReg)

	isMatch := gregex.IsMatch(ignoreReg, []byte(path))
	isMatch2 := gregex.IsMatch(ignoreReg, []byte(path2))
	logger.Infof(ctx, "isMatch:%v", isMatch)
	logger.Infof(ctx, "isMatch2:%v", isMatch2)
}

func daoSql() {

	var template *entity.AddonsCustomerProWxTemplate
	err := customerDao.AddonsCustomerProWxTemplate.Ctx(ctx).Where("type", 1).Scan(&template)
	if err != nil {
		logger.Error(ctx, err.Error())
	}

	type Filed struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}

	var filedSlice []*Filed
	err = gconv.Struct(template.Fields, &filedSlice)

	g.Dump(filedSlice)

}
