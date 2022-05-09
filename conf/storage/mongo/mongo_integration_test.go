//go:build integration

package mongo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartDBSuccessfully(t *testing.T) {
	client, _ := StartDB()
	assert.NotNil(t, client)
}
