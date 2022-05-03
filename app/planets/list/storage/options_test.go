package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestNewMongoOptions(t *testing.T) {
	mongoOptions := NewMongoOptions(3, 12)
	assert.Equal(t, int64(3), mongoOptions.offset)
	assert.Equal(t, int64(12), mongoOptions.limit)
}

func TestBuild(t *testing.T) {
	mongoOptions := NewMongoOptions(3, 12)
	expected := options.Find().SetSkip(mongoOptions.offset).SetLimit(mongoOptions.limit)

	result := mongoOptions.Build()
	assert.Equal(t, *expected.Skip, *result.Skip)
	assert.Equal(t, *expected.Limit, *result.Limit)
}
