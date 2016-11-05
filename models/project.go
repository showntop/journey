package models

import (
	"fmt"
)

var _ = fmt.Println

type Project struct {
	Base
	CategoryId  int    `json:"category_id"`
	Name        string `json:"name"`
	Version     string
	Description string
	Size        int
	LogoURL     string
	Dlink       string
	Assets      []*ProjectAsset
	TagArray    []*Tag `json:"-" pg:",many2many:project_tags"`
	Tags        []string
}

type ProjectAsset struct {
	Base
	ProjectId int
	URL       string
}

type Tag struct {
	Base
	Name string
}

type ProjectTag struct {
	Base
	ProjectId int
	TagId     int
}
