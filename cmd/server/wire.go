package main

import (
	user "rest-api/user"

	"github.com/google/wire"
)

func InitializeUserHandler() *user.UserHandler {
	wire.Build(
		user.NewUserService,
		user.NewUserHandler,
	)
	return &user.UserHandler{}
}
