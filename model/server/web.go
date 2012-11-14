package web

import "net/url"
import "net/http"
import "labix.org/v2/mgo"

type Context struct {
	Writer    http.ResponseWriter
	Request   *http.Request
	DBSession *mgo.Session
	Params    url.Values
}
