package stores

import (
	// "fmt"
	"crypto/tls"
	log2 "log"
	"os"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/pg.v5"

	. "github.com/showntop/journey/config"
)

type Store struct {
	Master   *pg.DB   //*sqlx.DB
	Replicas []*pg.DB //*sqlx.DB
	User     *UserStore
	Project  *ProjectStore
	Category *CategoryStore
	Asset    *AssetStore
	Post     *PostStore
}

var (
	StoreM *Store
)

func SetupStorage() {
	log.WithField("server", "starting").Info("init storage...")
	StoreM = &Store{}
	StoreM.Master = setupDB(Config.Database)

	StoreM.User = &UserStore{StoreM}
	StoreM.Project = &ProjectStore{StoreM}
	StoreM.Category = &CategoryStore{StoreM}
	StoreM.Asset = &AssetStore{StoreM}
	StoreM.Post = &PostStore{StoreM}
	log.Debugf("%+v", StoreM)

}

func setupDB(config map[string]interface{}) *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     config["addr"].(string),
		User:     config["user"].(string),
		Password: config["password"].(string),
		Database: config["dbname"].(string),
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		// Whether to use secure TCP/IP connections (TLS).

	})
	pg.SetQueryLogger(log2.New(os.Stdout, "postgres", 1))
	pg.SetLogger(log2.New(os.Stdout, "postgres", 1))
	log.WithField("server", "starting").Info("init db success...")
	log.Debug(db)
	return db
}
