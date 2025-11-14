package sys

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	baseCommon "dzhgo/internal/common"
	"dzhgo/internal/dao"
	"dzhgo/internal/model"
	"dzhgo/internal/model/do"
	"dzhgo/internal/model/entity"
	"dzhgo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gzdzh-cn/dzhcore"
)

func init() {
	service.RegisterBaseSysNoticeService(&sBaseSysNoticeService{})
	// 启动队列消费者，确保测试过程中队列消息能被处理
	startQueue()
	// 启动连接清理任务
	// NoticeConnectionManager.TaskCleanupInactiveConnections()
}

type sBaseSysNoticeService struct {
	*dzhcore.Service
}

// 队列消息结构
type NoticeQueueMessage struct {
	NoticeId  string `json:"noticeId"`
	UserId    string `json:"userId"`
	Action    string `json:"action"`
	Timestamp int64  `json:"timestamp"`
}

// Redis 队列配置
const (
	NoticeQueueKey     = "notice:queue"   // 队列键名
	NoticeQueueTimeout = 30 * time.Second // 队列超时时间
)

var (
	NoticeSend    = time.Duration(g.Cfg().MustGet(gctx.GetInitCtx(), "core.notice.send").Int()) * time.Second    // 发送时间，单位秒
	NoticeCleanup = time.Duration(g.Cfg().MustGet(gctx.GetInitCtx(), "core.notice.cleanup").Int()) * time.Second // 清理时间，单位秒
	NoticeExpire  = time.Duration(g.Cfg().MustGet(gctx.GetInitCtx(), "core.notice.expire").Int()) * time.Second  // 过期时间，单位秒
)

func NewBaseSysNoticeService() *sBaseSysNoticeService {
	return &sBaseSysNoticeService{
		&dzhcore.Service{
			Dao:   &dao.BaseSysNotice,
			Model: model.NewBaseSysNotice(),
			ListQueryOp: &dzhcore.QueryOp{
				FieldEQ:      []string{""},                         // 字段等于
				KeyWordField: []string{""},                         // 模糊搜索匹配的数据库字段
				AddOrderby:   g.MapStrStr{"nr.createTime": "DESC"}, // 添加排序
				Where: func(ctx context.Context) []g.Array { // 自定义条件
					return []g.Array{}
				},
				OrWhere: func(ctx context.Context) []g.Array { // or 自定义条件
					return []g.Array{}
				},
				Select: "",                  // 查询字段,多个字段用逗号隔开 如: id,name  或  a.id,a.name,b.name AS bname
				As:     "",                  //主表别名
				Join:   []*dzhcore.JoinOp{}, // 关联查询
				Extend: func(ctx g.Ctx, m *gdb.Model) *gdb.Model { // 追加其他条件
					return m
				},
				ModifyResult: func(ctx g.Ctx, data interface{}) interface{} { // 修改结果
					return data
				},
			},
			PageQueryOp: &dzhcore.QueryOp{
				FieldEQ:      []string{""},                         // 字段等于
				KeyWordField: []string{"n.title", "n.remark"},      // 模糊搜索匹配的数据库字段
				AddOrderby:   g.MapStrStr{"nr.createTime": "DESC"}, // 添加排序
				Where: func(ctx context.Context) []g.Array { // 自定义条件
					admin := baseCommon.GetAdmin(ctx)
					userId := admin.UserId

					// 条件搜索，noType 类型和关键词 keyword
					rmap := g.RequestFromCtx(ctx).GetMap()
					noType := gconv.String(rmap["noType"])

					whereData := []g.Array{
						{"nr.user_id = ?", userId},
						{"n.noType = ?", noType, noType != ""},
						{"nr.createTime BETWEEN ? AND ?", gconv.SliceAny(rmap["datetimerange"]), rmap["datetimerange"] != nil && rmap["datetimerange"] != ""},
					}

					return whereData
				},
				OrWhere: func(ctx context.Context) []g.Array { // or 自定义条件
					return []g.Array{}
				},
				Select: "nr.id,nr.createTime,n.title,n.remark,n.noType,nr.readTime,nr.status", // 查询字段,多个字段用逗号隔开 如: id,name  或  a.id,a.name,b.name AS bname
				As:     "n",                                                                   //主表别名
				Join: []*dzhcore.JoinOp{
					{
						Model:     model.NewBaseSysNoticeUserRead(),
						Alias:     "nr",
						Type:      "RightJoin",
						Condition: "nr.notice_id = n.id",
					},
				}, // 关联查询
				Extend: func(ctx g.Ctx, m *gdb.Model) *gdb.Model { // 追加其他条件
					return m
				},
				ModifyResult: func(ctx g.Ctx, data any) interface{} { // 修改结果

					type Pagination struct {
						Page        int `json:"page"`
						Size        int `json:"size"`
						Total       int `json:"total"`
						TotalNoRead int `json:"totalNoRead"`
					}
					type List struct {
						*entity.BaseSysNoticeUserRead
						Title  string `json:"title"`
						Remark string `json:"remark"`
						NoType string `json:"noType"`
					}
					type Page struct {
						Pagination Pagination `json:"pagination"`
						List       []*List    `json:"list"`
					}
					var (
						pageData Page
					)
					err := gconv.Struct(data, &pageData)
					if err != nil {
						g.Log().Error(ctx, err.Error())
						return err
					}

					admin := baseCommon.GetAdmin(ctx)
					userId := admin.UserId
					pageData.Pagination.TotalNoRead, err = dao.BaseSysNoticeUserRead.Ctx(ctx).Where("user_id = ? AND status = ?", userId, 0).Count()
					if err != nil {
						g.Log().Error(ctx, err.Error())
						return err
					}

					return pageData
				},
			},
			InsertParam: func(ctx context.Context) g.MapStrAny { // Add时插入参数
				return g.MapStrAny{}
			},
			Before: func(ctx context.Context) (err error) { // CRUD前的操作
				return nil
			},
			InfoIgnoreProperty: "",            // Info时忽略的字段,多个字段用逗号隔开
			UniqueKey:          g.MapStrStr{}, // 唯一键 key:字段名 value:错误信息
			NotNullKey:         g.MapStrStr{}, // 非空键 key:字段名 value:错误信息
		},
	}
}

