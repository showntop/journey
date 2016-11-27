package stores

import (
	// "gopkg.in/pg.v5"

	"github.com/showntop/journey/models"
)

type AssetStore struct {
	*Store
}

var (
	bootAssets = [...]*models.Asset{
		&models.Asset{
			URL:    "http://oginbhz1c.bkt.clouddn.com/5882b2b7d0a20cf482c772bf73094b36acaf997f.jpg",
			Action: "xxxx",
		},
	}

	bannerAssets = [...]*models.Asset{
		&models.Asset{
			URL:    "http://oginbhz1c.bkt.clouddn.com/5882b2b7d0a20cf482c772bf73094b36acaf997f.jpg",
			Action: "xxxx",
		},
		&models.Asset{
			URL:    "http://oginbhz1c.bkt.clouddn.com/5882b2b7d0a20cf482c772bf73094b36acaf997f.jpg",
			Action: "xxxx",
		},
		&models.Asset{
			URL:    "http://oginbhz1c.bkt.clouddn.com/5882b2b7d0a20cf482c772bf73094b36acaf997f.jpg",
			Action: "xxxx",
		},
		&models.Asset{
			URL:    "http://oginbhz1c.bkt.clouddn.com/5882b2b7d0a20cf482c772bf73094b36acaf997f.jpg",
			Action: "xxxx",
		},
	}
)

func (u *AssetStore) FindAllWith(ftype string) ([]*models.Asset, error) {
	if ftype == "boot" {
		return bootAssets[:], nil
	} else {
		return bannerAssets[:], nil
	}
}
