package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	Collection *mgo.Collection
)

type Object struct {
	PublishID string `json:"id"`
	PublishDate string `json:"publish_date"`
	Title string `json:"title"`
	Detail []string `json:"detail"`
}

type Objects struct {
	News []Object `json:"news"`
}

func init() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	//defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	Collection = session.DB("swarm").C("sinanews")
}

func GetOne(PublishID string) (ob Object, err error) {
	err = Collection.Find(bson.M{"publish_id": PublishID}).One(&ob)
	return
}

func GetAll() (result Objects, err error){
	obs := []Object{}
	err = Collection.Find(bson.M{}).All(&obs)
	result.News = obs
	return
}
