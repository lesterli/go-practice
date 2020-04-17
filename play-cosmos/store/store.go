package store

import (
	"strings"
	"time"

	"github.com/lesterli/go-practice/play-cosmos/config/db"
	"github.com/lesterli/go-practice/play-cosmos/logger"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session
var log = logger.GetLogger("store")

func Start() {
	db.Init()
	addrs := strings.Split(db.Addrs, ",")
	dialInfo := &mgo.DialInfo{
		Addrs:     addrs,
		Database:  db.Database,
		Username:  db.User,
		Password:  db.Passwd,
		Direct:    true,
		Timeout:   time.Second * 10,
		PoolLimit: 4096,
	}

	var err error
	session, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Error("Mongodb dial error with" + err.Error())
	}
	session.SetMode(mgo.Primary, true)
}

func Stop() {
	log.Info("release resource :mongoDb")
	session.Close()
}

func getSession() *mgo.Session {
	// max session num is 4096
	return session.Clone()
}

// ExecCollection - get collection object
func ExecCollection(collection string, s func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB(db.Database).C(collection)
	return s(c)
}

func Find(collection string, query interface{}) *mgo.Query {
	session := getSession()
	defer session.Close()
	c := session.DB(db.Database).C(collection)
	return c.Find(query)
}
