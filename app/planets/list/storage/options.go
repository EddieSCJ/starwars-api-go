package storage

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultOffset = 0
	defaultLimit  = 10
)

type MongoOptions struct {
	offset int64
	limit  int64
}

func NewMongoOptions(offset, limit int64) MongoOptions {
	return MongoOptions{offset: offset, limit: limit}
}

func (m *MongoOptions) Build() *options.FindOptions {
	findOptions := options.Find()
	findOptions.SetSkip(m.offset).SetLimit(m.limit)
	if m.offset < 0 {
		findOptions.SetSkip(defaultOffset)
		m.offset = defaultOffset
	}

	if m.limit == 0 {
		findOptions.SetLimit(defaultLimit)
		m.limit = defaultLimit
	}

	if m.offset > m.limit {
		findOptions.SetLimit(m.offset)
		m.limit = m.offset
	}

	return findOptions
}
