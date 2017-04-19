package stores

import (
	"github.com/showntop/journey/models"
)

type SubjectStore struct {
	*Store
}

func (c *SubjectStore) FindAll(offset, limit int) ([]*models.Subject, error) {
	subjects := []*models.Subject{}
	err := c.Master.Find(&subjects).Error

	return subjects, err
}

func (u *SubjectStore) Find(id int64) (*models.Subject, error) {
	subject := &models.Subject{}
	err := u.Master.First(subject, id).Error
	return subject, err
}