func startQueue() {
	// 使用全局上下文，确保队列消费者能够持续运行

	// 延迟启动，等待Redis连接初始化完成
	go func() {
		// 等待Redis连接初始化
		for i := 0; i < 30; i++ { // 最多等待30秒
			if dzhcore.Redis != nil {
				// 启动消费队列
				service.BaseSysNoticeService().StartRedisQueueConsumer()
				g.Log().Info(ctx, "Redis队列消费者启动完成")
				return
			}
			g.Log().Debug(ctx, "等待Redis连接初始化...", "attempt", i+1)
			time.Sleep(1 * time.Second)
		}
		g.Log().Error(ctx, "Redis连接初始化超时，队列消费者启动失败")
	}()
}

func (s *sBaseSysNoticeService) ServiceInfo(ctx context.Context, req *dzhcore.InfoReq) (data any, err error) {
	admin := baseCommon.GetAdmin(ctx)
	userId := admin.UserId

	type NoticeInfo struct {
		*entity.BaseSysNoticeUserRead
		Title  string `json:"title"`
		Remark string `json:"remark"`
		NoType string `json:"noType"`
	}

	var noticeInfo *NoticeInfo
	err = dao.BaseSysNotice.Ctx(ctx).As("n").
		RightJoin(dao.BaseSysNoticeUserRead.Table(), "nr", "nr.notice_id = n.id").
		Where("nr.user_id = ? AND nr.id = ?", userId, req.Id).
		Fields("nr.*,n.title,n.remark,n.noType").
		Scan(&noticeInfo)
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return nil, err
	}

	return noticeInfo, nil
}

// 更新阅读状态
func (s *sBaseSysNoticeService) ServiceUpdate(ctx context.Context, req *dzhcore.UpdateReq) (data any, err error) {
	admin := baseCommon.GetAdmin(ctx)
	userId := admin.UserId
	rmap := g.RequestFromCtx(ctx).GetMap()

	noticeId := rmap["id"].(string)
	noticeData := do.BaseSysNoticeUserRead{
		ReadTime: gtime.Now(),
		Status:   1,
	}

	_, err = dao.BaseSysNoticeUserRead.Ctx(ctx).Data(noticeData).Where("id = ? AND user_id = ?", noticeId, userId).Update()
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return nil, err
	}

	return nil, nil
}

