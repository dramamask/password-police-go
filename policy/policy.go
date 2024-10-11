package policy

import (
	"encoding/json"
	"fmt"
	"os"
)

var policy Policy = Policy{}

type Policy struct {
	MinLength int `json:"min-length"`
	MaxLength int `json:"max-length"`
}

// Init reads the password policy from the supplied file name
func LoadFromFile(fileName string) error {
	policyJson, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("error reading the password policy file. Make sure the file exists")
	}

	var policyFromFile Policy
	err = json.Unmarshal(policyJson, &policyFromFile)
	if err != nil {
		return fmt.Errorf("error parsing the password policy file. Make sure the file is valid JSON")
	}

	return Set(policyFromFile)
}

// Set takes the supplied password policy, validates it, and assigns it to the policy variable
func Set(myPolicy Policy) error {
	err := validatePolicy(myPolicy)
	if err != nil {
		return err
	}

	policy = myPolicy

	return nil
}

// Get returns the password policy (if it has been initialized)
func Get() (Policy, error) {
	if policy == (Policy{}) {
		return (Policy{}), fmt.Errorf("error: Password policy has not been initialized")
	}

	return policy, nil
}

// validatePolicy checks if the policy is valid
func validatePolicy(policy Policy) error {
	if policy.MinLength < 0 {
		return fmt.Errorf("error: Invalid password policy. Minimum length must be greater than or equal to 0")
	}

	if policy.MaxLength < 0 {
		return fmt.Errorf("error: Invalid password policy. Maximum length must be greater than or equal to 0")
	}

	if policy.MinLength > policy.MaxLength {
		return fmt.Errorf("error: Invalid password policy. Minimum length must be less than or equal to the maximum length")
	}

	return nil
}
