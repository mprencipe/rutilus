package password_uppercase_letter

import (
	"rutilus/internal/check"
	"rutilus/internal/util"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type PasswordUppercaseLetter struct {
}

func (c *PasswordUppercaseLetter) Describe() string {
	return "Password policy must require at least one uppercase letter"
}

func PolicyRequiresUppercaseCharacters(policy *iam.GetAccountPasswordPolicyOutput) bool {
	return policy.PasswordPolicy.RequireUppercaseCharacters
}

func (c *PasswordUppercaseLetter) Check() (check.CheckResult, error) {
	passwordPolicy := util.GetPasswordPolicy()

	if !PolicyRequiresUppercaseCharacters(passwordPolicy) {
		return check.Failure, nil
	}
	return check.Success, nil
}
