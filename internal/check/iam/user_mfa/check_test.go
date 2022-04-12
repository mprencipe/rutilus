package user_mfa

import (
	"rutilus/internal/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindUsersWithoutMFA_NoUsers(t *testing.T) {
	report := &util.CredentialReport{}

	assert.Empty(t, FindUsersWithoutMFA(report))
}

func TestFindUsersWithoutMFA_NoPasswordNoMFA(t *testing.T) {
	mfa := false
	passwordEnabled := false
	report := &util.CredentialReport{
		Users: []util.CredentialUser{
			{
				UserName:        "someuser",
				MfaActive:       &mfa,
				PasswordEnabled: &passwordEnabled,
			},
		},
	}

	assert.Empty(t, FindUsersWithoutMFA(report))
}

func TestFindUsersWithoutMFA_PasswordAndMFA(t *testing.T) {
	mfa := true
	passwordEnabled := true
	report := &util.CredentialReport{
		Users: []util.CredentialUser{
			{
				UserName:        "someuser",
				MfaActive:       &mfa,
				PasswordEnabled: &passwordEnabled,
			},
		},
	}

	assert.Empty(t, FindUsersWithoutMFA(report))
}

func TestFindUsersWithoutMFA_PasswordButNoMFA(t *testing.T) {
	mfa := false
	passwordEnabled := true
	report := &util.CredentialReport{
		Users: []util.CredentialUser{
			{
				UserName:        "someuser",
				MfaActive:       &mfa,
				PasswordEnabled: &passwordEnabled,
			},
		},
	}

	assert.Equal(t, 1, len(FindUsersWithoutMFA(report)))
}
