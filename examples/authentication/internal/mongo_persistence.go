package internal

import (
	"labix.org/v2/mgo"
)

type MongoPersistence struct {
	InitHandlers []func(*mgo.Database)
	DialUrl      string
	Database     string
	session      *mgo.Session
}

func (p *MongoPersistence) Init() {
	var err error
	p.session, err = mgo.Dial(p.DialUrl)
	if err != nil {
		panic(err)
	}
}

func (p MongoPersistence) InitOnRequest() interface{} {
	session := p.session.Clone()
	db := session.DB(p.Database)
	for _, handler := range p.InitHandlers {
		handler(db)
	}
	return session
}

func (p MongoPersistence) CloseAfterRequest(s interface{}) {
	session := s.(*mgo.Session)
	session.Close()
}
