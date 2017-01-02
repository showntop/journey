package stores

import (
	"fmt"

	"github.com/showntop/journey/models"
)

type ProjectStore struct {
	*Store
}

func (p *ProjectStore) FindWithTag(tagId int64, offset, limit int) ([]*models.Project, error) {
	var projects []*models.Project
	err := p.Master.Model(&projects).Related(&[]*models.Tag{}).Where("tag_id = ?", tagId).Error
	return projects, err
}

func (p *ProjectStore) FindWithKey(key string, offset, limit int) ([]*models.Project, error) {
	key = "%" + key + "%"
	var projects []*models.Project
	err := p.Master.Preload("Tags").Select("projects.id, projects.name, projects.version, projects.size, projects.logo_url, projects.dlink").Where("name like ? or description like ?", key, key).Find(&projects).Error
	return projects, err
}

func (p *ProjectStore) FindAll(offset, limit int) ([]*models.Project, error) {
	projects := []*models.Project{}
	err := p.Master.Preload("Tags").Select("projects.id, projects.name, projects.version, projects.size, projects.logo_url, projects.dlink").Offset(offset).Limit(limit).Find(&projects).Error
	fmt.Println("")
	// fmt.Printf("%v", rows)
	return projects, err
}

func (p *ProjectStore) FindAllByCategory(cid int64, offset, limit int) ([]*models.Project, error) {
	projects := []*models.Project{}
	err := p.Master.Preload("Tags").Where("category_id = ?", cid).Select("projects.id, projects.name, projects.version, projects.size, projects.logo_url, projects.dlink").Offset(offset).Limit(limit).Find(&projects).Error

	return projects, err
}

func (u *ProjectStore) Find(id int64) (*models.Project, error) {
	project := &models.Project{}
	err := u.Master.First(project, id).Error

	return project, err
}
