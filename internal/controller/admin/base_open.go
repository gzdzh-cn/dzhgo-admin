package admin

import (
	"context"
	logic "dzhgo/internal/logic/sys"
	"dzhgo/internal/service"
	"fmt"
	"net/http"
	"time"

	v1 "dzhgo/internal/api/admin_v1"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gzdzh-cn/dzhcore"
	"github.com/gzdzh-cn/dzhcore/coreconfig"
)

type BaseOpenController struct {
	*dzhcore.ControllerSimple
	baseSysLoginService service.IBaseSysLoginService
	baseOpenService     service.IBaseOpenService
}

func init() {
	var open = &BaseOpenController{
		ControllerSimple:    &dzhcore.ControllerSimple{Prefix: "/admin/base/open"},
		baseSysLoginService: logic.NewsBaseSysLoginService(),
		baseOpenService:     logic.NewsBaseOpenService(),
	}
	// 注册路由
	dzhcore.AddControllerSimple(open)

}

// 验证码接口
func (c *BaseOpenController) BaseOpenCaptcha(ctx context.Context, req *v1.BaseOpenCaptchaReq) (res *dzhcore.BaseRes, err error) {
	data, err := c.baseSysLoginService.Captcha(req)
	res = dzhcore.Ok(data)
	return
}

// eps 接口
func (c *BaseOpenController) Eps(ctx context.Context, req *v1.BaseOpenEpsReq) (res *dzhcore.BaseRes, err error) {
	if !coreconfig.Config.Core.Eps {
		g.Log().Error(ctx, "eps is not open")
		res = dzhcore.Ok(nil)
		return
	}
	data, err := c.baseOpenService.AdminEPS(ctx)
	if err != nil {
		g.Log().Error(ctx, "eps error", err)
		return dzhcore.Fail(err.Error()), err
	}
	res = dzhcore.Ok(data)
	return
}

// login 接口
func (c *BaseOpenController) Login(ctx context.Context, req *v1.BaseOpenLoginReq) (res *dzhcore.BaseRes, err error) {
	data, err := c.baseSysLoginService.Login(ctx, req)
	if err != nil {
		return
	}
	res = dzhcore.Ok(data)
	return
}

// 站点配置
func (c *BaseOpenController) GetSetting(ctx context.Context, req *v1.GetSettingReq) (res *dzhcore.BaseRes, err error) {
	data, err := service.BaseOpenService().GetSetting(ctx, req)
	if err != nil {
		return
	}
	res = dzhcore.Ok(data)
	return
}

// RefreshToken 刷新token
func (c *BaseOpenController) RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *dzhcore.BaseRes, err error) {
	data, err := c.baseSysLoginService.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return
	}
	res = dzhcore.Ok(data)
	return
}

// 版本
func (c *BaseOpenController) Versions(ctx context.Context, req *v1.VersionsReq) (res *dzhcore.BaseRes, err error) {
	data, err := service.BaseOpenService().Versions(ctx, req)
	if err != nil {
		return
	}
	res = dzhcore.Ok(data)
	return
}

// 服务器信息
func (c *BaseOpenController) ServerInfo(ctx context.Context, req *v1.ServerInfoReq) (res *dzhcore.BaseRes, err error) {
	data, err := service.BaseOpenService().ServerInfo(ctx)
	if err != nil {
		return
	}
	res = dzhcore.Ok(data)
	return
}

