package role

import "github.com/netbrain/gonk/app/user"
import "labix.org/v2/mgo/bson"

type Role struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Name  string        ",omitempty"
	Users []user.User   ",omitempty"
}
