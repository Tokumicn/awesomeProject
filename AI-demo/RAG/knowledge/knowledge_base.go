package knowledge

import "gorm.io/gorm"

// KnowledgeBase 资料库表
type KnowledgeBase struct {
	gorm.Model
	Name       string // 知识库名称''
	ViewType   int    // 知识库展示类型 0-列表 1-卡片 3-表格
	CategoryID uint   // 知识库分类ID
}

// KnowledgeBaseIntention 知识库意图表
type KnowledgeBaseIntention struct {
	gorm.Model
	IntentionID      int    // 意图ID
	KnowledgeBaseID  uint   // 知识库ID
	RejectAnswer     string // 拒绝意图话术
	RejectAnswerType int    // 拒绝意图话术类型 0-纯文本 1-富文本 2-图文消息 3-单图片
}

// 资料库词槽表
type KnowledgeBaseSlot struct {
	gorm.Model
	KnowledgeBaseID   uint   // 知识库ID
	SlotID            uint   // 词槽ID
	MissingAnswer     string // 缺失词槽话术
	MissingAnswerType int    // 缺失词槽话术类型 0-纯文本 1-富文本 2-图文消息 3-单图片
	SlotSort          uint   // 词槽卡片排序0，1，2，3，4...
}
