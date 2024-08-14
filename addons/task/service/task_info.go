package service

import (
	"dzhgo/addons/task/dao"
	"time"

	"dzhgo/addons/task/model"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gzdzh-cn/dzhcore"
	"github.com/robfig/cron"
)

type TaskInfoService struct {
	*dzhcore.Service
}

func NewTaskInfoService() *TaskInfoService {
	return &TaskInfoService{
		&dzhcore.Service{
			Dao:   &dao.AddonsTaskInfo,
			Model: model.NewTaskInfo(),
			PageQueryOp: &dzhcore.QueryOp{
				FieldEQ: []string{"status", "type"},
			},
			UniqueKey: map[string]string{
				"name": "任务名称不能重复",
			},
		},
	}
}

func (s *TaskInfoService) ModifyAfter(ctx g.Ctx, method string, param g.MapStrAny) (err error) {
	g.Log().Info(ctx, "TaskInfoService.ModifyAfter", method, param)
	if method == "Add" {
		if gconv.Int(param["status"]) == 1 {
			id, err := s.Dao.Ctx(ctx).Where("name = ?", param["name"]).Value("id")
			if err != nil {
				return err
			}
			return dzhcore.ClusterRunFunc(ctx, "TaskAddTask("+id.String()+")")
		}
	}
	if method == "Update" {
		id := gconv.String(param["id"])
		if gconv.Int(param["status"]) == 1 {
			return dzhcore.ClusterRunFunc(ctx, "TaskStartFunc("+id+")")
		} else {
			return dzhcore.ClusterRunFunc(ctx, "TaskStopFunc("+id+")")
		}
	}
	if method == "Delete" {
		id := gconv.String(param["id"])
		return dzhcore.ClusterRunFunc(ctx, "TaskStopFunc("+id+")")
	}
	return nil
}

// Record 保存任务记录,成功任务每个任务保留最新20条日志,失败日志不会删除
func (s *TaskInfoService) Record(ctx g.Ctx, id string, status int, detail string) error {

	taskLog := model.NewTaskLog()
	_, err := dzhcore.DBM(taskLog).Data(g.Map{
		"id":     dzhcore.NodeSnowflake.Generate().String(),
		"taskId": id,
		"status": status,
		"detail": detail,
	}).Insert()
	if err != nil {
		return err
	}
	if status == 1 {
		record, err := dzhcore.DBM(taskLog).Where("taskId = ?", id).Where("status", 1).Order("id", "desc").Offset(19).One()
		if err != nil {
			return err
		}
		if record.IsEmpty() {
			return nil
		}
		minId := record["id"].Int()
		_, err = dzhcore.DBM(taskLog).Where("taskId = ?", id).Where("status", 1).Where("id < ?", minId).Delete()
		if err != nil {
			return err
		}
	}
	return err
}

// Once 执行一次任务
func (s *TaskInfoService) Once(ctx g.Ctx, id int64) error {
	record, err := s.Dao.Ctx(ctx).Where("id = ?", id).One()
	if err != nil {
		return err
	}
	if record.IsEmpty() {
		return gerror.New("任务不存在")
	}
	funcString := record["service"].String()
	return dzhcore.ClusterRunFunc(ctx, funcString)
}

// Log 获取任务日志
func (s *TaskInfoService) Log(ctx g.Ctx, param g.MapStrStr) (data interface{}, err error) {
	var (
		Total = 0
		Page  = 1
		Size  = 10
	)
	taskLog := model.NewTaskLog()
	m := dzhcore.DBM(taskLog).LeftJoin("addons_task_info", "addons_task_info.id = addons_task_log.taskId")

	if id, ok := param["id"]; ok {
		m = m.Where("taskId = ?", id)
	}
	if status, ok := param["status"]; ok {
		m = m.Where("status = ?", status)
	}
	Total, err = m.Clone().Count()
	m = m.Fields("addons_task_log.*,addons_task_info.name as taskName")
	if err != nil {
		return nil, err
	}
	if page, ok := param["page"]; ok {
		Page = gconv.Int(page)

	}
	if size, ok := param["size"]; ok {
		Size = gconv.Int(size)
	}
	m = m.Limit(Size).Offset((Page - 1) * Size)

	result, err := m.Order("id", "desc").All()
	// g.Dump(result)
	if err != nil {
		return nil, err
	}
	if result.IsEmpty() {
		return g.Map{
			"list": g.Slice{},
			"pagination": g.Map{
				"total": Total,
				"size":  Size,
				"page":  Page,
			},
		}, nil
	}
	// g.Log().Info(ctx, "TaskInfoService.Log", result)
	data = g.Map{
		"list": result,
		"pagination": g.Map{
			"total": Total,
			"size":  Size,
			"page":  Page,
		},
	}
	return
}

// SetNextRunTime 更新下次执行时间
func (s *TaskInfoService) SetNextRunTime(ctx g.Ctx, cronId string, cron string) error {
	// 更新下次执行时间
	nextTime, e := getCronNextTime(cron, time.Now())

	if e == nil {
		_, err := s.Dao.Ctx(ctx).Where("id = ?", cronId).Data("nextRunTime", nextTime).Update()
		if err != nil {
			return err
		}
	} else {
		g.Log().Debug(ctx, "获取下次执行时间失败", e)
	}

	return nil
}

// getCronNextTime 获取下一次Cron的执行时间
func getCronNextTime(cronStr string, t time.Time) (nextTime time.Time, err error) {
	p := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
	s, err := p.Parse(cronStr)
	if err != nil {
		return
	}
	nextTime = s.Next(t)
	return
}
