package gonk

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Config map[string]interface{}

type Gonk struct {
	rmap        *RouteMap
	router      *mux.Router
	config      Config
	persistence Persistence
}

func NewGonk(c Config) *Gonk {
	g := &Gonk{
		config: c,
		rmap:   NewRouteMap(),
	}

	if persistence, present := g.config["Persistence"]; present {
		g.persistence = persistence.(Persistence)
	}

	g.router = mux.NewRouter()
	v, ok := c["StrictSlash"].(bool)
	if ok {
		g.router.StrictSlash(v)
	}
	return g
}

func (g Gonk) Handle(r string, h HTTPHandler) *mux.Route {
	g.addDefaultHandlerConfig(&h)
	g.rmap.Add(r, &h)
	return g.router.Handle(r, h)
}

func (g Gonk) addDefaultHandlerConfig(h *HTTPHandler) {
	var handler interface{}
	var present bool

	if h.Persistence == nil {
		h.Persistence = g.persistence
	}

	if handler, present = g.config["InitContextHandler"]; present {
		if f, ok := handler.(func(h HTTPHandler, w http.ResponseWriter, r *http.Request) Context); ok {
			h.Defaults.InitContextHandler = f
		}
	}

	if handler, present = g.config["FailureHandler"]; present {
		if f, ok := handler.(func(h HTTPHandler, w http.ResponseWriter)); ok {
			h.Defaults.FailureHandler = f
		}
	}

	if handler, present = g.config["RenderViewHandler"]; present {
		if f, ok := handler.(func(h HTTPHandler, w http.ResponseWriter, data interface{})); ok {
			h.Defaults.RenderViewHandler = f
		}
	}

	if handler, present = g.config["RenderStatusHandler"]; present {
		if f, ok := handler.(func(h HTTPHandler, w http.ResponseWriter)); ok {
			h.Defaults.RenderStatusHandler = f
		}
	}

}

func (g Gonk) Start() {
	//init persistence layer
	g.persistence.Init()

	//start http server
	serverAddr, _ := g.config["ServerAddr"].(string)
	if err := http.ListenAndServe(serverAddr, g.router); err != nil {
		panic(err)
	}
}
