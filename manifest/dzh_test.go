package test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"testing"
)

type AddonsCustomerProClues struct {
	Id               int
	SerialId         sql.NullString
	Level            sql.NullInt64
	AccountId        sql.NullInt64
	CreatedName      sql.NullString
	Name             sql.NullString
	Mobile           sql.NullString
	Wechat           sql.NullString
	SourceFrom       sql.NullString
	GuestIpInfo      sql.NullString
	ServicesIds      sql.NullString
	ServicesId       sql.NullInt64
	ProjectId        sql.NullInt64
	LastFollowupTime sql.NullTime
	FollowupType     sql.NullString
	OceanTime        sql.NullTime
	CreateTime       sql.NullTime
	Keywords         sql.NullString
}

// parseLevel 解析 level 列的值
func parseLevel(levelBytes []byte) (sql.NullInt64, error) {
	var level sql.NullInt64
	levelStr := string(levelBytes)
	if levelStr == "" {
		level.Valid = false
	} else {
		levelInt, err := strconv.ParseInt(levelStr, 10, 64)
		if err != nil {
			return level, err
		}
		level.Valid = true
		level.Int64 = levelInt
	}
	return level, nil
}

func Test_sql(t *testing.T) {
	dsn := "root:dzh123456@tcp(127.0.0.1:3306)/dzhgo_go?charset=utf8mb4&parseTime=True"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}
	fmt.Println("Connected!")
	// 定义 SQL 查询语句，使用 * 查询所有列
	query := `
        SELECT *
        FROM
            addons_customer_pro_clues AS clues
        WHERE
            (clues.ocean_time IS NULL)
            AND deleted_at IS NULL
        GROUP BY
            clues.id
        ORDER BY
            clues.createTime DESC
        LIMIT
            20
        OFFSET
            0
    `

	// 执行查询
	rows, err := db.Query(query)
	if err != nil {
		t.Fatalf("Failed to execute query: %v", err)
	}
	defer rows.Close()

}
