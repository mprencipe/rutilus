package rotated_access_keys

import (
	"rutilus/internal/util"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func timeWithPointer(t time.Time) *time.Time {
	return &t
}

func TestFindNonRotatedAccessKeys_NoUsers(t *testing.T) {
	report := &util.CredentialReport{}

	assert.Empty(t, FindNonRotatedAccessKeys(report))
}

func TestFindNonRotatedAccessKeys_AccessKeyRotated_UserNotFound(t *testing.T) {
	report := &util.CredentialReport{
		Users: []util.CredentialUser{
			{
				UserName:              "someuser",
				AccessKey1LastRotated: timeWithPointer(time.Now().AddDate(0, -25, 0)),
			},
		},
	}

	assert.Equal(t, 1, len(FindNonRotatedAccessKeys(report)))
}

func TestFindNonRotatedAccessKeys_UsersFound(t *testing.T) {
	for _, user := range []util.CredentialUser{
		{
			UserName:              "someuser1",
			AccessKey1LastRotated: timeWithPointer(time.Now().AddDate(0, -100, 0)),
			AccessKey2LastRotated: nil,
		},
		{
			UserName:              "someuser2",
			AccessKey1LastRotated: nil,
			AccessKey2LastRotated: timeWithPointer(time.Now().AddDate(0, -100, 0)),
		},
	} {
		report := &util.CredentialReport{
			Users: []util.CredentialUser{
				user,
			},
		}

		assert.Equal(t, 1, len(FindNonRotatedAccessKeys(report)))
	}
}
