package gonk

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"runtime/debug"
)

type HTTPHandler struct {
	Before      func(ctx Context) (bool, error)
	Handler     func(ctx Context) interface{}
	After       func(ctx Context) (bool, error)
	ToDefer     func(ctx Context)
	View        string
	Status      HTTPStatus
	Defaults    DefaultHandlers
	Persistence Persistence
}

type HTTPStatus struct {
	Code int
	Text string
}

type DefaultHandlers struct {
	InitContextHandler  func(h HTTPHandler, w http.ResponseWriter, r *http.Request) Context
	FailureHandler      func(h HTTPHandler, w http.ResponseWriter)
	RenderViewHandler   func(h HTTPHandler, w http.ResponseWriter, data interface{})
	RenderStatusHandler func(h HTTPHandler, w http.ResponseWriter)
}

func (h HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//init failure handling
	defer func() {
		if h.Defaults.FailureHandler == nil {
			h.FailureHandler(w)
		} else {
			h.Defaults.FailureHandler(h, w)
		}
	}()

	//init context
	var ctx Context
	if h.Defaults.InitContextHandler == nil {
		ctx = &DefaultContext{}
		h.InitContext(w, r, ctx)
	} else {
		ctx = h.Defaults.InitContextHandler(h, w, r)
	}

	//init persistence
	if h.Persistence != nil {
		pObj := h.Persistence.InitOnRequest()
		defer h.Persistence.CloseAfterRequest(pObj)
	}

	//init other defer funcs
	h.executeToDefer(ctx)

	//execute the before logic
	if proceed, err := h.executeBefore(ctx); !proceed {
		if err != nil {
			panic(err)
		}
		return
	}

	//execute the handler
	data := h.executeHandler(ctx)

	//execute the after logic
	if proceed, err := h.executeAfter(ctx); !proceed {
		if err != nil {
			panic(err)
		}
		return
	}

	//render the view
	if h.Defaults.RenderViewHandler == nil {
		h.RenderView(w, &data)
	} else {
		h.Defaults.RenderViewHandler(h, w, &data)
	}

	//render http status & text
	if h.Defaults.RenderStatusHandler == nil {
		h.RenderStatus(w)
	} else {
		h.Defaults.RenderStatusHandler(h, w)
	}
}

func (h HTTPHandler) FailureHandler(w http.ResponseWriter) {
	if r := recover(); r != nil {
		http.Error(w, fmt.Sprintf("PANIC: %s - %s", r, debug.Stack()), http.StatusInternalServerError)
		log.Printf("PANIC: %s - %s", r, debug.Stack())
	}
}

func (h HTTPHandler) InitContext(w http.ResponseWriter, r *http.Request, ctx Context) {
	r.ParseMultipartForm(0)
	for k, v := range mux.Vars(r) {
		r.Form.Add(k, v)
	}

	ctx.SetWriter(w)
	ctx.SetRequest(r)
	ctx.SetParams(&r.Form)
}

func (h HTTPHandler) executeToDefer(ctx Context) {
	if h.ToDefer != nil {
		defer h.ToDefer(ctx)
	}
}

func (h HTTPHandler) executeBefore(ctx Context) (bool, error) {
	if h.Before != nil {
		return h.Before(ctx)
	}
	return true, nil
}

func (h HTTPHandler) executeHandler(ctx Context) interface{} {
	if h.Handler != nil {
		return h.Handler(ctx)
	}
	return nil
}

func (h HTTPHandler) executeAfter(ctx Context) (bool, error) {
	if h.After != nil {
		return h.After(ctx)
	}
	return true, nil
}

func (h HTTPHandler) RenderView(w http.ResponseWriter, data interface{}) {
	if h.View != "" {
		tmpl, err := template.ParseFiles("app/" + h.View + ".html")
		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			panic(err)
		}
	}
}

func (h HTTPHandler) RenderStatus(w http.ResponseWriter) {
	if h.Status.Code != 0 {
		text := h.Status.Text
		if text == "" {
			text = http.StatusText(h.Status.Code)
		}
		http.Error(w, text, h.Status.Code)
	}
}

func CollectAndWrap(funcs ...func(ctx Context)) func(ctx Context) {
	return func(ctx Context) {
		for _, f := range funcs {
			f(ctx)
		}
	}
}
