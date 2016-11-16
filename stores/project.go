package stores

import (
	"github.com/showntop/journey/models"
)

type ProjectStore struct {
	*Store
}

func (p *ProjectStore) FindWith(tagId int64, key string, offset, limit int) ([]*models.Project, error) {
	p.Master.Model(&projects).Where(where, ...)
}

func (u *ProjectStore) FindAll(offset, limit int) ([]*models.Project, error) {
	projects := []*models.Project{}
	err := u.Master.Model(&projects).Column("project.*", "TagArray").Offset(offset).Limit(limit).Select()

	///pos
	for _, p := range projects {
		for _, tag := range p.TagArray {
			p.Tags = append(p.Tags, tag.Name)
		}
	}
	//pos
	return projects, err
}

func (u *ProjectStore) FindAllByCategory(cid int, offset, limit int) ([]*models.Project, error) {
	projects := []*models.Project{}
	err := u.Master.Model(&projects).Column("project.*", "TagArray").Where("category_id = ?", cid).Offset(offset).Limit(limit).Select()

	///pos
	for _, p := range projects {
		for _, tag := range p.TagArray {
			p.Tags = append(p.Tags, tag.Name)
		}
	}
	//pos

	return projects, err
}

func (u *ProjectStore) Find(id int) (*models.Project, error) {
	project := &models.Project{}
	err := u.Master.Model(project).Column("project.*", "Assets").Where("id = ?", id).First()

	return project, err
}
