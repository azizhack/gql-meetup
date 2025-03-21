//go:generate go run github.com/99designs/gqlgen --v

package graph

import "github.com/azizhack/gql-meetup/postgres"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	MeetupRepo *postgres.MeetupsRepo
	UserRepo   *postgres.UsersRepo
}
