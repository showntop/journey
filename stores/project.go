package stores

import (
	"github.com/showntop/journey/models"
)

type ProjectStore struct {
	*Store
}

func (p *ProjectStore) FindWith(querySQL string, offset, limit int) (models.ProjectList, error) {
	var projects models.ProjectList
	err := p.Master.Model(&models.Project{}).Column("Id", "Name", "Size", "Version", "LogoURL", "Dlink", "Tags").Where(querySQL).Select(&projects)

	return projects, err
}

// func (p *ProjectStore) FindWith(querySQL string, offset, limit int) (models.ProjectList, error) {
// 	var projects models.ProjectList
// 	_, err := p.Master.Query(&projects, "SELECT projects.id, projects.name, projects.version, projects.size, projects.logo_url, projects.dlink FROM projects WHERE "+querySQL)

// 	idsStr := ""
// 	for _, p := range projects {
// 		idsStr += (strconv.FormatInt(p.Id, 10) + ",")
// 	}
// 	var ptags map[string]string
// 	p.Master.Query(&ptags, "SELECT pt.project_id, t.name FROM tags t JOIN project_tags pt ON t.id = pt.tag_id WHERE pt.project_id in (?)", idsStr[:len(idsStr)-1])
// 	fmt.Printf("ptags: %v", ptags)
// 	return projects, err
// }

func (p *ProjectStore) FindWithTag(tagId int64, offset, limit int) ([]*models.Project, error) {
	var projects []*models.Project
	err := p.Master.Model(&projects).Column("project.*", "TagArray").Where("tag_id = ?", tagId).Select()
	return projects, err
}

func (p *ProjectStore) FindWithKey(key string, offset, limit int) ([]*models.Project, error) {
	var projects []*models.Project
	err := p.Master.Model(&projects).Column("project.*", "TagArray").Where("name like ? or description like ?", key).Select()
	return projects, err
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
