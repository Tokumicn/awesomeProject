package knowledge

import (
	"gorm.io/gorm"
	"time"
)

type BadCase struct {
	gorm.Model
	CaseID         string    // badcase唯一标识 时间+随机数生成
	Type           string    // 类型： intentMistake-意图误触 unidentifiedSlot-词槽未识别 slotMistake-词槽误触 slotLack-词槽缺失 errorCondition-跳转条件错误
	SessionSource  int       // 来源：1-会话质检 2-常规教育
	SessionID      string    // 会话ID
	Query          string    // 用户问题
	QueryIndex     int       // 问句在会话中的下标位置
	ReportUserID   uint      // 上报用户ID
	ReportUserName string    // 上报用户姓名
	ReportTime     time.Time // 上报时间
	FinishUserID   uint      // 完成用户ID
	FinishUserName string    // 完成用户姓名
	FinishTime     time.Time // 完成时间
	Status         string    // 状态：ignore-忽略 confirm-待确认 repair-待修复 verify-待验证 filing-已归档 resolve-已解决
	Replay         int       // 是否重播过： 1-是 0-否
	ReplayInfo     string    // 重播信息
	VerifyReplay   int       // 是否验证时重播过： 1-是 0-否
	VerifyType     int       // 验证类型：1-手动验证 2-测试集批量验证
	TestID         uint      // 测试集ID
	SessionTime    time.Time // 会话时间
	ReasonType     string    // 原因类型：1-意图歧义 2-标注数据有误 3-教育数量不足 4-其他
	OtherReason    string    // 其他原因 当ReasonType=4 填写该字段
	Remark         string    // 备注
}

type BadCaseQuery struct {
	gorm.Model
	SameID        string    // 表示query来自于同一个会话
	Query         string    // 用户问题
	QueryIndex    int       // 问句在会话中的下标位置
	QueryTime     time.Time // 问题时间
	AnswerContent string    // 机器人回答内容
	SessionID     string    // 会话ID
}

type BadCaseQuery struct {
	ID        uint // 自增主键ID
	BadCaseID uint // badcase_id
	QueryID   uint // query_id
}
