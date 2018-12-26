package main

import (
	//"enhance/pgx"
	//"io" //testing
	"ac"
	"log"
	"model"
	"net/http"
)

// 返回当前用户
// 数据库设置在线表，uid，ip对照
func currentUser(handler *Handler) (*model.User, bool) {

	r := handler.Request

	session_id, _ := r.Cookie("session_id")

	remote_ip := r.RemoteAddr

	pool := createConnPool()

	var user_id, login_ip string

	user := model.User{}
	err := pool.QueryRow("select user_id,login_ip from shipcrews.dbo.online_user where session_id = $1 ", session_id).Scan(&user_id, &login_ip)
	if err != nil {
		log.Fatalf("account 23", err)
	}

	if user_id != "" {
		if remote_ip == login_ip {
			user.Uid = user_id
			return &user, true
		} else {
			user.Uid = user_id
			return &user, false
		}
	} else {
		return nil, false
	}
	/*if !ok {
		return nil, false
	}

	username = username.(string)

	user := User{}

	c := handler.DB.C(USERS)

	// 检查用户名
	err := c.Find(bson.M{"username": username}).One(&user)

	if err != nil {
		return nil, false
	}
	*/
}

// 登录
//data : username password type domain
// uri : /account/login{json:(|.json)}
func loginHandler(handler *Handler) {
	log.Print("login begin")
	//return
	//io.WriteString(handler.ResponseWriter, "login working")

	req := handler.Request
	rw := handler.ResponseWriter
	username := req.PostFormValue("username")
	if username == "" || req.Method != "POST" {
		//filter.SetData(req, map[string]interface{}{"error": "非法请求"})
		//req.Form.Set(filter.CONTENT_TPL_KEY, "/template/login.html")
		return
	}

	//vars := mux.Vars(req)

	//suffix := vars["json"]

	// 处理用户登录
	passwd := req.PostFormValue("passwd")
	//userLogin
	user, err := ac.Login(username, passwd)
	if err != nil {
		/* if suffix != "" {
			logger.Errorln("login error:", err)
			fmt.Fprint(rw, `{"ok":0,"error":"`+err.Error()+`"}`)
			return
		}

		req.Form.Set(filter.CONTENT_TPL_KEY, "/template/login.html")
		filter.SetData(req, map[string]interface{}{"username": username, "error": err.Error()})
		*/return
	}
	log.Print(user.Uid, user.Name, req.UserAgent(), req.Referer(), "login")
	//logger.Debugf("remember_me is %q\n", req.FormValue("remember_me"))
	// 登录成功，种cookie
	http.SetCookie(rw, &http.Cookie{
		Name:  "islogin",
		Value: user.Uid,
	})

	// 支持跳转到源页面
	//uri := "/"
	//values := filter.NewFlash(rw, req).Flashes("uri")
	//if values != nil {
	//	uri = values[0].(string)
	//}
	//logger.Debugln("uri===", uri)
	//util.Redirect(rw, req, uri)
	http.Redirect(rw, req, req.Referer(), http.StatusFound)
}
