package password_lowercase_letter

import (
	"rutilus/internal/check"
	"rutilus/internal/util"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type PasswordLowercaseLetter struct {
}

func (c *PasswordLowercaseLetter) Describe() string {
	return "Password policy must require at least one lowercase letter"
}

func PolicyRequiresLowercaseCharacters(policy *iam.GetAccountPasswordPolicyOutput) bool {
	return policy.PasswordPolicy.RequireLowercaseCharacters
}

func (c *PasswordLowercaseLetter) Check() (check.CheckResult, error) {
	passwordPolicy := util.GetPasswordPolicy()

	if !PolicyRequiresLowercaseCharacters(passwordPolicy) {
		return check.Failure, nil
	}
	return check.Success, nil
}
