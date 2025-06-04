package dto

import (
	"gorm.io/gorm"
)

// Examine 审核
type Examine struct {
	gorm.Model
	OldKnowledge string `gorm:"type:text column:old_knowledge" json:"old_knowledge"`         // 变更前快照：分类信息、标准问、默认答案、相似问
	NewKnowledge string `gorm:"type:text column:new_knowledge" json:"new_knowledge"`         // 变更后快照：分类信息、标准问、默认答案、相似问
	Question     string `gorm:"type:text column:question" json:"question"`                   // 标准问
	CategoryPath string `gorm:"type:varchar(256) column:category_path" json:"category_path"` // 分类路径
	ExamineType  uint8  `gorm:"type:tinyint(4) column:examine_type" json:"examine_type"`     // 操作类型  1-更新 2-发布 3-下线
	CreateUserID uint   `gorm:"type:int(11) column:create_user_id" json:"create_user_id"`    // 提交人ID
	CheckUserID  uint   `gorm:"type:int(11) column:check_user_id" json:"check_user_id"`      // 审核人ID
	SubmitStatus uint8  `gorm:"type:tinyint(4) column:submit_status" json:"submit_status"`   // 提交时状态
	Result       uint8  `gorm:"type:tinyint(4) column:result" json:"result"`                 // 审核结果 0-未审核 1-通过 2-不通过
}
