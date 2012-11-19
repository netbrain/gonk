package user

import (
	"code.google.com/p/go.crypto/bcrypt"
	"labix.org/v2/mgo/bson"
)

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Email    string        ",omitempty"
	Password []byte        ",omitempty"
}

//encrypts and sets the password on the user
//returns nil on success and error on fail
func (u *User) SetPassword(password string) error {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = encrypted
	return nil
}

//hashes password and compares it against stored password hash.
//returns nil on success and error on fail
func (u *User) PasswordEquals(password string) error {
	return bcrypt.CompareHashAndPassword(u.Password, []byte(password))
}

func (u *User) Equals(other User) bool {
	return u.ID == other.ID
}
