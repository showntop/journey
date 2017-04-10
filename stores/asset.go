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
			URL:    "http://ohan0t6mr.bkt.clouddn.com/bar/jpg/huochairen.jpg",
			Action: "com.hotkuy.app.android.activity.IntroduceActivity",
			Params: "app_id:1",
		},
	}

	bannerAssets = [...]*models.Asset{
		&models.Asset{
			URL:    "http://ohan0t6mr.bkt.clouddn.com/bar/jpg/huochairen.jpg",
			Action: "com.hotkuy.app.android.activity.IntroduceActivity",
			Params: "app_id:1",
		},

		&models.Asset{
			URL:    "http://ohan0t6mr.bkt.clouddn.com/bar/jpg/wangzherongyao.jpg",
			Action: "com.hotkuy.app.android.activity.IntroduceActivity",
			Params: "app_id:200002",
		},
		&models.Asset{
			URL:    "http://ohan0t6mr.bkt.clouddn.com/bar/jpg/yinyangshi.jpg",
			Action: "com.hotkuy.app.android.activity.IntroduceActivity",
			Params: "app_id:200001",
		},
		&models.Asset{
			URL:    "http://oizfueie4.bkt.clouddn.com/banner_wangzherongyao.jpg",
			Action: "com.hotkuy.app.android.activity.IntroduceActivity",
			Params: "app_id:10004",
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
