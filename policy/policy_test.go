package policy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	t.Run("Non existant policy file", func(t *testing.T) {
		err := LoadFromFile("non-existant-policy-file.json")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "Make sure the file exists")
	})

	t.Run("Empty policy file", func(t *testing.T) {
		err := LoadFromFile("./test/empty-pwd-policy.json")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "Make sure the file is valid JSON")
	})

	t.Run("Malformed JSON policy file", func(t *testing.T) {
		err := LoadFromFile("./test/malformed-json-pwd-policy.json")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "Make sure the file is valid JSON")
	})
}

func TestValidatePolicy(t *testing.T) {
	t.Run("Invalid minimum length", func(t *testing.T) {
		policy := Policy{
			MinLength: -1,
			MaxLength: 10,
		}

		err := validatePolicy(policy)

		assert.NotNil(t, err)
	})

	t.Run("Invalid maximum length", func(t *testing.T) {
		policy := Policy{
			MinLength: 10,
			MaxLength: -1,
		}

		err := validatePolicy(policy)

		assert.NotNil(t, err)
	})

	t.Run("Minimum length greater than maximum length", func(t *testing.T) {
		policy := Policy{
			MinLength: 20,
			MaxLength: 10,
		}

		err := validatePolicy(policy)

		assert.NotNil(t, err)
	})

	t.Run("Valid policy", func(t *testing.T) {
		policy := Policy{
			MinLength: 10,
			MaxLength: 20,
		}

		err := validatePolicy(policy)

		assert.Nil(t, err)
	})

}
