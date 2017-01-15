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
			URL:    "http://oizfueie4.bkt.clouddn.com/assets/boots/splashPage.png",
			Action: "com.hotkuy.app.android.activity.WebViewActivity",
			Params: "urlStr:htttp://www.baidu.com",
		},
	}

	bannerAssets = [...]*models.Asset{
		&models.Asset{
			URL:    "http://oizfueie4.bkt.clouddn.com/assets/banners/banner_xiaoxiaole.jpg",
			Action: "com.hotkuy.app.android.activity.WebViewActivity",
			Params: "urlStr:htttp://www.baidu.com",
		},
		&models.Asset{
			URL:    "http://oizfueie4.bkt.clouddn.com/assets/banners/banner_dnfphoto.png",
			Action: "com.hotkuy.app.android.activity.WebViewActivity",
			Params: "urlStr:htttp://www.baidu.com",
		},
		&models.Asset{
			URL:    "http://oizfueie4.bkt.clouddn.com/banner_wangzherongyao.jpg",
			Action: "com.hotkuy.app.android.activity.WebViewActivity",
			Params: "urlStr:htttp://www.baidu.com",
		},
		&models.Asset{
			URL:    "http://oizfueie4.bkt.clouddn.com/assets/banners/banner_yinyangshi.png",
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
