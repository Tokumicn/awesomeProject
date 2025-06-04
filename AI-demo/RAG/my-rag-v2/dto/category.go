package dto

import "gorm.io/gorm"

// Category 分类标签：树型结构通过parent_id关联
type Category struct {
	gorm.Model
	Name          string `gorm:"type:text column:name" json:"name"`                       // 分类名
	Order         uint8  `gorm:"type:int(8) column:order" json:"order"`                   // 同文件夹下的排序
	ParentID      uint8  `gorm:"type:int(11) column:parent_id" json:"parent_id"`          // 父分类ID
	KnowledgeType uint8  `gorm:"type:int(2) column:knowledge_type" json:"knowledge_type"` // 0: 标准FAQ  1: 寒暄 2：任务型
}
