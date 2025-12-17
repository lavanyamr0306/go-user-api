package service

import (
	"context"
	"time"

	db "github.com/lavanyamr0306/go-user-api/internal/db/sqlc"
	"github.com/lavanyamr0306/go-user-api/internal/models"
)

type UserService struct {
	store *db.Queries
}

func NewUserService(store *db.Queries) *UserService {
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
func toModel(user db.User) models.User {
	return models.User{
		ID:   int(user.ID),
		Name: user.Name,
		DOB:  user.Dob,
		Age:  calculateAge(user.Dob),
	}
}

// List Users
func (s *UserService) ListUsers(ctx context.Context) ([]models.User, error) {
	users, err := s.store.ListUsers(ctx)
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
func (s *UserService) GetUser(ctx context.Context, id int32) (models.User, error) {
	user, err := s.store.GetUserByID(ctx, id)
	if err != nil {
		return models.User{}, err
	}
	return toModel(user), nil
}

// Create User
func (s *UserService) CreateUser(ctx context.Context, name string, dob time.Time) error {
	params := db.CreateUserParams{
		Name: name,
		Dob:  dob,
	}
	return s.store.CreateUser(ctx, params)
}

// Update User
func (s *UserService) UpdateUser(ctx context.Context, id int32, name string, dob time.Time) error {
	params := db.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dob,
	}
	return s.store.UpdateUser(ctx, params)
}

// Delete User
func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
	return s.store.DeleteUser(ctx, id)
}
