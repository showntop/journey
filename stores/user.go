package stores

import (
	"gopkg.in/pg.v5"

	"github.com/showntop/journey/models"
)

type UserStore struct {
	*Store
}

func (u *UserStore) Create(user *models.User) error {

	err := u.Master.Insert(user)
	return err
}

func (u *UserStore) FindBy(filedV string) (*models.User, error) {
	user := &models.User{}
	err := u.Master.Model(user).Where("username = ?", filedV).First()
	if pg.ErrNoRows == err {
		return nil, nil
	}
	return user, err
}
