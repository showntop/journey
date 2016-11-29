package stores

import (
	"fmt"
	"github.com/showntop/journey/models"
)

type ProjectStore struct {
	*Store
}

// func (p *ProjectStore) FindWith(querySQL string, offset, limit int) (interface{}, error) {
// 	var projects []*models.Project
// 	err := p.Master.Offset(offset).Limit(limit).Model(&[]*models.Project{}).Column("project.id", "TagArray")

// 	return projects, err
// }

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
	err := p.Master.Model(&projects).Related(&[]*models.Tag{}).Where("tag_id = ?", tagId).Error
	return projects, err
}

func (p *ProjectStore) FindWithKey(key string, offset, limit int) ([]*models.Project, error) {
	var projects []*models.Project
	err := p.Master.Model(&projects).Related([]*models.Tag{}).Where("name like ? or description like ?", key).Error
	return projects, err
}

func (p *ProjectStore) FindAll(offset, limit int) ([]*models.Project, error) {
	projects := []*models.Project{}
	err := p.Master.Preload("Tags").Select("projects.id, projects.name, projects.version, projects.size, projects.logo_url, projects.download_url").Offset(offset).Limit(limit).Find(&projects).Error
	// err := p.Master.Select("projects.id, projects.name, projects.version, projects.size, projects.logo_url, projects.download_url, tags.name").Table("projects").Offset(offset).Limit(limit).Joins("LEFT JOIN project_tags ON projects.id = project_tags.project_id").Joins("LEFT JOIN tags ON project_tags.tag_id = tags.id").Scan(&projects).Error
	///pos
	// for _, p := range projects {
	// 	for _, tag := range p.TagArray {
	// 		p.Tags = append(p.Tags, tag.Name)
	// 	}
	// }
	//pos
	fmt.Println("")
	// fmt.Printf("%v", rows)
	return projects, err
}

func (u *ProjectStore) FindAllByCategory(cid int, offset, limit int) ([]*models.Project, error) {
	projects := []*models.Project{}
	err := u.Master.Model(&projects).Related(&[]*models.Tag{}, "tags").Where("category_id = ?", cid).Offset(offset).Limit(limit).Error

	///pos
	// for _, p := range projects {
	// 	for _, tag := range p.TagArray {
	// 		p.Tags = append(p.Tags, tag.Name)
	// 	}
	// }
	//pos

	return projects, err
}

func (u *ProjectStore) Find(id int64) (*models.Project, error) {
	project := &models.Project{}
	err := u.Master.First(project, id).Error

	return project, err
}
