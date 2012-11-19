package main

import (
	. "github.com/netbrain/gonk"
	"github.com/netbrain/gonk/examples/authentication/app/role"
	"github.com/netbrain/gonk/examples/authentication/app/user"
	. "github.com/netbrain/gonk/examples/authentication/internal"
	"labix.org/v2/mgo"
)

func main() {
	c := Config{
		//HTTPConfig
		"ServerAddr":  "localhost:8080",
		"StrictSlash": true,
		//HTTP Handlers
		"InitContextHandler":  nil,
		"FailureHandler":      nil,
		"RenderViewHandler":   nil,
		"RenderStatusHandler": nil,
		"DefaultBefore":       nil,
		"DefaultAfter":        nil,
		"DefaultToDefer":      nil,
		//DB
		"Persistence": &MongoPersistence{
			DialUrl: "localhost",
			InitHandlers: []func(*mgo.Database){
				role.Init,
				user.Init,
			},
			Database: "authexample",
		},
	}

	g := NewGonk(c)
	g.Handle("/", HTTPHandler{View: "index"})

	// user controller
	g.Handle("/user/", HTTPHandler{
		Handler: user.Index,
		View:    "user/index",
	})

	g.Handle("/user/create", HTTPHandler{
		Handler: user.Create,
		View:    "user/create",
	}).Methods("GET")

	g.Handle("/user/create", HTTPHandler{
		Handler: user.Create,
	}).Methods("POST")

	g.Handle("/user/retrieve/{id}", HTTPHandler{
		Handler: user.Retrieve,
		View:    "user/retrieve",
	})

	g.Handle("/user/update/{id}", HTTPHandler{
		Handler: user.Update,
		View:    "user/update",
	})

	g.Handle("/user/delete/{id}", HTTPHandler{
		Handler: user.Delete,
	})

	// role controller
	g.Handle("/role/", HTTPHandler{
		Handler: role.Index,
		View:    "role/index",
	})

	g.Handle("/role/retrieve/{id}", HTTPHandler{
		Handler: role.Retrieve,
		View:    "role/retrieve",
	})

	g.Handle("/role/create", HTTPHandler{
		Handler: role.Create,
		View:    "role/create",
	}).Methods("GET")

	g.Handle("/role/create", HTTPHandler{
		Handler: role.Create,
	}).Methods("POST")

	g.Handle("/role/delete/{id}", HTTPHandler{
		Handler: role.Delete,
	})

	g.Handle("/role/update/{id}", HTTPHandler{
		Handler: role.Update,
		View:    "role/update",
	})

	g.Handle("/role/{id}/user/", HTTPHandler{
		Handler: role.UserIndex,
		View:    "role/user/index",
	})

	//Could also handle this with a POST/DELETE http method instead of .../add & .../remove
	g.Handle("/role/{id}/user/{uid}/add", HTTPHandler{
		Handler: role.UserAdd,
	})

	g.Handle("/role/{id}/user/{uid}/remove", HTTPHandler{
		Handler: role.UserRemove,
	})

	g.Start()
}
