package admin

import (
	"dzhgo/addons/task/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gzdzh-cn/dzhcore"
)

type TaskInfoController struct {
	*dzhcore.Controller
}

func init() {
	var taskInfoController = &TaskInfoController{
		&dzhcore.Controller{
			Prefix:  "/admin/task/info",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page", "Start", "Stop"},
			Service: service.NewTaskInfoService(),
		},
	}
	// 注册路由
	dzhcore.AddController(taskInfoController)
}

// TaskInfoStopReq 请求参数
type TaskInfoStopReq struct {
	g.Meta `path:"/stop" method:"GET"`
	ID     int64 `json:"id" v:"required#请输入id"`
}

// Stop 停止任务
func (c *TaskInfoController) Stop(ctx g.Ctx, req *TaskInfoStopReq) (res *dzhcore.BaseRes, err error) {

	err = dzhcore.ClusterRunFunc(ctx, "TaskStopFunc("+gconv.String(req.ID)+")")
	if err != nil {
		return dzhcore.Fail(err.Error()), err
	}
	res = dzhcore.Ok("停止成功")
	return
}

// TaskInfoStartReq 请求参数
type TaskInfoStartReq struct {
	g.Meta `path:"/start" method:"GET"`
	ID     int64 `json:"id" v:"required#请输入id"`
}

// Start 启动任务
func (c *TaskInfoController) Start(ctx g.Ctx, req *TaskInfoStartReq) (res *dzhcore.BaseRes, err error) {

	err = dzhcore.ClusterRunFunc(ctx, "TaskStartFunc("+gconv.String(req.ID)+")")
	if err != nil {
		return dzhcore.Fail(err.Error()), err
	}
	res = dzhcore.Ok("启动成功")
	return
}

// TaskInfoOnceReq 请求参数
type TaskInfoOnceReq struct {
	g.Meta `path:"/once" method:"POST"`
	ID     int64 `json:"id" v:"required#请输入id"`
}

// Once 执行一次
func (c *TaskInfoController) Once(ctx g.Ctx, req *TaskInfoOnceReq) (res *dzhcore.BaseRes, err error) {
	err = c.Service.(*service.TaskInfoService).Once(ctx, req.ID)
	if err != nil {
		return dzhcore.Fail(err.Error()), err
	}
	res = dzhcore.Ok("执行成功")
	return
}

// TaskInfoLogReq 请求参数
type TaskInfoLogReq struct {
	g.Meta `path:"/log" method:"GET"`
	ID     int64 `json:"id"`
	Status int   `json:"status"`
}

// Log 任务日志
func (c *TaskInfoController) Log(ctx g.Ctx, req *TaskInfoLogReq) (res *dzhcore.BaseRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	param := r.GetQueryMapStrStr()
	data, err := c.Service.(*service.TaskInfoService).Log(ctx, param)
	if err != nil {
		return dzhcore.Fail(err.Error()), err
	}
	res = dzhcore.Ok(data)
	return
}
