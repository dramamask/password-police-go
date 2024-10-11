package password

import (
	"fmt"

	"github.com/dramamask/password-police-go/misc"
	"github.com/dramamask/password-police-go/policy"
)

// Validate checks if the password adheres to the password policy
func Validate(password string) error {
	policyRules, err := policy.Get()
	if err != nil {
		return err
	}

	err = validateLength(policyRules, password)
	if err != nil {
		return err
	}

	return nil
}

// validateLength checks if the password length adheres to the password policy
func validateLength(policyRules policy.Policy, password string) error {
	if len(password) < policyRules.MinLength {
		return fmt.Errorf("password must be more than %d %s long", policyRules.MinLength, misc.Plurality("character", policyRules.MinLength))
	}

	if len(password) > policyRules.MaxLength {
		return fmt.Errorf("password must be less than %d %s long", policyRules.MaxLength, misc.Plurality("character", policyRules.MaxLength))
	}

	return nil
}
