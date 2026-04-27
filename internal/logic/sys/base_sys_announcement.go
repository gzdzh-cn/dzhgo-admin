package sys

import (
	"context"
	"dzhgo/internal/common"
	"dzhgo/internal/dao"
	"dzhgo/internal/model"
	"dzhgo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gzdzh-cn/dzhcore"
)

func init() {
	service.RegisterBaseSysAnnouncementService(&sBaseSysAnnouncementService{})
}

type sBaseSysAnnouncementService struct {
	*dzhcore.Service
}

func NewsBaseSysAnnouncementService() *sBaseSysAnnouncementService {
	return &sBaseSysAnnouncementService{
		&dzhcore.Service{
			Dao:   &dao.BaseSysAnnouncement,
			Model: model.NewBaseSysAnnouncement(),
			ListQueryOp: &dzhcore.QueryOp{
				KeyWordField: []string{"title"},
				AddOrderby:   g.MapStrStr{"top": "DESC", "createTime": "DESC"},
				Where: func(ctx context.Context) []g.Array {
					return []g.Array{}
				},
				OrWhere: func(ctx context.Context) []g.Array {
					return []g.Array{}
				},
				Extend: func(ctx g.Ctx, m *gdb.Model) *gdb.Model {
					return m
				},
				ModifyResult: func(ctx g.Ctx, data any) any {
					return data
				},
			},
			PageQueryOp: &dzhcore.QueryOp{
				KeyWordField: []string{"title"},
				AddOrderby:   g.MapStrStr{"top": "DESC", "createTime": "DESC"},
				Where: func(ctx context.Context) []g.Array {
					return []g.Array{}
				},
				OrWhere: func(ctx context.Context) []g.Array {
					return []g.Array{}
				},
				Extend: func(ctx g.Ctx, m *gdb.Model) *gdb.Model {
					return m
				},
				ModifyResult: func(ctx g.Ctx, data any) any {
					return data
				},
			},
		},
	}
}

func (s *sBaseSysAnnouncementService) Test(ctx context.Context) (err error) {
	return nil
}

// ServiceList 重写父方法，返回公告列表 + 未读数量
func (s *sBaseSysAnnouncementService) ServiceList(ctx context.Context, req *dzhcore.ListReq) (data any, err error) {
	admin := common.GetAdmin(ctx)

	type AnnouncementItem struct {
		Id         string `json:"id"`
		Title      string `json:"title"`
		Content    string `json:"content"`
		Type       int    `json:"type"`
		Status     int    `json:"status"`
		Top        int    `json:"top"`
		CreateTime string `json:"createTime"`
		IsRead     int    `json:"isRead"`
	}

	// 查询公告列表
	var list []AnnouncementItem
	err = dao.BaseSysAnnouncement.Ctx(ctx).
		Fields("id, title, content, type, status, top, createTime").
		Where("status", 1).
		Order("top DESC, createTime DESC").
		Scan(&list)
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = []AnnouncementItem{}
	}

	// 统计当前用户的已读公告ID
	readMap := make(map[string]bool)
	if len(list) > 0 {
		ids := make([]string, len(list))
		for i, item := range list {
			ids[i] = item.Id
		}
		var readRecords []struct {
			AnnouncementId string `json:"announcementId"`
		}
		err = dao.BaseSysAnnouncementRead.Ctx(ctx).
			Fields("announcement_id").
			Where("user_id", admin.UserId).
			WhereIn("announcement_id", ids).
			Scan(&readRecords)
		if err != nil {
			return nil, err
		}
		for _, r := range readRecords {
			readMap[r.AnnouncementId] = true
		}
	}

	// 标记已读状态
	unreadCount := 0
	for i := range list {
		if readMap[list[i].Id] {
			list[i].IsRead = 1
		} else {
			list[i].IsRead = 0
			unreadCount++
		}
	}

	return g.Map{
		"list":        list,
		"unreadCount": unreadCount,
	}, nil
}

// MarkRead 标记公告为已读
func (s *sBaseSysAnnouncementService) MarkRead(ctx context.Context, announcementId string) (err error) {
	admin := common.GetAdmin(ctx)
	count, _ := dao.BaseSysAnnouncementRead.Ctx(ctx).
		Where("user_id", admin.UserId).
		Where("announcement_id", announcementId).
		Count()
	if count > 0 {
		return nil
	}
	_, err = dao.BaseSysAnnouncementRead.Ctx(ctx).Insert(g.Map{
		"id":              dzhcore.NodeSnowflake.Generate().String(),
		"user_id":         admin.UserId,
		"announcement_id": announcementId,
	})
	return
}

// ModifyAfter 监听Update操作，当前端传read=1时标记已读
func (s *sBaseSysAnnouncementService) ModifyAfter(ctx context.Context, method string, param g.MapStrAny) (err error) {
	if method == "Update" && gconv.Int(param["read"]) == 1 {
		admin := common.GetAdmin(ctx)
		announcementId := gconv.String(param["id"])
		if announcementId == "" {
			return nil
		}
		// 判断是否已读，避免重复插入
		count, _ := dao.BaseSysAnnouncementRead.Ctx(ctx).
			Where("user_id", admin.UserId).
			Where("announcement_id", announcementId).
			Count()
		if count == 0 {
			_, err = dao.BaseSysAnnouncementRead.Ctx(ctx).Insert(g.Map{
				"id":              dzhcore.NodeSnowflake.Generate().String(),
				"user_id":         admin.UserId,
				"announcement_id": announcementId,
			})
		}
	}
	return
}

// MarkAllRead 标记所有公告为已读
func (s *sBaseSysAnnouncementService) MarkAllRead(ctx context.Context) (err error) {
	admin := common.GetAdmin(ctx)

	// 获取所有未读的公告ID
	var allIds []string
	err = dao.BaseSysAnnouncement.Ctx(ctx).
		Fields("id").
		Where("status", 1).
		Scan(&allIds)
	if err != nil {
		return err
	}

	// 获取已读的公告ID
	var readIds []string
	err = dao.BaseSysAnnouncementRead.Ctx(ctx).
		Fields("announcement_id").
		Where("user_id", admin.UserId).
		Scan(&readIds)
	if err != nil {
		return err
	}
	readMap := make(map[string]bool)
	for _, id := range readIds {
		readMap[id] = true
	}

	// 批量插入未读的
	for _, id := range allIds {
		if !readMap[id] {
			_, err = dao.BaseSysAnnouncementRead.Ctx(ctx).Insert(g.Map{
				"id":              dzhcore.NodeSnowflake.Generate().String(),
				"user_id":         admin.UserId,
				"announcement_id": id,
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}
