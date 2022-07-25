package password_symbol

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/stretchr/testify/assert"
)

func TestPolicyRequiresSymbol_Required(t *testing.T) {
	policy := &iam.GetAccountPasswordPolicyOutput{PasswordPolicy: &types.PasswordPolicy{RequireSymbols: true}}

	assert.True(t, PolicyRequiresSymbol(policy))
}

func TestPolicyRequiresSymbol_NotRequired(t *testing.T) {
	policy := &iam.GetAccountPasswordPolicyOutput{PasswordPolicy: &types.PasswordPolicy{RequireSymbols: false}}

	assert.False(t, PolicyRequiresSymbol(policy))
}
