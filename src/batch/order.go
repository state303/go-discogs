package batch

import (
	"context"
	"gorm.io/gorm"
)

type Order interface {
	getContext() context.Context
	getChunkSize() int
	getFilePath() string
	getDB() *gorm.DB
}

type orderImpl struct {
	ctx       context.Context
	chunkSize int
	filepath  string
	db        *gorm.DB
}

func (o *orderImpl) getContext() context.Context {
	return o.ctx
}

func (o *orderImpl) getChunkSize() int {
	return o.chunkSize
}

func (o *orderImpl) getFilePath() string {
	return o.filepath
}

func (o *orderImpl) getDB() *gorm.DB {
	return o.db.Session(&gorm.Session{})
}

func NewOrder(ctx context.Context, chunkSize int, filepath string, db *gorm.DB) Order {
	return &orderImpl{
		ctx:       ctx,
		chunkSize: chunkSize,
		filepath:  filepath,
		db:        db,
	}
}
