package policy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDescription(t *testing.T) {
	t.Run("Password can be any length", func(t *testing.T) {
		policy = Policy{
			MinLength: 0,
			MaxLength: 0,
		}

		description := Describe()

		assert.Equal(t, "Password can be any length.", description)
	})

	t.Run("Password must be less than 10 characters long.", func(t *testing.T) {
		policy = Policy{
			MinLength: 0,
			MaxLength: 10,
		}

		description := Describe()

		assert.Equal(t, "Password must be less than 10 characters long.", description)
	})

	t.Run("Password must be more than 10 characters long", func(t *testing.T) {
		policy = Policy{
			MinLength: 10,
			MaxLength: 0,
		}

		description := Describe()

		assert.Equal(t, "Password must be more than 10 characters long.", description)
	})

	t.Run("Password must be between 10 and 20 characters long", func(t *testing.T) {
		policy = Policy{
			MinLength: 10,
			MaxLength: 20,
		}

		description := Describe()

		assert.Equal(t, "Password must be between 10 and 20 characters long.", description)
	})
}
