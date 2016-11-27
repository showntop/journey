package stores

import (
	"github.com/showntop/journey/models"
)

type PostStore struct {
	*Store
}

func (p *PostStore) FindByProject(projectId int, offset, limit int) ([]*models.Post, error) {
	posts := []*models.Post{}
	err := p.Master.Model(&posts).Column("post.*", "Comments").Offset(offset).Limit(limit).Select()
	//pos
	return posts, err
}

func (p *PostStore) Create(post *models.Post) error {

	err := p.Master.Insert(post)
	return err
}

func (p *PostStore) CreateLike(like *models.PostLike) error {

	err := p.Master.Insert(like)
	return err
}

func (p *PostStore) CreateComment(comment *models.PostComment) error {

	err := p.Master.Insert(comment)
	return err
}
