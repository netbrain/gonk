package gonk

type RouteMap struct {
	byHandler map[*HTTPHandler]string
	byRoute   map[string]*HTTPHandler
}

func NewRouteMap() *RouteMap {
	return &RouteMap{
		byHandler: make(map[*HTTPHandler]string),
		byRoute:   make(map[string]*HTTPHandler),
	}
}

func (m RouteMap) Add(route string, handler *HTTPHandler) {
	m.byHandler[handler] = route
	m.byRoute[route] = handler
}

func (m RouteMap) GetByView(route string) *HTTPHandler {
	return m.byRoute[route]
}

func (m RouteMap) GetByHandler(handler *HTTPHandler) string {
	return m.byHandler[handler]
}
