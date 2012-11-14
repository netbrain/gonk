package user

import (
	"github.com/netbrain/gonk/model/server"
	"net/http"
)

func Index(ctx web.Context) interface{} {
	return All()
}

func Create(ctx web.Context) interface{} {
	if ctx.Request.Method == "POST" {
		email := ctx.Params.Get("email")
		password := ctx.Params.Get("password")
		if email != "" && password != "" {
			u := User{Email: email}
			u.SetPassword(password)
			SaveOrUpdate(u)
			//redirects to Index
			http.Redirect(ctx.Writer, ctx.Request, "/user/", http.StatusTemporaryRedirect)
		} else {
			panic("Missing parameters")
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
			email := ctx.Params.Get("email")
			password := ctx.Params.Get("password")
			if email != "" && password != "" {
				u := Get(id)
				u.Email = email
				u.SetPassword(password)
				SaveOrUpdate(u)
			} else {
				panic("Missing parameters")
			}
		} else {
			return Get(id)
		}
	} else {
		panic("Missing id parameter")
	}
	http.Redirect(ctx.Writer, ctx.Request, "/user/", http.StatusTemporaryRedirect)
	return nil
}

func Delete(ctx web.Context) interface{} {
	if id := ctx.Params.Get("id"); id != "" {
		Remove(id)
		//redirects to Index
		http.Redirect(ctx.Writer, ctx.Request, "/user/", http.StatusTemporaryRedirect)
		return nil
	} else {
		panic("Missing id parameter")
	}
	return nil
}
