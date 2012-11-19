package gonk

type Persistence interface {
	Init()
	InitOnRequest() interface{}
	CloseAfterRequest(interface{})
}
