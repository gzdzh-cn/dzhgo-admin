package test

import (
	"dzhgo/addons/dzh3164/dao"
	"dzhgo/addons/dzh3164/logic/sys"
	"dzhgo/addons/dzh3164/service"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gzdzh-cn/dzhcore"
)

var ctx = gctx.New()

func init() {
	service.RegisterDzh3164ConvergeService(sys.NewsDzh3164ConvergeService())
	dzhcore.NewInit()

}

func Test(t *testing.T) {
	dao.AddonsDzh3164Category.Ctx(ctx).As("c FORCE INDEX(idx_addons_dzh3164_category_id)").All()
}
