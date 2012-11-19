package gonk

type Route map[string]HTTPHandler

/*import (
	"github.com/gorilla/mux"
	"github.com/netbrain/gonk/app/role"
	"github.com/netbrain/gonk/app/user"
	. "github.com/netbrain/gonk/model/web"
	"labix.org/v2/mgo"
)

var router *mux.Router

func init() {
	/*router = mux.NewRouter()
	router.StrictSlash(true)

	// front page
	router.Handle("/", HTTPHandler{view: "index"})

	// user controller
	router.Handle("/user/", HTTPHandler{
		before: initSessionBeforeRequest,
		handler: user.Index,
		todefer:  closeSessionAfterRequest,
		view:    "user/index"})

	router.Handle("/user/create", HTTPHandler{
		handler: user.Create,
		dbinit:  user.Init,
		view:    "user/create"}).
		Methods("GET")

	router.Handle("/user/create", HTTPHandler{
		handler: user.Create,
		dbinit:  user.Init}).
		Methods("POST")

	router.Handle("/user/retrieve/{id}", HTTPHandler{
		handler: user.Retrieve,
		dbinit:  user.Init,
		view:    "user/retrieve"})

	router.Handle("/user/update/{id}", HTTPHandler{
		handler: user.Update,
		dbinit:  user.Init,
		view:    "user/update"})

	router.Handle("/user/delete/{id}", HTTPHandler{
		handler: user.Delete,
		dbinit:  user.Init})

	// role controller
	router.Handle("/role/", HTTPHandler{
		handler: role.Index,
		dbinit:  role.Init,
		view:    "role/index"})

	router.Handle("/role/retrieve/{id}", HTTPHandler{
		handler: role.Retrieve,
		dbinit:  role.Init,
		view:    "role/retrieve"})

	router.Handle("/role/create", HTTPHandler{
		handler: role.Create,
		dbinit:  role.Init,
		view:    "role/create"}).
		Methods("GET")

	router.Handle("/role/create", HTTPHandler{
		handler: role.Create,
		dbinit:  role.Init}).
		Methods("POST")

	router.Handle("/role/delete/{id}", HTTPHandler{
		handler: role.Delete,
		dbinit:  role.Init})

	router.Handle("/role/update/{id}", HTTPHandler{
		handler: role.Update,
		dbinit:  role.Init,
		view:    "role/update"})

	router.Handle("/role/{id}/user/", HTTPHandler{
		handler: role.UserIndex,
		dbinit: func(arg *mgo.Database) {
			role.Init(arg)
			user.Init(arg)
		},
		view: "role/user/index"})

	//Could also handle this with a POST/DELETE http method instead of .../add & .../remove
	router.Handle("/role/{id}/user/{uid}/add", HTTPHandler{
		handler: role.UserAdd,
		dbinit: func(arg *mgo.Database) {
			role.Init(arg)
			user.Init(arg)
		}})

	router.Handle("/role/{id}/user/{uid}/remove", HTTPHandler{
		handler: role.UserRemove,
		dbinit: func(arg *mgo.Database) {
			role.Init(arg)
			user.Init(arg)
		}})

}
*/
