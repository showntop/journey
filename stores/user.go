package stores

import (
	"github.com/showntop/journey/models"
)

type UserStore struct {
	*Store
}

func (u *UserStore) Create(user *models.User) error {

	err := u.Master.Create(user).Error
	return err
}

func (u *UserStore) FindBy(filedV string) (*models.User, error) {
	user := &models.User{}
	err := u.Master.Where("username = ?", filedV).First(user)
	if err.RecordNotFound() {
		return nil, nil
	}
	return user, err.Error
}
