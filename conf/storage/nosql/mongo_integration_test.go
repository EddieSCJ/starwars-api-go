//go:build integration

package nosql

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStartDBSuccessfully(t *testing.T) {
	client, _ := StartDB()
	assert.NotNil(t, client)
}
