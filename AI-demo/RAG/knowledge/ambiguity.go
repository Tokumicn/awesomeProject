package knowledge

import "gorm.io/gorm"

// Ambiguity 歧义问
type Ambiguity struct {
	gorm.Model
	Question      string  // 意图相似问或标注问
	CategoryID    uint    // 分类ID
	DocID         uint    // 意图ID
	IntentionID   uint    // 歧义对ID
	Status        int     // 状态 0-默认 1-交换 2-转移 3-删除 4-左移 5-右移
	Version       int     // 乐观锁版本号
	AmbiguityRate float64 // 歧义率
	Type          int     // 问句类型 TODO
}
