package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitCache(t *testing.T) {
	t.Parallel()
	t.Run("Test in memory cache", func(t *testing.T) {
		InitCache()
		GetCache().Set("A", 1)
		a := GetCache().Get("A")
		assert.Equal(t, a, 1)
	})
}
