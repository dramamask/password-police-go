package policy

import (
	"encoding/json"
	"fmt"
	"os"
)

type Policy struct {
	MinLength int `json:"min-length"`
	MaxLength int `json:"max-length"`
}

// ParsePolicy reads the password policy from the 'pwd-policy.json' file
func ParsePolicy() (policy Policy) {
	policyJson, err := os.ReadFile("pwd-policy.json")
	if err != nil {
		fmt.Println("Error reading password policy file 'pwd-policy.json'. Make sure the file exists.")
	}

	err = json.Unmarshal(policyJson, &policy)
	if err != nil {
		fmt.Println("Error parsing password policy file 'pwd-policy.json'. Make sure the file is valid JSON.")
	}

	if !validatePolicy(policy) {
		fmt.Println("Error: Invalid password policy. Make sure all the rules are valid. For example, is the minimum length less than the maximum length?")
	}

	return policy
}

// validatePolicy checks if the policy is valid
func validatePolicy(policy Policy) bool {
	if policy.MinLength < 0 {
		return false
	}

	if policy.MaxLength < 0 {
		return false
	}

	if policy.MinLength > policy.MaxLength {
		return false
	}

	return true
}

// AdheresToPolicy checks if the password adheres to the policy
func AdheresToPolicy(password string, policy Policy) bool {
	if len(password) < policy.MinLength {
		return false
	}

	if len(password) > policy.MaxLength {
		return false
	}

	return true
}
