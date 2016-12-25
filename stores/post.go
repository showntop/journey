package stores

import (
	"github.com/showntop/journey/models"
)

type PostStore struct {
	*Store
}

func (p *PostStore) FindByProject(projectId int, offset, limit int) ([]*models.Post, error) {
	posts := []*models.Post{}
	err := p.Master.Where("project_id = ?", projectId).Find(&posts).Error
	//pos
	return posts, err
}

func (p *PostStore) Create(post *models.Post) error {

	err := p.Master.Create(post).Error
	return err
}

func (p *PostStore) CreateLike(like *models.PostLike) error {

	err := p.Master.Create(like).Error
	return err
}

func (p *PostStore) CreateComment(comment *models.PostComment) error {

	err := p.Master.Create(comment).Error
	return err
}
