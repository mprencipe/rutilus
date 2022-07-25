package root_account_not_used

import (
	"errors"
	"rutilus/internal/check"
	"rutilus/internal/util"
	"time"
)

type RootAccountNotUsed struct {
}

func (c *RootAccountNotUsed) Describe() string {
	return "Root users should not be used"
}

func checkAccessKeyUsed(accessKeyUsedTime *time.Time, limit time.Time) bool {
	if accessKeyUsedTime == nil {
		return false
	}
	return accessKeyUsedTime.After(limit)
}

func CheckRootAccountUsed(report *util.CredentialReport) (bool, error) {
	rootAccountFound := false
	for _, u := range report.Users {
		if u.UserName == "<root_account>" {
			rootAccountFound = true
			limit := time.Now().AddDate(0, 0, -7)
			if checkAccessKeyUsed(u.AccessKey1LastUsed, limit) || checkAccessKeyUsed(u.AccessKey2LastUsed, limit) {
				return true, nil
			} else {
				return false, nil
			}
		}
	}
	if !rootAccountFound {
		return false, errors.New("Root account not found")
	}
	return true, nil
}

func (c *RootAccountNotUsed) Check() (check.CheckResult, error) {
	report := util.GetCredentialReport()
	rootAccountUsed, err := CheckRootAccountUsed(report)
	if err != nil {
		return check.Failure, err
	}

	if rootAccountUsed {
		return check.Failure, nil
	}

	return check.Success, nil
}
