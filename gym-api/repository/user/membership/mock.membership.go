package membership

import (
	"gym-membership/business/membership"

	"github.com/stretchr/testify/mock"
)

type MockMembershipRepository struct {
	data []*membership.Membership
	Mock mock.Mock
}

func (repository *MockMembershipRepository) GetAllMember() (memberships []*membership.Membership, err error) {
	return repository.data, nil
}