// 删除用户消息
func (s *sBaseSysNoticeService) ServiceDelete(ctx context.Context, req *dzhcore.DeleteReq) (data any, err error) {
	admin := baseCommon.GetAdmin(ctx)
	userId := admin.UserId
	rmap := g.RequestFromCtx(ctx).GetMap()

	noticeIds := gconv.Strings(rmap["ids"])

	if len(noticeIds) == 0 {
		g.Log().Warning(ctx, "没有有效的 noticeIds 需要删除")
		return nil, nil
	}

	_, err = dao.BaseSysNoticeUserRead.Ctx(ctx).Where("id IN (?) AND user_id = ?", noticeIds, userId).Delete()
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return nil, err
	}

	// 加入操作日志
	baseCommon.RecordActionLog(ctx, "消息通知", userId, fmt.Sprintf("删除消息，消息ID：%v", noticeIds))

	return nil, nil
}

// 一键已阅
func (s *sBaseSysNoticeService) ServiceReadAll(ctx context.Context) (data any, err error) {
	admin := baseCommon.GetAdmin(ctx)
	userId := admin.UserId

	_, err = dao.BaseSysNoticeUserRead.Ctx(ctx).Where("user_id = ? AND status = ?", userId, 0).Data(g.Map{"status": 1}).Update()
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return nil, err
	}
	// 加入操作日志
	baseCommon.RecordActionLog(ctx, "消息通知", userId, "消息一键已阅")

	return nil, nil
}

// NoticeAdd 给指定用户推送消息（保持接口兼容性）
func (s *sBaseSysNoticeService) NoticeAdd(ctx context.Context, notice *entity.BaseSysNotice, userIdSlice *[]string) (data any, err error) {
	if userIdSlice == nil || len(*userIdSlice) == 0 {
		return s.NoticeAddWithTarget(ctx, notice, baseCommon.ToAllAdmins())
	}
	return s.NoticeAddWithTarget(ctx, notice, baseCommon.ToUsers(*userIdSlice))
}

// NoticeAddWithTarget 使用接口多态的消息推送
func (s *sBaseSysNoticeService) NoticeAddWithTarget(ctx context.Context, notice *entity.BaseSysNotice, target baseCommon.NoticeTarget) (data any, err error) {
	id := dzhcore.CreateSnowflakeId()
	insertData := do.BaseSysNotice{
		Id:     id,
		Title:  notice.Title,
		NoType: notice.NoType,
		Remark: notice.Remark,
	}
	_, err = dao.BaseSysNotice.Ctx(ctx).Data(insertData).Insert()
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return nil, err
	}

	newCtx := gctx.New()
	// 异步处理消息通知，避免阻塞主流程
	go func(ctx context.Context) {
		defer func() {
			if r := recover(); r != nil {
				g.Log().Errorf(ctx, "消息通知处理发生 panic: %v", r)
			}
		}()

		noticeData := &entity.BaseSysNotice{
			Id:     id,
			Title:  notice.Title,
			NoType: notice.NoType,
			Remark: notice.Remark,
		}

		_, err := s.NoticeDoWithTarget(ctx, noticeData, target)
		if err != nil {
			g.Log().Errorf(ctx, "消息通知处理失败: %v", err)
		}
	}(newCtx)

	return nil, nil
}

// NoticeAddToAllUsers 添加通知并推送给全部用户
func (s *sBaseSysNoticeService) NoticeAddToAllUsers(ctx context.Context, notice *entity.BaseSysNotice) (data any, err error) {
	return s.NoticeAddWithTarget(ctx, notice, baseCommon.ToAllUsers())
}

// 消息通知处理（保持接口兼容性）
func (s *sBaseSysNoticeService) NoticeDo(ctx context.Context, notice *entity.BaseSysNotice, userIdSlice *[]string) (data any, err error) {
	if userIdSlice == nil || len(*userIdSlice) == 0 {
		return s.NoticeDoWithTarget(ctx, notice, baseCommon.ToAllAdmins())
	}
	return s.NoticeDoWithTarget(ctx, notice, baseCommon.ToUsers(*userIdSlice))
}

