package postgres

import (
	"github.com/azizhack/gql-meetup/entities"
	"github.com/go-pg/pg/v10"
)

type MeetupsRepo struct {
	DB *pg.DB
}

func (r *MeetupsRepo) CreateMeetup(meetup *entities.Meetup) (*models.Meetup, error) {
	_, err := r.DB.Model(meetup).Returning("*").Insert()
	return meetup, err
}

func (r *MeetupsRepo) GetMeetups() ([]*entities.Meetup, error) {
	var meetups []*entities.Meetup
	err := r.DB.Model(&meetups).Select()
	return meetups, err
}

func (r *MeetupsRepo) GetMeetupsByUserID(userId string) ([]*entities.Meetup, error) {
	var meetups []*entities.Meetup
	err := r.DB.Model(&meetups).Where("user_id", userId).Select()
	return meetups, err
}
