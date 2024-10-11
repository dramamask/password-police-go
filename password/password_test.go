package password

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dramamask/password-police-go/policy"
)

func TestValidate(t *testing.T) {
	t.Run("Password policy has not been initialized", func(t *testing.T) {
		password := "1234567"
		err := Validate(password)

		assert.NotNil(t, err)
	})

	t.Run("Password is okay", func(t *testing.T) {
		myPolicy := policy.Policy{
			MinLength: 4,
			MaxLength: 8,
		}
		err := policy.Set(myPolicy)
		assert.Nil(t, err)

		password := "1234567"
		err = Validate(password)

		assert.Nil(t, err)
	})

	t.Run("Password is too short", func(t *testing.T) {
		myPolicy := policy.Policy{
			MinLength: 4,
			MaxLength: 8,
		}
		err := policy.Set(myPolicy)
		assert.Nil(t, err)

		password := "123"
		err = Validate(password)

		assert.NotNil(t, err)
	})

	t.Run("Password is too long", func(t *testing.T) {
		myPolicy := policy.Policy{
			MinLength: 4,
			MaxLength: 8,
		}
		err := policy.Set(myPolicy)
		assert.Nil(t, err)

		password := "123456789"
		err = Validate(password)

		assert.NotNil(t, err)
	})

	t.Run("Password is empty", func(t *testing.T) {
		myPolicy := policy.Policy{
			MinLength: 4,
			MaxLength: 8,
		}
		err := policy.Set(myPolicy)
		assert.Nil(t, err)

		password := ""
		err = Validate(password)

		assert.NotNil(t, err)
	})
}