// NoticeDoWithTarget 使用接口多态的消息处理
func (s *sBaseSysNoticeService) NoticeDoWithTarget(ctx context.Context, notice *entity.BaseSysNotice, target baseCommon.NoticeTarget) (data any, err error) {

	if notice.Id == "" {
		g.Log().Error(ctx, "消息ID为空")
		return nil, gerror.New("消息ID为空")
	}

	// 通过接口获取用户ID列表
	ids, err := target.GetUserIDs(ctx)
	if err != nil {
		g.Log().Error(ctx, "获取用户ID失败:", err)
		return nil, err
	}

	g.Log().Infof(ctx, "准备推送消息给 %d 个用户", len(ids))

	for _, id := range ids {
		_, err = s.NoticePushQueue(ctx, notice.Id, &id)
		if err != nil {
			g.Log().Errorf(ctx, "推送消息到队列失败: noticeId=%s, userId=%s, error=%v", notice.Id, id, err)
			continue
		}
	}

	g.Log().Infof(ctx, "成功推送 %d 条消息到队列", len(ids))
	return nil, nil
}

// 推送队列到 Redis
func (s *sBaseSysNoticeService) NoticePushQueue(ctx context.Context, noticeId string, userId *string) (data any, err error) {
	// 构造队列消息
	queueMessage := NoticeQueueMessage{
		NoticeId:  noticeId,
		UserId:    *userId,
		Action:    "notice_read_insert",
		Timestamp: time.Now().Unix(),
	}

	// 序列化消息
	messageBytes, err := json.Marshal(queueMessage)
	if err != nil {
		g.Log().Errorf(ctx, "序列化队列消息失败: %v", err)
		return nil, err
	}

	// 推送到 Redis 队列
	err = s.pushToRedisQueue(ctx, messageBytes)
	if err != nil {
		g.Log().Errorf(ctx, "推送消息到Redis队列失败: noticeId=%s, userId=%s, error=%v", noticeId, *userId, err)
		return nil, err
	}

	g.Log().Infof(ctx, "消息已推送到Redis队列: noticeId=%s, userId=%s", noticeId, *userId)
	return nil, nil
}

// 队列处理,把队列的数据插入到数据库
func (s *sBaseSysNoticeService) NoticeQueueDo(ctx context.Context, noticeId string, userId *string) (data any, err error) {
	// 检查是否已经存在记录，避免重复插入
	count, err := dao.BaseSysNoticeUserRead.Ctx(ctx).
		Where("user_id = ? AND notice_id = ?", *userId, noticeId).
		Count()
	if err != nil {
		g.Log().Errorf(ctx, "检查消息记录失败: %v", err)
		return nil, err
	}

	// 如果已存在，不重复插入
	if count > 0 {
		g.Log().Infof(ctx, "消息记录已存在，跳过插入: noticeId=%s, userId=%s", noticeId, *userId)
		return nil, nil
	}

	insertData := do.BaseSysNoticeUserRead{
		Id:       dzhcore.CreateSnowflakeId(),
		UserId:   userId,
		NoticeId: noticeId,
		Status:   0,
	}
	_, err = dao.BaseSysNoticeUserRead.Ctx(ctx).Data(insertData).Insert()
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return nil, err
	}

	// 通知到公众号
	// service.CustomerProWxTemplateService().TemplateNotify(ctx, clues, kf.UserId)

	// g.Log().Infof(ctx, "成功插入消息阅读记录: noticeId=%s, userId=%s", noticeId, *userId)
	return nil, nil
}

// 将消息推送到 Redis 队列
func (s *sBaseSysNoticeService) pushToRedisQueue(ctx context.Context, messageBytes []byte) error {

	redis := dzhcore.Redis
	// 使用 LPUSH 将消息推送到队列头部
	_, err := redis.LPush(ctx, NoticeQueueKey, messageBytes)
	if err != nil {
		return gerror.Newf("推送到Redis队列失败: %v", err)
	}

	// 设置队列过期时间，避免队列无限增长
	redis.Expire(ctx, NoticeQueueKey, int64(NoticeQueueTimeout.Seconds()))

	return nil
}

