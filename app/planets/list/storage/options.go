package storage

type MongoOptions struct {
	offset int64
	limit  int64
}

func NewMongoOptions(offset, limit int64) MongoOptions {
	return MongoOptions{offset: offset, limit: limit}
}
