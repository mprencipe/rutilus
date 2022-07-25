package password_uppercase_letter

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/stretchr/testify/assert"
)

func TestPolicyRequiresUppercaseCharacters_Required(t *testing.T) {
	policy := &iam.GetAccountPasswordPolicyOutput{PasswordPolicy: &types.PasswordPolicy{RequireUppercaseCharacters: true}}

	assert.True(t, PolicyRequiresUppercaseCharacters(policy))
}

func TestPolicyRequiresUppercaseCharacters_NotRequired(t *testing.T) {
	policy := &iam.GetAccountPasswordPolicyOutput{PasswordPolicy: &types.PasswordPolicy{RequireUppercaseCharacters: false}}

	assert.False(t, PolicyRequiresUppercaseCharacters(policy))
}
