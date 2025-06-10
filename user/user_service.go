package user

import "fmt"

type UserService struct {
	// potentially DB connection, etc.
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUser(id string) string {
	// Just a dummy example
	return fmt.Sprintf("User: %s", id)
}
