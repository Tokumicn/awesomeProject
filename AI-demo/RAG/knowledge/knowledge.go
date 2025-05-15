package knowledge

import "gorm.io/gorm"

// Knowledge 知识点 标准问
type Knowledge struct {
	gorm.Model
	Question         string  // 问题
	CategoryID       uint    // 分类ID
	Type             int     // 知识类型
	Status           int     // 状态
	Version          int     // 乐观锁版本号
	IntentionID      uint    // 意图ID
	InputIntentionID uint    // 导入时的意图ID
	Accuracy         float64 // 准确率
	Deactivated      bool    // 是否被禁用
	EduCount         int     // 教育数量：相似问+标注问
}

type Answer struct {
	Answer          string // 回答
	Type            int    // 回答类型 0-默认答案 10-其他答案
	RelatedQuestion string // 答案的关联问句
	AnswerType      int    // 回答类型 0-纯文本 1-富文本 2-图文消息 3-单图片
	KnowledgeID     uint   // 知识ID
}

// Category 分类表
type Category struct {
	gorm.Model
	Name          string // 分类名称
	Order         int    // 同层(同 parent_id )显示顺序
	ParentID      uint   // 父级ID
	KnowledgeType int    // 知识类型: 0-标准问 1-标注问 2-意图 3-任务型 10-寒暄库 20-资料库 22-任务词槽 23-资料词槽 30-流程意图
	Description   string // 分类描述
}

// Similar 知识库相似问
type Similar struct {
	gorm.Model
	KnowledgeID uint   // 知识ID
	Question    string // 相似问
	Type        int    // 相似问类型：0-相似问 1-推荐相似问(用户配置)
}

type Relation struct {
	gorm.Model
	KnowledgeID       uint   // 知识ID
	RelateKnowledgeID uint   // 关联知识点ID
	Type              int    // 关联类型：1-知识点关联问 2-默认回答关联问
	Question          string // 关联问
	RelatedDisplay    string // 关联显示
	AnswerID          string // 答案ID
}

type DiyRelation struct {
	gorm.Model
	KnowledgeID int64  // 知识ID
	Type        int    // 1-知识点关联问 2-默认回答关联问
	Question    string // 关联问
	AnswerID    uint   // 知识点的答案ID
	Order       int    // 排序
	QueryStatus int    // 查询状态：0-默认 1-精准命中 2-推荐回答 3-拒绝识别
}

type SlotInstance struct {
	gorm.Model
	KonowledgeBaseID  uint   // 资料库ID
	SlotInstanceID1   uint   // 词槽实例ID1
	SlotInstanceName1 string // 词槽实例名称1
	SlotInstanceID2   uint   // 词槽实例ID2
	SlotInstanceName2 string // 词槽实例名称2
	KnowledgeID       uint   // 关联的资料库知识点ID
	Question          string // 知识点名称(猜测是标准问)
	IntentionID       uint   // 意图ID
}
