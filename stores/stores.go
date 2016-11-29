package stores

import (
	// "fmt"
	// "crypto/tls"
	// log2 "log"
	// "os"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	. "github.com/showntop/journey/config"
)

type Store struct {
	Master   *gorm.DB   //*sqlx.DB
	Replicas []*gorm.DB //*sqlx.DB
	User     *UserStore
	Project  *ProjectStore
	Category *CategoryStore
	Asset    *AssetStore
	Post     *PostStore
	Subject  *SubjectStore
}

var (
	StoreM *Store
)

func SetupStorage() {
	log.WithField("server", "starting").Info("init storage...")
	StoreM = &Store{}
	StoreM.Master = setupDB()

	StoreM.User = &UserStore{StoreM}
	StoreM.Project = &ProjectStore{StoreM}
	StoreM.Category = &CategoryStore{StoreM}
	StoreM.Asset = &AssetStore{StoreM}
	StoreM.Post = &PostStore{StoreM}

	StoreM.Subject = &SubjectStore{StoreM}
}

func setupDB() *gorm.DB {
	db, err := gorm.Open("postgres", Config.Dbstr)
	db.LogMode(true)
	if err != nil {
		log.Errorln(err)
	}
	return db
}
