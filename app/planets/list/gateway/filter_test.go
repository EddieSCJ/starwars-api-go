package gateway

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFilter(t *testing.T) {
	t.Parallel()
	filter := NewFilter("tattooine", "1")
	assert.Equal(t, "tattooine", filter.name)
	assert.Equal(t, "1", filter.page)
}
