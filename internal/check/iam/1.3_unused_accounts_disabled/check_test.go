package unused_accounts_disabled

import (
	"rutilus/internal/util"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func timeWithPointer(t time.Time) *time.Time {
	return &t
}

func TestFindUnusedUsers_NoUsers(t *testing.T) {
	report := &util.CredentialReport{}

	assert.Empty(t, FindUnusedUsers(report))
}

func TestFindUnusedUsers_PassWordNotEnabled_UserNotFound(t *testing.T) {
	passwordEnabled := false
	report := &util.CredentialReport{
		Users: []util.CredentialUser{
			{
				UserName:         "someuser",
				PasswordEnabled:  &passwordEnabled,
				PasswordLastUsed: timeWithPointer(time.Now().AddDate(0, -100, 0)),
			},
		},
	}

	assert.Empty(t, FindUnusedUsers(report))
}

func TestFindUnusedUsers_UsersFound(t *testing.T) {
	passwordEnabled := true
	for _, user := range []util.CredentialUser{
		{
			UserName:         "someuser",
			PasswordEnabled:  &passwordEnabled,
			PasswordLastUsed: timeWithPointer(time.Now().AddDate(0, -100, 0)),
		},
		{
			UserName:            "someuser",
			PasswordEnabled:     &passwordEnabled,
			PasswordLastChanged: timeWithPointer(time.Now().AddDate(0, -100, 0)),
		},
		{
			UserName:              "someuser",
			PasswordEnabled:       &passwordEnabled,
			AccessKey1LastRotated: timeWithPointer(time.Now().AddDate(0, -100, 0)),
		},
		{
			UserName:           "someuser",
			PasswordEnabled:    &passwordEnabled,
			AccessKey1LastUsed: timeWithPointer(time.Now().AddDate(0, -100, 0)),
		},
		{
			UserName:              "someuser",
			PasswordEnabled:       &passwordEnabled,
			AccessKey2LastRotated: timeWithPointer(time.Now().AddDate(0, -100, 0)),
		},
		{
			UserName:           "someuser",
			PasswordEnabled:    &passwordEnabled,
			AccessKey2LastUsed: timeWithPointer(time.Now().AddDate(0, -100, 0)),
		},
	} {
		report := &util.CredentialReport{
			Users: []util.CredentialUser{
				user,
			},
		}

		assert.Equal(t, 1, len(FindUnusedUsers(report)))
	}
}
