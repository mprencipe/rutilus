package user_mfa

import (
	"rutilus/internal/check"
	"rutilus/internal/util"
)

type UsersWithPasswordsHaveMFA struct {
}

func (c *UsersWithPasswordsHaveMFA) Describe() string {
	return "Users with passwords should have MFA"
}

func FindUsersWithoutMFA(report *util.CredentialReport) []util.CredentialUser {
	ret := make([]util.CredentialUser, 0)
	for _, u := range report.Users {
		if u.PasswordEnabled != nil && *u.PasswordEnabled {
			if u.MfaActive != nil && !*u.MfaActive {
				ret = append(ret, u)
			}
		}
	}
	return ret
}

func (c *UsersWithPasswordsHaveMFA) Check() (check.CheckResult, error) {
	report := util.GetCredentialReport()
	users := FindUsersWithoutMFA(report)
	if len(users) > 0 {
		return check.Failure, nil
	}
	return check.Success, nil
}
