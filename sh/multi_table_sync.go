package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/olivere/elastic/v7"
)

// 聚合数据结构
type AggregatedClues struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	ServicesIds string    `json:"services_ids"`
	CreateTime  time.Time `json:"create_time"`
	GuestIpInfo string    `json:"guest_ip_info"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	FromPage    string    `json:"from_page"`
}

// 多表聚合同步器
type MultiTableSync struct {
	mysqlDB   *sql.DB
	esClient  *elastic.Client
	indexName string
}

// 创建同步器
func NewMultiTableSync(mysqlDSN, esURL, indexName string) (*MultiTableSync, error) {
	// 连接MySQL
	mysqlDB, err := sql.Open("mysql", mysqlDSN)
	if err != nil {
		return nil, fmt.Errorf("连接MySQL失败: %w", err)
	}

	// 连接ES
	esClient, err := elastic.NewClient(
		elastic.SetURL(esURL),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
	)
	if err != nil {
		return nil, fmt.Errorf("连接ES失败: %w", err)
	}

	return &MultiTableSync{
		mysqlDB:   mysqlDB,
		esClient:  esClient,
		indexName: indexName,
	}, nil
}

// 聚合查询数据
func (s *MultiTableSync) aggregateData(ctx context.Context, batchSize int) ([]*AggregatedClues, error) {
	query := `
		SELECT 
			c.id,
			c.name,
			c.services_ids,
			c.create_time,
			ca.guest_ip_info,
			ca.address,
			ca.city,
			ca.from_page
		FROM addons_customer_pro_clues c
		LEFT JOIN addons_customer_pro_clues_attr ca ON c.id = ca.clues_id
		WHERE c.deleted_at IS NULL
		ORDER BY c.create_time DESC
		LIMIT ?
	`

	rows, err := s.mysqlDB.QueryContext(ctx, query, batchSize)
	if err != nil {
		return nil, fmt.Errorf("查询数据失败: %w", err)
	}
	defer rows.Close()

	var results []*AggregatedClues
	for rows.Next() {
		var clues AggregatedClues
		err := rows.Scan(
			&clues.ID,
			&clues.Name,
			&clues.ServicesIds,
			&clues.CreateTime,
			&clues.GuestIpInfo,
			&clues.Address,
			&clues.City,
			&clues.FromPage,
		)
		if err != nil {
			return nil, fmt.Errorf("扫描数据失败: %w", err)
		}
		results = append(results, &clues)
	}

	return results, nil
}

// 同步到ES
func (s *MultiTableSync) syncToES(ctx context.Context, data []*AggregatedClues) error {
	if len(data) == 0 {
		return nil
	}

	// 创建批量请求
	bulkRequest := s.esClient.Bulk()

	for _, item := range data {
		// 添加到批量请求
		bulkRequest.Add(elastic.NewBulkIndexRequest().
			Index(s.indexName).
			Id(item.ID).
			Doc(item))
	}

	// 执行批量操作
	bulkResponse, err := bulkRequest.Do(ctx)
	if err != nil {
		return fmt.Errorf("批量同步失败: %w", err)
	}

	// 检查错误
	if bulkResponse.Errors {
		for _, item := range bulkResponse.Items {
			for _, action := range item {
				if action.Error != nil {
					log.Printf("同步失败 - ID: %s, 错误: %v", action.Id, action.Error)
				}
			}
		}
	}

	log.Printf("成功同步 %d 条数据", len(data))
	return nil
}

// 全量同步
func (s *MultiTableSync) FullSync(ctx context.Context) error {
	log.Println("开始全量同步...")

	// 清空ES索引
	_, err := s.esClient.DeleteIndex(s.indexName).Do(ctx)
	if err != nil {
		return fmt.Errorf("删除索引失败: %w", err)
	}

	// 创建索引
	_, err = s.esClient.CreateIndex(s.indexName).BodyString(`{
		"settings": {
			"number_of_replicas": 0,
			"number_of_shards": 1
		},
		"mappings": {
			"properties": {
				"id": {"type": "keyword"},
				"name": {"type": "text"},
				"services_ids": {"type": "keyword"},
				"create_time": {"type": "date"},
				"guest_ip_info": {"type": "text"},
				"address": {"type": "text"},
				"city": {"type": "keyword"},
				"from_page": {"type": "text"}
			}
		}
	}`).Do(ctx)
	if err != nil {
		return fmt.Errorf("创建索引失败: %w", err)
	}

	// 分批同步
	batchSize := 1000
	offset := 0

	for {
		// 查询数据
		data, err := s.aggregateData(ctx, batchSize)
		if err != nil {
			return err
		}

		if len(data) == 0 {
			break
		}

		// 同步到ES
		err = s.syncToES(ctx, data)
		if err != nil {
			return err
		}

		offset += len(data)
		log.Printf("已同步 %d 条数据", offset)

		// 如果数据量小于批次大小，说明已经同步完
		if len(data) < batchSize {
			break
		}
	}

	log.Println("全量同步完成")
	return nil
}

// 增量同步（基于时间戳）
func (s *MultiTableSync) IncrementalSync(ctx context.Context, lastSyncTime time.Time) error {
	log.Printf("开始增量同步，上次同步时间: %v", lastSyncTime)

	query := `
		SELECT 
			c.id,
			c.name,
			c.services_ids,
			c.create_time,
			ca.guest_ip_info,
			ca.address,
			ca.city,
			ca.from_page
		FROM addons_customer_pro_clues c
		LEFT JOIN addons_customer_pro_clues_attr ca ON c.id = ca.clues_id
		WHERE c.deleted_at IS NULL 
		AND c.update_time > ?
		ORDER BY c.update_time ASC
	`

	rows, err := s.mysqlDB.QueryContext(ctx, query, lastSyncTime)
	if err != nil {
		return fmt.Errorf("增量查询失败: %w", err)
	}
	defer rows.Close()

	var data []*AggregatedClues
	for rows.Next() {
		var clues AggregatedClues
		err := rows.Scan(
			&clues.ID,
			&clues.Name,
			&clues.ServicesIds,
			&clues.CreateTime,
			&clues.GuestIpInfo,
			&clues.Address,
			&clues.City,
			&clues.FromPage,
		)
		if err != nil {
			return fmt.Errorf("扫描数据失败: %w", err)
		}
		data = append(data, &clues)
	}

	if len(data) > 0 {
		err = s.syncToES(ctx, data)
		if err != nil {
			return err
		}
		log.Printf("增量同步完成，同步 %d 条数据", len(data))
	} else {
		log.Println("没有新数据需要同步")
	}

	return nil
}

// 使用示例
func ExampleUsage() {
	// 配置参数
	mysqlDSN := "root:password@tcp(localhost:3306)/dzhgo?charset=utf8mb4&parseTime=true"
	esURL := "http://localhost:9200"
	indexName := "addons_customer_pro_clues_aggregated"

	// 创建同步器
	sync, err := NewMultiTableSync(mysqlDSN, esURL, indexName)
	if err != nil {
		log.Fatalf("创建同步器失败: %v", err)
	}
	defer sync.mysqlDB.Close()

	ctx := context.Background()

	// 执行全量同步
	err = sync.FullSync(ctx)
	if err != nil {
		log.Fatalf("全量同步失败: %v", err)
	}

	// 定时增量同步
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	lastSyncTime := time.Now()

	for {
		select {
		case <-ticker.C:
			err := sync.IncrementalSync(ctx, lastSyncTime)
			if err != nil {
				log.Printf("增量同步失败: %v", err)
			} else {
				lastSyncTime = time.Now()
			}
		}
	}
}
