package knowledge

type IgnoreQuery struct {
	ID        uint   // 主键ID
	Query     string // 用户问句
	Source    int    // 忽略问句来源：0-聚类推荐 1-标注缺失
	VersionNo int    // 乐观锁版本号
}
