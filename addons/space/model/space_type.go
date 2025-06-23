package model

import "github.com/gzdzh-cn/dzhcore"

const TableNameSpaceType = "addons_space_type"

// SpaceType mapped from table <addons_space_type>
type SpaceType struct {
	*dzhcore.Model
	Name     string `gorm:"column:name;type:varchar(255);not null;comment:类别名称 " json:"name"` // 类别名称
	ParentID *int32 `gorm:"column:parentId;comment:父分类ID" json:"parentId"`                    // 父分类ID
}

// TableName SpaceType's table name
func (*SpaceType) TableName() string {
	return TableNameSpaceType
}

// GroupName SpaceType's table group
func (*SpaceType) GroupName() string {
	return "default"
}

// NewSpaceType create a new SpaceType
func NewSpaceType() *SpaceType {
	return &SpaceType{
		Model: dzhcore.NewModel(),
	}
}

// init 创建表
func init() {
	dzhcore.AddModel(&SpaceType{})
}
