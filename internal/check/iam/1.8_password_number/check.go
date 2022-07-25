package password_number

import (
	"rutilus/internal/check"
	"rutilus/internal/util"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type PasswordNumber struct {
}

func (c *PasswordNumber) Describe() string {
	return "Password policy must require at least one number"
}

func PolicyRequiresNumber(policy *iam.GetAccountPasswordPolicyOutput) bool {
	return policy.PasswordPolicy.RequireNumbers
}

func (c *PasswordNumber) Check() (check.CheckResult, error) {
	passwordPolicy := util.GetPasswordPolicy()

	if !PolicyRequiresNumber(passwordPolicy) {
		return check.Failure, nil
	}
	return check.Success, nil
}
