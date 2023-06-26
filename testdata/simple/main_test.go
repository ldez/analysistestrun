package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// BUG: here // want "There is a bug"
func TestSimple(t *testing.T) {
	assert.True(t, true)
}
