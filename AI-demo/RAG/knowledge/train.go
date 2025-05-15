package knowledge

import "gorm.io/gorm"

// 训练

// TrainInfo 训练信息
type TrainInfo struct {
	gorm.Model
	TraceID      string // 训练业务TraceID
	TrainModelID uint   // 模型表ID
	TrainCode    string // 训练类型编码: seq2seq, bert, route
	Status       int    // 0-初始化或排队中 1-训练中 2-训练完成 3-训练失败 4-训练取消
	// 5-模型推送中 6-模型推送失败 7-模型推送取消 8-模型推送完成 9-模型准去率未达到阈值不进行推送
	//  10-发布中 11-发布失败 20-生效被替代
	IsValid         int     // 是否有效 0-无效 1-有效
	TrainMode       int     // 训练模式 0-快速训练 1-深度训练
	IsPublishAuto   int     // 是否自动发布 0-手动发起模型训练否 1-定时任务发起模型训练
	BusPublishValue float32 // 自动发布阈值
	IncCount        int     // 参与训练的增量数据数量
	TrainCount      int     // 参与训练的总数据数量
	TrainRequest    string  // 请求模型训练的报文
	TrainResult     string  // 模型训练的响应报文
	TopResult       string  // Top相关信息，只在模型训练成功后记录
}

type TrainKnowledge struct {
	gorm.Model
	TrainID      uint   // 训练ID
	TrainModelID uint   // 模型类型ID
	TraceID      string // 训练业务TraceID
	KnowledgeID  uint   // 知识点ID
	Question     string // 标准问
	Status       int    // 训练状态 0-参与模型训练 1-模型训练后但发布前（等待审核发布）
}

type TrainType struct {
	gorm.Model
	TrainCode  string // 训练类型编码: seq2seq, bert, route
	Name       string // 模型名
	IsValid    int    // 是否有效 0-无效 1-有效
	TrainModel int    // 训练模型 0-不共存，标准和bert二选一 1-共存，路由模型和其他模型共存
}
