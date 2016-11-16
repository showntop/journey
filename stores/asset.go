package stores

import (
	// "gopkg.in/pg.v5"

	"github.com/showntop/journey/models"
)

type AssetStore struct {
	*Store
}

var (
	staticAssets = [...]string{"", ""}
)

func (u *AssetStore) FindAllWith(ftype string) ([]*models.Asset, error) {
	var assets []*models.Asset
	for _, url := range staticAssets {
		asset := &models.Asset{
			URL: url,
		}
		assets = append(assets, asset)
	}
	return assets, nil
}
