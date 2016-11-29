package models

import (
	"fmt"
)

var _ = fmt.Println

type Project struct {
	Base
	CategoryId  int         `json:"category_id"`
	Name        string      `json:"name"`
	Version     string      `json:"version"`
	Description string      `json:"description"`
	Size        int         `json:"size"`
	LogoURL     string      `json:"logo_url"`
	Dlink       string      `json:"download_url"`
	Assets      StringSlice `json:"assets,omitempty"`
	Tags        []Tag       `json:"tags" gorm:"many2many:project_tags;"`
	// Tags        []string `json:"tags"`
}

type Tag struct {
	Base
	Name string
	// Projects []Project `gorm:"many2many:project_tags;"`
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
