package funcs

import (
	"dzhgo/addons/task/model"
	"dzhgo/addons/task/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gzdzh-cn/dzhcore"
)

type TaskStartFunc struct {
}

func (t *TaskStartFunc) Func(ctx g.Ctx, id string) error {
	taskInfo := model.NewTaskInfo()
	_, err := dzhcore.DBM(taskInfo).Where("id = ?", id).Update(g.Map{"status": 1})
	if err != nil {
		return err
	}
	result, err := dzhcore.DBM(taskInfo).Where("id = ?", id).One()
	if err != nil {
		return err
	}
	if result["taskType"].Int() == 1 {
		every := result["every"].Uint() / 1000
		cron := "@every " + gconv.String(every) + "s"
		funcstring := result["service"].String()
		startDate := result["startDate"].String()
		err = service.EnableTask(ctx, id, funcstring, cron, startDate)

	} else {
		cron := result["cron"].String()
		funcstring := result["service"].String()
		startDate := result["startDate"].String()
		err = service.EnableTask(ctx, id, funcstring, cron, startDate)
	}
	return err

}
func (t *TaskStartFunc) IsSingleton() bool {
	return false
}
func (t *TaskStartFunc) IsAllWorker() bool {
	return true
}

func init() {
	dzhcore.RegisterFunc("TaskStartFunc", &TaskStartFunc{})
}
