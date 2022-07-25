package root_account_not_used

import (
	"rutilus/internal/util"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCheckRootAccountUsed_NoRootUser(t *testing.T) {
	report := &util.CredentialReport{}

	_, err := CheckRootAccountUsed(report)

	assert.NotNil(t, err, "Expected error")
}

func TestCheckRootAccountUsed_NotRecentlyUsed(t *testing.T) {
	accessKeyLastUsed := time.Now().AddDate(0, -1, 0)
	report := &util.CredentialReport{Users: []util.CredentialUser{
		{
			UserName:           "<root_account>",
			AccessKey1LastUsed: &accessKeyLastUsed,
		},
	}}

	result, err := CheckRootAccountUsed(report)

	assert.Nil(t, err, "Unexpected error")
	assert.False(t, result, "Expected root account not used")
}

func TestCheckRootAccountUsed_RecentlyUsed(t *testing.T) {
	accessKeyLastUsed := time.Now().AddDate(0, 0, -2)
	report := &util.CredentialReport{Users: []util.CredentialUser{
		{
			UserName:           "<root_account>",
			AccessKey1LastUsed: &accessKeyLastUsed,
		},
	}}

	result, err := CheckRootAccountUsed(report)

	assert.Nil(t, err, "Unexpected error")
	assert.True(t, result, "Expected root account used")
}
