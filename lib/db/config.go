package db

import "gopkg.in/mgo.v2"

type Cfg struct {
	url      string
	username string
	password string
}

var Configs = map[string]Cfg{

	"development": Cfg{
		url:      "localhost:27017",
		username: "",
		password: ""},

	"production": Cfg{
		url:      "",
		username: "",
		password: ""}}
