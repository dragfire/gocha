package db

import "gopkg.in/mgo.v2"

func GetConnection(env string) (db *mgo.Database, err error) {
	cfg := Configs[env]
	session, err := mgo.Dial(cfg.url)
	if err != nil {
		return nil, err
	}
	return session.DB("gocha"), nil
}
