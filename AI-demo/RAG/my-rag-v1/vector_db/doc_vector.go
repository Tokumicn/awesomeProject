package vector_db

import (
	"context"
	"github.com/pgvector/pgvector-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
)

// DocVector 定义向量数据模型
type DocVector struct {
	gorm.Model
	DocID       uint            `gorm:"column:doc_id"`       // 文档ID
	BlockID     uint            `gorm:"column:block_id"`     // 块ID
	Content     string          `gorm:"type:text"`           // 块内容
	Embedding   pgvector.Vector `gorm:"type:vector"`         // 块内容向量
	TrainStatus int             `gorm:"column:train_status"` // 训练状态
}

func NewDB(ctx context.Context) (*gorm.DB, error) {
	// 连接数据库
	var err error
	dsn := "host=localhost user=chatwiki password=postgres_password dbname=my_rag port=15432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.ErrorContext(ctx, "Failed to connect to database:", err)
		return nil, err
	}

	// 创建 pgvector 扩展
	err = db.Exec("CREATE EXTENSION IF NOT EXISTS vector").Error
	if err != nil {
		slog.ErrorContext(ctx, "Failed to create vector extension:", err)
		return nil, err
	}

	return db, nil
}

func AutoMigrate(ctx context.Context) error {
	// 自动迁移数据库结构
	err := db.AutoMigrate(&DocVector{})
	if err != nil {
		slog.ErrorContext(ctx, "Failed to migrate database:", err)
		return err
	}
	return nil
}
