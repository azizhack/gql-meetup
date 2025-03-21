package postgres

import (
	"github.com/azizhack/gql-meetup/models"
	"github.com/go-pg/pg/v10"
)

type UsersRepo struct {
	DB *pg.DB
}

func (r *UsersRepo) GetUserById(id string) (*models.User, error) {
	var user *models.User
	err := r.DB.Model(&user).Where("id = ?", id).First()
	return user, err
}
func (r *UsersRepo) GetUsers() ([]*models.User, error) {
	var user []*models.User
	err := r.DB.Model(&user).Select()
	return user, err
}
