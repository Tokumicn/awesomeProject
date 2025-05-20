package vector_db

import (
	"context"
	"github.com/pgvector/pgvector-go"
	"gorm.io/gorm"
)

var db *gorm.DB

func (DocVector) TableName() string {
	return "doc_vectors"
}

func (v *DocVector) Create() (err error) {
	result := db.Create(&v)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (v *DocVector) SimilarQuery(ctx context.Context, qVector pgvector.Vector) ([]DocVector, error) {
	// 使用向量相似度查询
	// 注意：这里使用 PostgreSQL 的向量操作符 <-> 来计算余弦相似度
	var similarVectors []DocVector

	db.Raw("SELECT * FROM doc_vectors ORDER BY embedding <-> ? LIMIT 2", qVector).
		Scan(&similarVectors)

	return similarVectors, nil
}
