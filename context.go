package gonk

import "net/url"
import "net/http"

type Context interface {
	Writer() http.ResponseWriter
	SetWriter(http.ResponseWriter)
	Request() *http.Request
	SetRequest(*http.Request)
	Params() *url.Values
	SetParams(*url.Values)
}

type DefaultContext struct {
	writer      http.ResponseWriter
	request     *http.Request
	params      *url.Values
	persistence Persistence
}

func (d DefaultContext) Writer() http.ResponseWriter {
	return d.writer
}

func (d DefaultContext) Request() *http.Request {
	return d.request
}

func (d DefaultContext) Params() *url.Values {
	return d.params
}

func (d *DefaultContext) SetWriter(writer http.ResponseWriter) {
	d.writer = writer
}

func (d *DefaultContext) SetRequest(request *http.Request) {
	d.request = request
}

func (d *DefaultContext) SetParams(params *url.Values) {
	d.params = params
}
