package role

import (
	"github.com/netbrain/gonk/app/user"
	"github.com/netbrain/gonk/model/server"
	"net/http"
)

func Index(ctx web.Context) interface{} {
	return All()
}

func Create(ctx web.Context) interface{} {
	if ctx.Request.Method == "POST" {
		if name := ctx.Params.Get("name"); name != "" {
			r := Role{Name: name}
			SaveOrUpdate(r)
			//redirects to Index
			http.Redirect(ctx.Writer, ctx.Request, "/role/", http.StatusTemporaryRedirect)
		} else {
			panic("Missing name parameter")
		}
	}
	return nil
}

func Retrieve(ctx web.Context) interface{} {
	if id := ctx.Params.Get("id"); id != "" {
		return Get(id)
	} else {
		panic("Missing id parameter")
	}
	return nil
}

func Update(ctx web.Context) interface{} {
	if id := ctx.Params.Get("id"); id != "" {
		if ctx.Request.Method == "POST" {
			if name := ctx.Params.Get("name"); name != "" {
				r := Get(id)
				r.Name = name
				SaveOrUpdate(r)
			} else {
				panic("Missing name parameter")
			}
		} else {
			return Get(id)
		}
	} else {
		panic("Missing id parameter")
	}
	http.Redirect(ctx.Writer, ctx.Request, "/role/", http.StatusTemporaryRedirect)
	return nil
}

func Delete(ctx web.Context) interface{} {
	if id := ctx.Params.Get("id"); id != "" {
		Remove(id)
		//redirects to Index
		http.Redirect(ctx.Writer, ctx.Request, "/role/", http.StatusTemporaryRedirect)
		return nil
	} else {
		panic("Missing id parameter")
	}
	return nil
}

func UserIndex(ctx web.Context) interface{} {
	id := ctx.Params.Get("id")
	return struct {
		ID    string
		Users []user.User
	}{
		id,
		user.All(),
	}
}

func UserAdd(ctx web.Context) interface{} {
	id := ctx.Params.Get("id")
	uid := ctx.Params.Get("uid")
	if id != "" && uid != "" {
		r := Get(id)
		u := user.User{ID: user.Get(uid).ID}
		var appended bool
		if r.Users, appended = appendIfMissing(r.Users, u); appended {
			SaveOrUpdate(r)
		}
	} else {
		panic("Missing parameters")
	}
	http.Redirect(ctx.Writer, ctx.Request, "/role/retrieve/"+id, http.StatusTemporaryRedirect)
	return nil
}

func UserRemove(ctx web.Context) interface{} {
	id := ctx.Params.Get("id")
	uid := ctx.Params.Get("uid")
	if id != "" && uid != "" {
		r := Get(id)
		u := user.User{ID: user.Get(uid).ID}
		var removed bool
		if r.Users, removed = removeFromSlice(r.Users, u); removed {
			SaveOrUpdate(r)
		}
	} else {
		panic("Missing parameters")
	}
	http.Redirect(ctx.Writer, ctx.Request, "/role/retrieve/"+id, http.StatusTemporaryRedirect)
	return nil
}

func appendIfMissing(slice []user.User, u user.User) ([]user.User, bool) {
	for _, ele := range slice {
		if u.Equals(ele) {
			return slice, false
		}
	}
	return append(slice, u), true
}

func removeFromSlice(slice []user.User, u user.User) ([]user.User, bool) {
	removed := false
	newslice := make([]user.User, 0, len(slice))
	for _, ele := range slice {
		if !u.Equals(ele) && ele.ID != "" {
			newslice = append(newslice, ele)
		} else {
			removed = true
		}
	}
	return newslice, removed
}
