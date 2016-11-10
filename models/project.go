package models

import (
	"fmt"
)

var _ = fmt.Println

type Project struct {
	Base
	CategoryId  int    `json:"category_id"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Size        int    `json:"size"`
	LogoURL     string `json:"logo_url"`
	Dlink       string `json:"download_url"`
	Assets      []*ProjectAsset
	TagArray    []*Tag   `json:"-" pg:",many2many:project_tags"`
	Tags        []string `json:"tags"`
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
