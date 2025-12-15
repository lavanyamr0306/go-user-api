package service

import (
	"time"

	"go-user-api/db/sqlc"
	"go-user-api/internal/models"
)

type UserService struct {
	store *sqlc.Queries
}

func NewUserService(store *sqlc.Queries) *UserService {
	return &UserService{store: store}
}

// Calculate age
func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}

// Convert SQLC user to models.User with age
func toModel(user sqlc.User) models.User {
	return models.User{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob,
		Age:  calculateAge(user.Dob),
	}
}

// List Users
func (s *UserService) ListUsers() ([]models.User, error) {
	users, err := s.store.ListUsers()
	if err != nil {
		return nil, err
	}
	result := make([]models.User, len(users))
	for i, u := range users {
		result[i] = toModel(u)
	}
	return result, nil
}

// Get User by ID
func (s *UserService) GetUser(id int) (models.User, error) {
	user, err := s.store.GetUserByID(int32(id))
	if err != nil {
		return models.User{}, err
	}
	return toModel(user), nil
}

// Create User
func (s *UserService) CreateUser(name string, dob time.Time) error {
	return s.store.CreateUser(name, dob)
}

// Update User
func (s *UserService) UpdateUser(id int, name string, dob time.Time) error {
	return s.store.UpdateUser(name, dob, int32(id))
}

// Delete User
func (s *UserService) DeleteUser(id int) error {
	return s.store.DeleteUser(int32(id))
}
