package root_account_not_used

import (
	"rutilus/internal/util"
	"testing"
	"time"
)

func TestCheckRootAccountUsed_NoRootUser(t *testing.T) {
	report := &util.CredentialReport{}
	_, err := CheckRootAccountUsed(report)
	if err == nil {
		t.Fatalf("Expected error")
	}
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
	if err != nil {
		t.Fatalf("Unexpected error")
	}
	if result {
		t.Fatalf("Expected root account not used")
	}
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
	if err != nil {
		t.Fatalf("Unexpected error")
	}
	if !result {
		t.Fatalf("Expected root account was used")
	}
}