// 从 Redis 队列消费消息
func (s *sBaseSysNoticeService) consumeFromRedisQueue(ctx context.Context) error {

	redis := dzhcore.Redis
	if redis == nil {
		g.Log().Error(ctx, "Redis连接未初始化，无法消费队列")
		return gerror.New("Redis连接未初始化")
	}

	// 使用 BRPOP 阻塞式消费队列消息
	result, err := redis.BRPop(ctx, 0, NoticeQueueKey)
	if err != nil {
		g.Log().Error(ctx, "从Redis队列消费消息失败:", err)
		return gerror.Newf("从Redis队列消费消息失败: %v", err)
	}

	if len(result) < 2 {
		g.Log().Error(ctx, "队列消息格式错误:", "result", result)
		return gerror.New("队列消息格式错误")
	}

	// 解析消息
	var queueMessage NoticeQueueMessage
	err = json.Unmarshal([]byte(result[1].String()), &queueMessage)
	if err != nil {
		g.Log().Errorf(ctx, "解析队列消息失败: %v", err)
		return err
	}

	// 处理消息
	_, err = s.NoticeQueueDo(ctx, queueMessage.NoticeId, &queueMessage.UserId)
	if err != nil {
		g.Log().Errorf(ctx, "处理队列消息失败: %v", err)
		return err
	}

	g.Log().Infof(ctx, "成功处理队列消息: noticeId=%s, userId=%s", queueMessage.NoticeId, queueMessage.UserId)
	return nil
}

// 启动 Redis 队列消费者
func (s *sBaseSysNoticeService) StartRedisQueueConsumer() {
	newCtx := gctx.New()
	go func(ctx context.Context) {

		for {
			select {
			case <-ctx.Done():
				g.Log().Info(ctx, "Redis队列消费者已停止")
				return
			default:
				// 消费队列消息
				err := s.consumeFromRedisQueue(ctx)
				if err != nil {
					g.Log().Errorf(ctx, "消费队列消息失败: %v", err)
					// 短暂延迟后继续
					time.Sleep(1 * time.Second)
				}
			}
		}
	}(newCtx)
}

// 获取全部管理员ID
func getAdminUserIds(ctx context.Context) (data []string, err error) {
	ids, err := dao.BaseSysUser.Ctx(ctx).As("a").
		LeftJoin(dao.BaseSysUserRole.Table(), "user_role", "user_role.userId = a.id").
		LeftJoin(dao.BaseSysRole.Table(), "role", "role.id = user_role.roleId").
		Where("role.id = ?", 1). // 假设角色ID为1的是管理员
		Array("a.id")
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return nil, err
	}
	return gconv.Strings(ids), nil
}

// 获取全部用户ID
func getAllUserIds(ctx context.Context) (data []string, err error) {
	ids, err := dao.BaseSysUser.Ctx(ctx).As("a").
		Array("a.id")
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return nil, err
	}
	return gconv.Strings(ids), nil
}

// 全局连接管理器
var (
	NoticeConnectionManager = &ConnectionManager{
		Connections: make(map[string]*ConnectionInfo),
		mutex:       &sync.RWMutex{},
	}
)

// ConnectionInfo 连接信息
type ConnectionInfo struct {
	ConnectionID  string    `json:"connectionID"`
	LastHeartbeat time.Time `json:"lastHeartbeat"`
	CreateTime    time.Time `json:"createTime"`
	IsActive      bool      `json:"isActive"`
}

// ConnectionManager 连接管理器
type ConnectionManager struct {
	Connections map[string]*ConnectionInfo
	mutex       *sync.RWMutex
}

// AddConnection 添加连接
func (cm *ConnectionManager) AddConnection(connectionID string) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	cm.Connections[connectionID] = &ConnectionInfo{
		ConnectionID:  connectionID,
		LastHeartbeat: time.Now(),
		CreateTime:    time.Now(),
		IsActive:      true,
	}
	// g.Log().Debug(context.Background(), "连接已添加到管理器", "connectionID", connectionID, "total_connections", len(cm.Connections))
}

// RemoveConnection 移除连接
func (cm *ConnectionManager) RemoveConnection(connectionID string) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	delete(cm.Connections, connectionID)
	// g.Log().Debug(context.Background(), "连接已从管理器移除", "connectionID", connectionID, "total_connections", len(cm.Connections))
}

