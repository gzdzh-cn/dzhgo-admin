package task

import (
	baseModel "dzhgo/internal/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"dzhgo/addons/task/model"
	"github.com/gzdzh-cn/dzhcore"

	_ "dzhgo/addons/task/controller"
	_ "dzhgo/addons/task/funcs"
	_ "dzhgo/addons/task/middleware"
	_ "dzhgo/addons/task/packed"
)

func NewInit() {
	var (
		taskInfo = model.NewTaskInfo()
		ctx      = gctx.GetInitCtx()
	)

	g.Log().Debug(ctx, "addon task init start ...")
	g.Log().Debugf(ctx, "task version:%v", Version)
	dzhcore.FillInitData(ctx, "task", &model.TaskInfo{})
	dzhcore.FillInitData(ctx, "task", &baseModel.BaseSysMenu{})
	g.Log().Debug(ctx, "addon task init finished ...")

	result, err := dzhcore.DBM(taskInfo).Where("status = ?", 1).All()
	if err != nil {
		panic(err)
	}
	for _, v := range result {
		id := v["id"].String()
		err = dzhcore.RunFunc(ctx, "TaskAddTask("+id+")")
		if err != nil {
			return
		}
	}

	g.Log().Debug(ctx, "module task init finished ...")

}
