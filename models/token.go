package models

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/groupcache/lru"
	"github.com/satori/go.uuid"
)

const (
	KEY_NAMESPACE = "sun-token:"
)

type Store struct {
	Cache *lru.Cache
}

var StoreM *Store = &Store{}

func init() {
	StoreM.Cache = lru.New(100)
	log.WithField("server", "starting").Info("init storage success...")
}

type Token struct {
}

func CreateTokenFor(user *User) string {
	token := uuid.NewV4().String()
	StoreM.Cache.Add(fmt.Sprintf("%s%s", KEY_NAMESPACE, token), *user)
	return token
}

func RetriveUserFromToken(token string) (*User, error) {
	value, ok := StoreM.Cache.Get(fmt.Sprintf("%s%s", KEY_NAMESPACE, token))
	if !ok {
		return nil, fmt.Errorf("unrecgonized toke %s", token)
	}
	user := value.(User)
	return &user, nil
}
