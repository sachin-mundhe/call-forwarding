package database

import (
	"fmt"
	"sync"

	harperdb "github.com/HarperDB-Add-Ons/sdk-go"
	"github.com/call-forwarding/config"
)

type harperDB struct {
	Client *harperdb.Client
}

var (
	once     sync.Once
	dbClient *harperDB
)

func NewHarperDatabase(conf *config.Config) Database {
	once.Do(func() {
		client := harperdb.NewClient(fmt.Sprintf("%s:%d", conf.Db.Host, conf.Db.Port), conf.Db.User, conf.Db.Password)
		dbClient = &harperDB{client}
	})
	return dbClient
}

func (p *harperDB) GetDb() *harperdb.Client {
	return dbClient.Client
}
