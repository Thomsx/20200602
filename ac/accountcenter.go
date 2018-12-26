package ac

import (
	//"errors"
	"enhance/pgx"
	//"fmt"
	//"log"
	"model"
)

var DBconfig = pgx.ConnConfig{
	Host:     "localhost",
	User:     "postgres",
	Password: "123",
	Database: "shipcrew",
}
var ConnPoolConfig = pgx.ConnPoolConfig{DBconfig, 5, nil}

/*
type User struct {
	Uid      string //用户iD
	Utype    int    //类型 0：未认证 1:普通（个人船员）2:企业
	Name     string //名字
	Password string //密码
	Email    string //email
	State    int
	IsAdmin  bool   //管理员权限
}
*/
func Login(username string, passwd string) (*model.User, error) {
	user := model.NewUser()
	conn, err := pgx.NewConnPool(ConnPoolConfig)

	if err != nil {
		panic(err)
	}
	defer conn.Close()
	err = conn.QueryRow("select user_id,user_state from shipcrew.web.web_user where user_name=$1", username).Scan(&user.Uid, &user.State)
	if err != nil {
		//没有找到此人会抛出错误，查无此人在此处理
		user.Uid = "Hacker-"
		user.State = 0
		user.Name = username
		return user, nil
	}

	return user, nil
}
