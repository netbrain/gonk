package user

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var collection *mgo.Collection

func Init(db *mgo.Database) {
	collection = db.C("user")
}

func All() []User {
	var result []User = []User{}
	err := collection.Find(nil).All(&result)
	if err != nil {
		panic(err)
	}
	return result
}

func Get(id string) User {
	result := User{}
	if err := collection.FindId(bson.ObjectIdHex(id)).One(&result); err != nil {
		panic(err)
	}
	return result
}

func save(user User) {
	if err := collection.Insert(user); err != nil {
		panic(err)
	}
}

func update(user User) {
	if err := collection.UpdateId(user.ID, user); err != nil {
		panic(err)
	}
}

func SaveOrUpdate(user User) {
	if user.ID == "" {
		save(user)
	} else {
		update(user)
	}
}

func Remove(id string) {
	if err := collection.RemoveId(bson.ObjectIdHex(id)); err != nil {
		panic(err)
	}
}
