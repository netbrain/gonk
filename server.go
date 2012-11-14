package main

import "net/http"
import "html/template"
import "github.com/netbrain/gonk/model/server"
import "github.com/gorilla/mux"
import "labix.org/v2/mgo"
import "fmt"
import "log"
import "runtime/debug"

type HTTPHandler struct {
	handler       func(ctx web.Context) interface{}
	dbinit        func(db *mgo.Database)
	view          string
	dbinitialized bool
}

func initServer() {
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}

func (h HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			http.Error(w, fmt.Sprintf("PANIC: %s - %s", r, debug.Stack()), http.StatusInternalServerError)
			log.Printf("PANIC: %s - %s", r, debug.Stack())
		}
	}()
	r.ParseMultipartForm(0) //parse request parameters

	ctx := web.Context{Writer: w, Request: r, Params: r.Form}
	if h.dbinit != nil {
		ctx.DBSession = session.Clone()
		defer ctx.DBSession.Close()
		h.dbinit(ctx.DBSession.DB("gonk"))
	}

	//parse mux variables
	for k, v := range mux.Vars(r) {
		ctx.Params.Add(k, v)
	}

	var data interface{}
	if h.handler != nil {
		data = h.handler(ctx)
	}

	if h.view != "" {
		tmpl, err := template.ParseFiles("app/" + h.view + ".html")
		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			panic(err)
		}
	}
}
