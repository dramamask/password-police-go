package misc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPlurality(t *testing.T) {
	t.Run("1 character", func(t *testing.T) {
		plurality := Plurality("character", 1)

		assert.Equal(t, "character", plurality)
	})

	t.Run("2 numbers", func(t *testing.T) {

		plurality := Plurality("number", 2)

		assert.Equal(t, "numbers", plurality)
	})

	t.Run("5 symbols", func(t *testing.T) {

		plurality := Plurality("symbol", 5)

		assert.Equal(t, "symbols", plurality)
	})
}
