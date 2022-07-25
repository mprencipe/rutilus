package unused_accounts_disabled

import (
	"rutilus/internal/check"
	"rutilus/internal/util"
	"time"
)

type UnusedAccountsDisabled struct {
}

func (c *UnusedAccountsDisabled) Describe() string {
	return "Unused users (90 days) are disabled"
}

func FindUnusedUsers(report *util.CredentialReport) []util.CredentialUser {
	ret := make([]util.CredentialUser, 0)
	for _, u := range report.Users {
		if u.PasswordEnabled != nil && *u.PasswordEnabled {
			if util.AnyTimeIsNotNilAndLaterThan(
				time.Now().AddDate(0, 0, -90), []*time.Time{
					u.PasswordLastUsed,
					u.PasswordLastChanged,
					u.AccessKey1LastRotated,
					u.AccessKey1LastUsed,
					u.AccessKey2LastRotated,
					u.AccessKey2LastUsed,
				}) {
				ret = append(ret, u)
			}
		}
	}
	return ret
}

func (c *UnusedAccountsDisabled) Check() (check.CheckResult, error) {
	report := util.GetCredentialReport()
	users := FindUnusedUsers(report)
	if len(users) > 0 {
		return check.Failure, nil
	}
	return check.Success, nil
}