// sse 流式推送
func (c *BaseOpenController) NoticeSse(ctx context.Context, req *v1.NoticeSseReq) (res *dzhcore.BaseRes, err error) {

	r := g.RequestFromCtx(ctx)
	//  流式回应
	rw := r.Response.RawWriter()
	flusher, ok := rw.(http.Flusher)
	if !ok {
		g.Log().Error(ctx, "NoticeSse error")
		r.Response.WriteStatusExit(500)
		return
	}
	r.Response.Header().Set("Content-Type", "text/event-stream")
	r.Response.Header().Set("Cache-Control", "no-cache")
	r.Response.Header().Set("Connection", "keep-alive")
	r.Response.Header().Set("X-Accel-Buffering", "no") // 禁用nginx缓冲
	r.Response.Header().Set("Access-Control-Allow-Origin", "*")
	r.Response.Header().Set("Access-Control-Allow-Headers", "Cache-Control")

	// 设置响应写入器超时，确保能及时检测到断开
	// r.Response.SetTimeout(5 * time.Second) // GoFrame不支持此方法

	// 添加连接ID用于日志追踪
	connectionID := fmt.Sprintf("sse_%d", time.Now().UnixNano())
	g.Log().Debug(ctx, "SSE连接建立", "connectionID", connectionID)

	// 将连接添加到管理器
	logic.NoticeConnectionManager.AddConnection(connectionID)

	// 立即发送一次包含connectionID的数据
	initialData := fmt.Sprintf(`{"connectionID":"%s","status":"connected"}`, connectionID)
	_, err = fmt.Fprintf(rw, "data: %s\n\n", initialData)
	if err != nil {
		g.Log().Errorf(ctx, "发送初始数据失败: %v, connectionID: %s, error: %v", err, connectionID, err)
		return
	}
	flusher.Flush()
	g.Log().Debugf(ctx, "连接建立，发送初始数据: %s,当前时间: %s", initialData, time.Now().Format(time.DateTime))

	// 使用两个定时器：一个用于发送数据，一个用于检查连接状态
	dataTicker := time.NewTicker(logic.NoticeSend)     // 发送时间
	checkTicker := time.NewTicker(logic.NoticeCleanup) // 清理时间

	// 确保函数结束时记录连接断开并从管理器移除
	defer func() {
		logic.NoticeConnectionManager.RemoveConnection(connectionID)
		dataTicker.Stop()
		checkTicker.Stop()
		g.Log().Debug(ctx, "SSE连接已断开", "connectionID", connectionID)
	}()

	for {
		select {
		case <-ctx.Done():
			// 上下文被取消，退出循环
			g.Log().Debug(ctx, "SSE 连接已断开，上下文被取消", "connectionID", connectionID)
			return
		case <-checkTicker.C:
			// 检查连接是否仍然有效
			if rw == nil || r.Response.Writer == nil {
				g.Log().Debug(ctx, "检测到连接已断开，退出循环", "connectionID", connectionID)
				return
			}

			// 检查请求是否仍然有效
			if r.Request != nil && r.Request.Context().Err() != nil {
				g.Log().Debug(ctx, "请求上下文已取消，退出循环", "connectionID", connectionID)
				return
			}

			// 直接尝试发送心跳数据来检测连接状态
			heartbeatData := fmt.Sprintf("data: {\"type\":\"heartbeat\",\"connectionID\":\"%s\"}\n\n", connectionID)
			_, writeErr := fmt.Fprintf(rw, heartbeatData)
			if writeErr != nil {
				g.Log().Debug(ctx, "心跳写入失败，连接已断开", "connectionID", connectionID, "error", writeErr)
				return
			}

			// 强制刷新，确保数据发送到网络
			flusher.Flush()

			// 检查连接是否还在管理器中（是否被清理）
			if !logic.NoticeConnectionManager.IsConnectionActive(connectionID) {
				g.Log().Debug(ctx, "连接已被清理任务删除，停止SSE协程", "connectionID", connectionID)
				return
			}

		case <-dataTicker.C:
			// 发送数据前检查连接状态
			if !logic.NoticeConnectionManager.IsConnectionActive(connectionID) {
				g.Log().Debug(ctx, "发送数据前检测到连接已被清理，停止SSE协程", "connectionID", connectionID)
				return
			}

			// 定时器触发，发送包含connectionID的JSON数据
			jsonData := fmt.Sprintf(`{"connectionID":"%s","status":"connected"}`, connectionID)
			_, err = fmt.Fprintf(rw, "data: %s\n\n", jsonData)
			if err != nil {
				g.Log().Errorf(ctx, "写入 SSE 数据失败: %v, connectionID: %s, error: %v", err, connectionID, err)
				return
			}
			g.Log().Debugf(ctx, "发送数据: %s，当前时间: %s,最后心跳: %s", jsonData, time.Now().Format(time.DateTime), logic.NoticeConnectionManager.Connections[connectionID].LastHeartbeat.Format(time.DateTime))

			// 刷新缓冲区，将数据发送给客户端
			flusher.Flush()
		}
	}

}

// NoticeSseChecked 前端心跳检测接口
func (c *BaseOpenController) NoticeSseChecked(ctx context.Context, req *v1.NoticeSseCheckedReq) (res *dzhcore.BaseRes, err error) {

	// 调试
	// 当前存在的全部 id
	// connectionSlice := logic.NoticeConnectionManager.GetAllConnectionIDsAndLastHeartbeat()
	// for index, connection := range connectionSlice {
	// 	g.Log().Debug(ctx, "连接检查开始")
	// 	g.Log().Debugf(ctx, "序号: %d,id: %s,最后心跳时间: %s", index+1, connection["id"], connection["lastHeartbeat"])
	// 	g.Log().Debug(ctx, "连接检查结束")
	// }

	var status *logic.ConnectionInfo
	status = logic.NoticeConnectionManager.Connections[req.ConnectionID]
	// 更新心跳时间
	if logic.NoticeConnectionManager.UpdateHeartbeat(req.ConnectionID) {
		// 返回当前连接状态
		// g.Log().Debugf(ctx, "sse心跳更新成功: %s,最后心跳时间: %s", status.ConnectionID, status.LastHeartbeat.Format(time.DateTime))

	} else {
		// 连接不存在或已过期
		// g.Log().Debugf(ctx, "连接不存在或已过期: %s", req.ConnectionID)
		// 返回状态断开，返回 status:Disconnect
		return dzhcore.Ok(map[string]string{"status": "Disconnect"}), nil
	}
	return dzhcore.Ok(status), nil
}

// isConnectionAlive 使用多种方法检测连接是否仍然活跃
func isConnectionAlive(rw http.ResponseWriter, r *ghttp.Request, ctx context.Context) bool {
	// 方法1: 检查响应写入器
	if rw == nil {
		return false
	}

	// 方法2: 检查请求上下文
	if r.Request != nil && r.Request.Context().Err() != nil {
		return false
	}

	// 方法3: 尝试写入测试数据
	_, err := fmt.Fprintf(rw, "")
	if err != nil {
		return false
	}

	// 方法4: 检查HTTP连接状态
	if r.Request != nil && r.Request.Body != nil {
		// 尝试读取一个字节来检测连接状态
		buf := make([]byte, 1)
		_, readErr := r.Request.Body.Read(buf)
		if readErr != nil && readErr.Error() != "EOF" {
			return false
		}
	}

	return true
}
