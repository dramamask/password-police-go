package main

import (
	"fmt"

	"github.com/dramamask/password-police/policy"
)

func main() {
	password := "1234567"
	policyRules := policy.ParsePolicy()

	//fmt.Printf("Password policy: %v\n", policyRules)

	fmt.Println(policy.GetDescription(policyRules))

	if policy.AdheresToPolicy(password, policyRules) {
		fmt.Println("Password is okay")
		return
	}

	fmt.Println("Password is not okay")

}
