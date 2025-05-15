package knowledge

import (
	"gorm.io/gorm"
	"time"
)

// MarkInfo 标注库信息
type MarkInfo struct {
	gorm.Model
	Source    int    // 标注库来源: 0-常规教育 1-人工质检 2-外部导入
	SourceID  uint   // 标注库来源ID 导入时为0
	DocID     uint   // FAQ的数据库ID
	Query     string // 标注库问句
	VersionNo int    // 乐观锁版本号
	IsRNN     int    // 是否为RNN 1-RNN 0-GBRT
}

type MakeQuery struct {
	gorm.Model
	MarkTaskID    uint      // 标注任务ID
	SearchID      uint      // 问句ID
	Query         string    // 问句
	SessionRound  int       // 会话轮次
	SessionID     string    // 会话ID
	QueryTime     time.Time // 提问时间
	UserID        uint      // 用户ID
	Source        int       // 问句来源：0-推荐标注 1-转人工 2-点踩 3-拒识 4-外部数据 5-未识别 6-人工质检
	MarkIsSame    int       // 所有标注人结果是否一致 0-一致 1-不一致
	DialogIsSame  int       // 所有对话结果是否一致 0-一致 1-不一致
	DialogResult  string    // 机器人回答完整答案
	MarkType      int       // 标注类型 0-待标注 1-忽略 2-缺失 3-匹配
	MarkResult    string    // 标注结果 eg: [{ docID=12, querytion="你好" }]
	KnowledgeType int       // 知识库类型 0-Faq标准问 1-寒暄 2-资料意图 3-任务型
	VersionNo     int       // 乐观锁版本号
}

// MarkQueryExternal 外部导入标注
type MarkQueryExternal struct {
	gorm.Model
	Query     string // 问句
	HasInMark int    // 是否已在标注任务重 0-否 1-是
}

// MarkUser 标注人信息
type MarkUser struct {
	gorm.Model
	MarkTaskID uint // 标注任务ID
	UserID     uint // 用户ID
	MarkStatus int  // 标注状态 0-未标注 1-标注中 2-标注完成
}

// MarkTask 标注任务
type MarkTask struct {
	gorm.Model
	Name            string    // 标注任务名称
	Type            int       // 标注任务类型 0-标注 1-质检
	IsRecommend     int       // 是否智能推荐 0-不使用 1-使用
	MarkCount       int       // 标注数据量 标注任务的问句数量，质检任务则为会话数量
	Source          string    // 标注任务来源 0-推荐标注 1-转人工 2-点踩 3-拒识 4-外部数据
	Status          int       // 标注任务状态 0-未开始 1-进行中 2-已完成 3-待审核 4-审核中 5-审核完成
	MarkFinishTime  time.Time // 标注任务完成时间
	AuditFinishTime time.Time // 审核任务完成时间
}
