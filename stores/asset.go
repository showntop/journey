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
			Action: "com.hotkuy.app.android.activity.WebViewActivity",
			Params: "urlStr:htttp://www.baidu.com",
		},
	}

	bannerAssets = [...]*models.Asset{
		&models.Asset{
			URL:    "http://oginbhz1c.bkt.clouddn.com/5882b2b7d0a20cf482c772bf73094b36acaf997f.jpg",
			Action: "com.hotkuy.app.android.activity.WebViewActivity",
			Params: "urlStr:htttp://www.baidu.com",
		},
		&models.Asset{
			URL:    "http://oginbhz1c.bkt.clouddn.com/5882b2b7d0a20cf482c772bf73094b36acaf997f.jpg",
			Action: "com.hotkuy.app.android.activity.WebViewActivity",
			Params: "urlStr:htttp://www.baidu.com",
		},
		&models.Asset{
			URL:    "http://oginbhz1c.bkt.clouddn.com/5882b2b7d0a20cf482c772bf73094b36acaf997f.jpg",
			Action: "com.hotkuy.app.android.activity.WebViewActivity",
			Params: "urlStr:htttp://www.baidu.com",
		},
		&models.Asset{
			URL:    "http://oginbhz1c.bkt.clouddn.com/5882b2b7d0a20cf482c772bf73094b36acaf997f.jpg",
			Action: "com.hotkuy.app.android.activity.WebViewActivity",
			Params: "urlStr:htttp://www.baidu.com",
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
