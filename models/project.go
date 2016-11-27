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

type ProjectList []*struct {
	Id      int64    `json:"id"`
	Name    string   `json:"name"`
	Size    int64    `json:"size"`
	Version string   `json:"version"`
	LogoURL string   `json:"logo_url"`
	Dlink   string   `json:"dlink"`
	Tags    []string `json:tags`
}
