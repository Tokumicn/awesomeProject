package dto

import "gorm.io/gorm"

// Knowledge 知识
type Knowledge struct {
	gorm.Model
	CategoryID      string `gorm:"type:int(11) column:category_id" json:"category_id"`              // 分类ID
	Question        string `gorm:"type:text" json:"question"`                                       // 标准问
	KnowledgeType   string `gorm:"type:tinyint(4) column:knowledge_type" json:"knowledge_type"`     // FAQ类型 0-标准FAQ 1-寒暄 2-任务型
	KnowledgeStatus uint8  `gorm:"type:tinyint(4) column:knowledge_status" json:"knowledge_status"` // 知识状态
	// 100-未生效 110-已生效 120-编辑中
	// 200-上线审核 210-更新审核 220-下线审核
	// 300-上线发布中 310-更新发布中 320-下线发布中
	// 400-上线驳回 410-更新驳回 420-下线驳回
	// 500-上线失败 510-更新失败 520-下线失败
	Deactivate   uint8 `gorm:"type:tinyint(4) column:deactivate" json:"deactivate"`      // 是否已停用该FAQ答案 0-未停用 1-停用
	CreateUserID uint  `gorm:"type:int(11) column:create_user_id" json:"create_user_id"` // 创建人ID
	UpdateUserID uint  `gorm:"type:int(11) column:update_user_id" json:"update_user_id"` // 更新人ID
}

// KnowledgeAnswer 知识的答案
type KnowledgeAnswer struct {
	gorm.Model
	KnowledgeID     uint   `gorm:"type:int(11) column:knowledge_id" json:"knowledge_id"`                 // 知识ID
	Answer          string `gorm:"type:mediumtext column:answer" json:"answer"`                          // answer 答案
	AnswerType      uint8  `gorm:"type:int(2) column:answer_type" json:"answer_type"`                    // answer_type 答案类型 0-纯文本 1-富文本 2-图文 3-单图片
	RelatedQuestion string `gorm:"type:varchar(2048re) column:related_question" json:"related_question"` // related_question 关联问
}

// KnowledgeSimilar 相似问
type KnowledgeSimilar struct {
	gorm.Model
	KnowledgeID uint   `gorm:"type:int(11) column:knowledge_id" json:"knowledge_id"` // 知识ID
	Question    string `gorm:"type:text" json:"question"`                            // 相似问
}
