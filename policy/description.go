package policy

import (
	"fmt"
)

/**
* Example password description
* Password must be a minimum of 20 characters in length with at least one number, one uppercase letter, and one symbol.
 */

// GetDescription returns the description for the password policy
func GetDescription(policy Policy) (description string) {
	return getLengthDescription(policy)
}

// getLengthDescription returns the description for the password length policy
func getLengthDescription(policy Policy) string {
	if policy.MinLength == 0 && policy.MaxLength == 0 {
		return "Password can be any length"
	}

	if policy.MinLength == 0 {
		return fmt.Sprintf("Password must be less than %d %s long", policy.MaxLength, getPlurality("character", policy.MaxLength))
	}

	if policy.MaxLength == 0 {
		return fmt.Sprintf("Password must be more than %d %s long", policy.MinLength, getPlurality("character", policy.MinLength))
	}

	return fmt.Sprintf("Password must be between %d and %d %s long", policy.MinLength, policy.MaxLength, getPlurality("character", policy.MaxLength))
}

// getPlurality returns the word with the correct plurality based on the amount
func getPlurality(word string, amount int) string {
	if amount == 1 {
		return word
	}

	return word + "s"
}
