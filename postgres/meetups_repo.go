package postgres

import (
	"github.com/azizhack/gql-meetup/models"
	"github.com/go-pg/pg/v10"
)

type MeetupsRepo struct {
	DB *pg.DB
}

func (r *MeetupsRepo) CreateMeetup(meetup *models.Meetup) (*models.Meetup, error) {
	_, err := r.DB.Model(meetup).Returning("*").Insert()
	return meetup, err
}

func (r *MeetupsRepo) GetMeetups() ([]*models.Meetup, error) {
	var meetups []*models.Meetup
	err := r.DB.Model(&meetups).Select()
	return meetups, err
}

func (r *MeetupsRepo) GetMeetupsByUserID(userId string) ([]*models.Meetup, error) {
	var meetups []*models.Meetup
	err := r.DB.Model(&meetups).Where("user_id = ?", userId).Select()
	return meetups, err
}
