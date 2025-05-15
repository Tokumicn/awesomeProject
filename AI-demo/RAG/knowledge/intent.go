package knowledge

import (
	"gorm.io/gorm"
	"time"
)

// IntentDiscovery 意图发现表
type IntentDiscovery struct {
	gorm.Model
	TraceID          string    // 流水ID
	ClusterID        string    // 运行日志返回：8位日期 + task_id + 类id
	IntentDate       time.Time // 意图任务执行日期
	NewUserQueryRate float64   // 该类中新用户问题占比
	Score            float32   // 类打分、类大小、类内聚、类间距、标注占比、缺失占比 各自乘以一个权重，之后求和，再归一化一个分数
	Question         string    // 推荐标准问，处于类中心的问句
	DoFlag           int       // 0-待处理 1-忽略
	StatusResult     string    // 状态回调保温
	Status           int       // 数据运营台处理相应状态：0-初始 1-运行 2-成功 3-失败
	QueryList        string    // 用户问句集合
	DataResult       string    // 数据运营台回调报文
	VersionNo        int       // 乐观锁版本号
}

// IntentDiscoveryDetail 意图发现详情
type IntentDiscoveryDetail struct {
	gorm.Model
	DiscoveryID uint    // 意图发现ID
	Query       string  // 用户问句
	Source      string  // 用户问句类型： query-用户问题 lack-标注缺失问句 reject-拒识问句
	Distance    float32 // 用户问句到类中心的距离
	VersionNo   int     // 乐观锁版本号
}
