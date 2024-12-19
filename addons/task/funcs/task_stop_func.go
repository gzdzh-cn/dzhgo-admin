package funcs

import (
	"dzhgo/addons/task/model"
	"dzhgo/addons/task/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gzdzh-cn/dzhcore"
)

type TaskStopFunc struct {
}

func (t *TaskStopFunc) Func(ctx g.Ctx, id string) error {
	taskInfo := model.NewTaskInfo()
	_, err := dzhcore.DBM(taskInfo).Where("id = ?", id).Update(g.Map{"status": 0})
	if err != nil {
		return err
	}

	return service.DisableTask(ctx, id)
}

func (t *TaskStopFunc) IsSingleton() bool {
	return false
}

func (t *TaskStopFunc) IsAllWorker() bool {
	return true
}

func init() {
	dzhcore.RegisterFunc("TaskStopFunc", &TaskStopFunc{})
}
