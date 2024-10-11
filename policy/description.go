package policy

import (
	"fmt"

	"github.com/dramamask/password-police-go/misc"
)

/**
* Example password description
* Password must be a minimum of 20 characters in length with at least one number, one uppercase letter, and one symbol.
 */

// GetDescription returns the description for the password policy
func Describe() (description string) {
	description = getLengthDescription(policy)
	description += "."

	return description
}

// getLengthDescription returns the description for the password length policy
func getLengthDescription(policy Policy) string {
	if policy.MinLength == 0 && policy.MaxLength == 0 {
		return "Password can be any length"
	}

	if policy.MinLength == 0 {
		return fmt.Sprintf("Password must be less than %d %s long", policy.MaxLength, misc.Plurality("character", policy.MaxLength))
	}

	if policy.MaxLength == 0 {
		return fmt.Sprintf("Password must be more than %d %s long", policy.MinLength, misc.Plurality("character", policy.MinLength))
	}

	return fmt.Sprintf("Password must be between %d and %d %s long", policy.MinLength, policy.MaxLength, misc.Plurality("character", policy.MaxLength))
}
