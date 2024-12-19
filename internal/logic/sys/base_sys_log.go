package sys

import (
	"dzhgo/internal/common"
	"dzhgo/internal/dao"
	"dzhgo/internal/model"
	"dzhgo/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gzdzh-cn/dzhcore"
)

func init() {
	service.RegisterBaseSysLogService(NewsBaseSysLogService())

}

type sBaseSysLogService struct {
	*dzhcore.Service
}

func NewsBaseSysLogService() *sBaseSysLogService {
	return &sBaseSysLogService{
		&dzhcore.Service{
			Dao:   &dao.BaseSysLog,
			Model: model.NewBaseSysLog(),
			PageQueryOp: &dzhcore.QueryOp{
				KeyWordField: []string{"name", "action", "params", "ip"},
				Select:       "base_sys_log.*,user.name ",
				Join: []*dzhcore.JoinOp{
					{
						Model:     model.NewBaseSysUser(),
						Alias:     "user",
						Type:      "LeftJoin",
						Condition: "user.id = base_sys_log.userID",
					},
				},
			},
		},
	}
}

// Record 记录日志
func (s *sBaseSysLogService) Record(ctx g.Ctx) {
	var (
		admin  = common.GetAdmin(ctx)
		r      = g.RequestFromCtx(ctx)
		userId string
	)

	if admin != nil {
		userId = admin.UserId
	}

	baseSysLog := model.NewBaseSysLog()
	baseSysLog.UserID = userId
	baseSysLog.Action = r.Method + ":" + r.URL.Path
	baseSysLog.IP = r.GetClientIp()
	baseSysLog.IPAddr = r.GetClientIp()
	baseSysLog.Params = r.GetBodyString()
	m := s.Dao.Ctx(ctx)

	_, err := m.Insert(g.Map{
		"id":     dzhcore.NodeSnowflake.Generate().String(),
		"userId": baseSysLog.UserID,
		"action": baseSysLog.Action,
		"ip":     baseSysLog.IP,
		"ipAddr": baseSysLog.IPAddr,
		"params": baseSysLog.Params,
	})
	if err != nil {
		return
	}
}

// Clear 清除日志
func (s *sBaseSysLogService) Clear(isAll bool) (err error) {
	m := s.Dao.Ctx(ctx)
	if isAll {
		_, err = m.Delete("1=1")
	} else {
		keepDays := gconv.Int(service.BaseSysConfService().GetValue("logKeep"))
		_, err = m.Delete("createTime < ?", gtime.Now().AddDate(0, 0, -keepDays).String())
	}
	return
}
