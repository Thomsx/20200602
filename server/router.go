package main

import (
	//"fmt"
	"model"
	"net/http"
	"strings"
)

// WRoute 是代表对应请求的路由规则以及权限的结构体.
type WRoute struct {
	URL         string
	Permission  int
	HandlerFunc THandlerFunc
}

// HandlerFunc 是自定义的请求处理函数,接受*Handler作为参数.
type THandlerFunc func(*Handler)

// Handler 是包含一些请求上下文的结构体.
type Handler struct {
	http.ResponseWriter
	*http.Request
	//StartTime    time.Time     //接受请求时间
}

// NewHandler返回含有请求上下文的Handler.
func NewHandler(w http.ResponseWriter, r *http.Request) *Handler {
	return &Handler{
		ResponseWriter: w,
		Request:        r,
		//StartTime:      time.Now(),
	}
}

//处理程序
func handlerFun(wroute WRoute) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*
			defer func() {
				if e := recover(); e != nil {
					fmt.Println("panic:", e)
				}
			}()*/

		handler := NewHandler(w, r)
		//defer handler.Session.Close()

		url := r.Method + " " + r.URL.Path
		if r.URL.RawQuery != "" {
			url += "?" + r.URL.RawQuery
		}

		if wroute.Permission&Everyone == Everyone {
			wroute.HandlerFunc(handler)
		}
		var (
			user *model.User
			ok   bool
		)
		if wroute.Permission&Authenticated == Authenticated {
			user, ok = currentUser(handler)
			//fmt.Println(user.Uid)
			if !ok {
				http.Redirect(w, r, "/static", http.StatusFound)
				return
			}

			if wroute.Permission&AdministratorOnly == AdministratorOnly {
				if !user.IsAdmin {
					//message(handler, "没有权限", "对不起，你没有权限进行该操作", "error")
					return
				}
			}

			wroute.HandlerFunc(handler)
		}
	}
}

//http转https
func red(w http.ResponseWriter, req *http.Request) {
	var url string
	if strings.Contains(req.Host, ":") {
		url = "https://" + req.Host[:strings.Index(req.Host, ":")] + ":" + req.RequestURI
	} else {
		url = "https://" + req.Host + ":" + req.RequestURI
	}
	http.Redirect(w, req, url, http.StatusMovedPermanently)

}

// 路由规则
var routes = []WRoute{
	{"/", Everyone, indexHandler},
	{"/login", Everyone, loginHandler},
}
