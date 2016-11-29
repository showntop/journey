package stores

import (
	"github.com/showntop/journey/models"
)

type CategoryStore struct {
	*Store
}

func (c *CategoryStore) FindAll() ([]*models.Category, error) {
	categories := []*models.Category{}
	err := c.Master.Find(&categories).Error

	return categories, err
}
