package password_symbol

import (
	"rutilus/internal/check"
	"rutilus/internal/util"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type PasswordSymbol struct {
}

func (c *PasswordSymbol) Describe() string {
	return "Password policy must require at least one symbol"
}

func PolicyRequiresSymbol(policy *iam.GetAccountPasswordPolicyOutput) bool {
	return policy.PasswordPolicy.RequireSymbols
}

func (c *PasswordSymbol) Check() (check.CheckResult, error) {
	passwordPolicy := util.GetPasswordPolicy()

	if !PolicyRequiresSymbol(passwordPolicy) {
		return check.Failure, nil
	}
	return check.Success, nil
}
