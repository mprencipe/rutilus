package rotated_access_keys

import (
	"rutilus/internal/check"
	"rutilus/internal/util"
	"time"
)

type RotatedAccessKeys struct {
}

func (c *RotatedAccessKeys) Describe() string {
	return "Access keys are rotated every 90 days or less"
}

func FindNonRotatedAccessKeys(report *util.CredentialReport) []util.CredentialUser {
	ret := make([]util.CredentialUser, 0)
	for _, u := range report.Users {
		if util.AnyTimeIsNotNilAndLaterThan(
			time.Now().AddDate(0, 0, -90), []*time.Time{
				u.AccessKey1LastRotated,
				u.AccessKey2LastRotated,
			}) {
			ret = append(ret, u)
		}

	}
	return ret
}

func (c *RotatedAccessKeys) Check() (check.CheckResult, error) {
	report := util.GetCredentialReport()
	users := FindNonRotatedAccessKeys(report)
	if len(users) > 0 {
		return check.Failure, nil
	}
	return check.Success, nil
}
