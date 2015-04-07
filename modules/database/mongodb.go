package database

import (
	"fmt"
	mgo "gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
	"log"
	_ "time"
)

var (
	session *mgo.Session
)

func StartMongoDb() {
	current_session, err := mgo.Dial("107.167.180.219:27017")
	if err != nil {
		log.Println("err connecting to mongodb!")
		log.Println("error: ", err)
		return
	}
	fmt.Println("connected to mongodb!")
	session = current_session
}
