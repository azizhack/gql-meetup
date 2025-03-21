package postgres

import (
	"github.com/azizhack/gql-meetup/entities"
	"github.com/go-pg/pg/v10"
)

type UsersRepo struct {
	DB *pg.DB
}

func (r *UsersRepo) GetUserById(id string) (*entities.User, error) {
	var user *entities.User
	err := r.DB.Model(&user).Where("id", id).First()
	return user, err
}
func (r *UsersRepo) GetUsers() ([]*entities.User, error) {
	var user []*entities.User
	err := r.DB.Model(&user).Select()
	return user, err
}
