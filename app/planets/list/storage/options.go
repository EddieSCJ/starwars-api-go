package storage

import "go.mongodb.org/mongo-driver/mongo/options"

type MongoOptions struct {
	offset int64
	limit  int64
}

func NewMongoOptions(offset, limit int64) MongoOptions {
	return MongoOptions{offset: offset, limit: limit}
}

func (m *MongoOptions) Build() *options.FindOptions {
	return options.Find().
		SetSkip(m.offset).
		SetLimit(m.limit)
}
