package role

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var collection *mgo.Collection

func Init(db *mgo.Database) {
	collection = db.C("role")
}

func All() []Role {
	var result []Role = []Role{}
	err := collection.Find(nil).All(&result)
	if err != nil {
		panic(err)
	}
	return result
}

func Get(id string) Role {
	result := Role{}
	if err := collection.FindId(bson.ObjectIdHex(id)).One(&result); err != nil {
		panic(err)
	}
	return result
}

func save(role Role) {
	if err := collection.Insert(role); err != nil {
		panic(err)
	}
}

func update(role Role) {
	if err := collection.UpdateId(role.ID, role); err != nil {
		panic(err)
	}
}

func SaveOrUpdate(role Role) {
	if role.ID == "" {
		save(role)
	} else {
		update(role)
	}
}

func Remove(id string) {
	if err := collection.RemoveId(bson.ObjectIdHex(id)); err != nil {
		panic(err)
	}
}
