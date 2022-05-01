package gateway

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFilter(t *testing.T) {
	filter := NewFilter("tattooine", "1")
	assert.Equal(t, "tattooine", filter.name)
	assert.Equal(t, "1", filter.page)
}