// UpdateHeartbeat 更新心跳
func (cm *ConnectionManager) UpdateHeartbeat(connectionID string) bool {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	if conn, exists := cm.Connections[connectionID]; exists {
		conn.LastHeartbeat = time.Now()
		conn.IsActive = true
		return true
	}
	return false
}

// GetConnectionCount 获取连接数量
func (cm *ConnectionManager) GetConnectionCount() int {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	return len(cm.Connections)
}

// 获取全部 id和最后心跳时间
func (cm *ConnectionManager) GetAllConnectionIDsAndLastHeartbeat() []map[string]string {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	// id 和最后心跳时间的的字符串数组，格式是
	// [
	//   {
	// 	id:"11",
	// 	lastHeartbeat:2025-08-16 10:00:00
	//   }
	// ]
	connectionSlice := make([]map[string]string, 0, len(cm.Connections))
	for id, conn := range cm.Connections {
		connectionSlice = append(connectionSlice, map[string]string{
			"id":            id,
			"lastHeartbeat": conn.LastHeartbeat.Format(time.DateTime),
		})
	}
	return connectionSlice
}

// 获取全部 id
func (cm *ConnectionManager) GetAllConnectionIDs() []string {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	ids := make([]string, 0, len(cm.Connections))
	for id, _ := range cm.Connections {
		ids = append(ids, id)
	}
	return ids
}

// IsConnectionActive 检查连接是否还在管理器中
func (cm *ConnectionManager) IsConnectionActive(connectionID string) bool {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	_, exists := cm.Connections[connectionID]
	return exists // 如果连接存在，就认为是活跃的
}

// CleanupInactiveConnections 清理不活跃的连接
func (cm *ConnectionManager) CleanupInactiveConnections() {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	now := time.Now()
	inactiveConnections := []string{}

	for connectionID, conn := range cm.Connections {
		if now.Sub(conn.LastHeartbeat) > NoticeExpire {
			conn.IsActive = false
			inactiveConnections = append(inactiveConnections, connectionID)
		}
	}

	// 立即删除不活跃的连接，触发对应SSE协程停止
	for _, connectionID := range inactiveConnections {
		if conn, exists := cm.Connections[connectionID]; exists {
			inactiveDuration := now.Sub(conn.LastHeartbeat)
			delete(cm.Connections, connectionID) // 立即删除连接，SSE协程将在下次检查时停止
			g.Log().Debugf(context.Background(), "清理不活跃连接: %s, 无活动时长: %v, 将触发SSE协程停止\n", connectionID, inactiveDuration)
		}
	}

	if len(inactiveConnections) > 0 {
		g.Log().Debug(context.Background(), "连接清理完成，SSE协程将在10秒内停止", "cleaned_count", len(inactiveConnections), "remaining_connections", len(cm.Connections))
	}
}

// 定时清理
func (cm *ConnectionManager) TaskCleanupInactiveConnections() {
	ticker := time.NewTicker(time.Second * 10) // 每10秒清理一次
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			cm.CleanupInactiveConnections()
		}
	}
}

// // 启动定时清理任务

// 检查队列状态
func (s *sBaseSysNoticeService) CheckQueueStatus(ctx context.Context) (data any, err error) {
	redis := dzhcore.Redis
	if redis == nil {
		return g.Map{
			"redis_connected": false,
			"error":           "Redis连接未初始化",
		}, nil
	}

	// 检查队列长度
	queueLength, err := redis.LLen(ctx, NoticeQueueKey)
	if err != nil {
		return g.Map{
			"redis_connected": true,
			"queue_length":    -1,
			"error":           err.Error(),
		}, nil
	}

	// 检查队列中的消息
	var messages []string
	if queueLength > 0 {
		// 获取队列中的前几条消息（不删除）
		queueMessages, err := redis.LRange(ctx, NoticeQueueKey, 0, 4)
		if err == nil {
			for _, msg := range queueMessages {
				messages = append(messages, msg.String())
			}
		}
	}

	return g.Map{
		"redis_connected": true,
		"queue_length":    queueLength,
		"queue_messages":  messages,
		"queue_key":       NoticeQueueKey,
	}, nil
}
