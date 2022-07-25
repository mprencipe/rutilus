package password_lowercase_letter

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/stretchr/testify/assert"
)

func TestPolicyRequiresLowercaseCharacters_Required(t *testing.T) {
	policy := &iam.GetAccountPasswordPolicyOutput{PasswordPolicy: &types.PasswordPolicy{RequireLowercaseCharacters: true}}

	assert.True(t, PolicyRequiresLowercaseCharacters(policy))
}

func TestPolicyRequiresLowercaseCharacters_NotRequired(t *testing.T) {
	policy := &iam.GetAccountPasswordPolicyOutput{PasswordPolicy: &types.PasswordPolicy{RequireLowercaseCharacters: false}}

	assert.False(t, PolicyRequiresLowercaseCharacters(policy))
}
