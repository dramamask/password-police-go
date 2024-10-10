package policy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDescription(t *testing.T) {
	t.Run("Password can be any length", func(t *testing.T) {
		policy := Policy{
			MinLength: 0,
			MaxLength: 0,
		}

		description := GetDescription(policy)

		assert.Equal(t, "Password can be any length", description)
	})

	t.Run("Password must be less than 10 characters long", func(t *testing.T) {
		policy := Policy{
			MinLength: 0,
			MaxLength: 10,
		}

		description := GetDescription(policy)

		assert.Equal(t, "Password must be less than 10 characters long", description)
	})

	t.Run("Password must be more than 10 characters long", func(t *testing.T) {
		policy := Policy{
			MinLength: 10,
			MaxLength: 0,
		}

		description := GetDescription(policy)

		assert.Equal(t, "Password must be more than 10 characters long", description)
	})

	t.Run("Password must be between 10 and 20 characters long", func(t *testing.T) {
		policy := Policy{
			MinLength: 10,
			MaxLength: 20,
		}

		description := GetDescription(policy)

		assert.Equal(t, "Password must be between 10 and 20 characters long", description)
	})
}

func TestValidatePolicy(t *testing.T) {
	t.Run("Invalid minimum length", func(t *testing.T) {
		policy := Policy{
			MinLength: -1,
			MaxLength: 10,
		}

		valid := validatePolicy(policy)

		assert.False(t, valid)
	})

	t.Run("Invalid maximum length", func(t *testing.T) {
		policy := Policy{
			MinLength: 10,
			MaxLength: -1,
		}

		valid := validatePolicy(policy)

		assert.False(t, valid)
	})

	t.Run("Minimum length greater than maximum length", func(t *testing.T) {
		policy := Policy{
			MinLength: 20,
			MaxLength: 10,
		}

		valid := validatePolicy(policy)

		assert.False(t, valid)
	})

	t.Run("Valid policy", func(t *testing.T) {
		policy := Policy{
			MinLength: 10,
			MaxLength: 20,
		}

		valid := validatePolicy(policy)

		assert.True(t, valid)
	})
}

func TestGetPlurality(t *testing.T) {
	t.Run("1 character", func(t *testing.T) {
		plurality := getPlurality("character", 1)

		assert.Equal(t, "character", plurality)
	})

	t.Run("2 numbers", func(t *testing.T) {

		plurality := getPlurality("number", 2)

		assert.Equal(t, "numbers", plurality)
	})

	t.Run("5 symbols", func(t *testing.T) {

		plurality := getPlurality("symbol", 5)

		assert.Equal(t, "symbols", plurality)
	})
}
