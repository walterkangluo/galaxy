package policy

import (
	"github.com/DSiSc/craft/types"
	"github.com/DSiSc/validator/tools/account"
	"github.com/stretchr/testify/assert"
	"testing"
)

func mock_NewSoloPolicy() *SoloPolicy {
	policy, _ := NewSoloPolicy()
	return policy
}

var MockAccount = account.Account{
	Address: types.Address{
		0x33, 0x3c, 0x33, 0x10, 0x82, 0x4b, 0x7c, 0x68, 0x51, 0x33,
		0xf2, 0xbe, 0xdb, 0x2c, 0xa4, 0xb8, 0xb4, 0xdf, 0x63, 0x3d,
	},
}

func Test_NewSoloPolicy(t *testing.T) {
	asserts := assert.New(t)
	policy, err := NewSoloPolicy()
	asserts.NotNil(policy)
	asserts.Nil(err)
	asserts.Equal(SOLO_POLICY, policy.name, "they should not be equal")
}

func Test_PolicyName(t *testing.T) {
	asserts := assert.New(t)
	policy := mock_NewSoloPolicy()
	policyName := policy.PolicyName()
	asserts.Equal(SOLO_POLICY, policyName, "they should not be equal")
}

func Test_GetParticipates(t *testing.T) {
	asserts := assert.New(t)
	policy := mock_NewSoloPolicy()
	address, err := policy.GetParticipates()
	asserts.NotNil(address)
	asserts.Nil(err)
	asserts.Equal(1, len(address), "they should not be equal")
}

func Test_getMembers(t *testing.T) {
	asserts := assert.New(t)
	policy := mock_NewSoloPolicy()
	members := policy.getMembers()
	asserts.Equal(MockAccount, members, "they should not be equal")
}
