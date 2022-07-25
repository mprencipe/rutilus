package password_number

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/stretchr/testify/assert"
)

func TestPolicyRequiresSymbol_Required(t *testing.T) {
	policy := &iam.GetAccountPasswordPolicyOutput{PasswordPolicy: &types.PasswordPolicy{RequireNumbers: true}}

	assert.True(t, PolicyRequiresNumber(policy))
}

func TestPolicyRequiresSymbol_NotRequired(t *testing.T) {
	policy := &iam.GetAccountPasswordPolicyOutput{PasswordPolicy: &types.PasswordPolicy{RequireNumbers: false}}

	assert.False(t, PolicyRequiresNumber(policy))
}
