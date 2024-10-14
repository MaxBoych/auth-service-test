package userInMemoryRepository

import (
	"auth-service-test/internal/domain"
	"context"
)

func (r *UserRepository) UpdateOne(_ context.Context, curr domain.User, data domain.UpdatedUserData) error {
	curr = getUpdatedUser(curr, data)
	r.users.Store(curr.Username, curr)

	return nil
}

func getUpdatedUser(curr domain.User, data domain.UpdatedUserData) domain.User {
	if data.Role != nil {
		curr.Role = *data.Role
	}
	if data.Firstname != nil {
		curr.Firstname = *data.Firstname
	}
	if data.Lastname != nil {
		curr.Lastname = *data.Lastname
	}
	if data.Email != nil {
		curr.Email = *data.Email
	}
	if data.PhoneNumber != nil {
		curr.PhoneNumber = *data.PhoneNumber
	}

	return curr
}
