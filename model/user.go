package model

//用户结构 普通 企业 企业域名
type User struct {
	Uid      string //用户iD
	Utype    int    //类型 0：未认证 1:普通（个人船员）2:企业
	Name     string //名字
	Password string //密码
	Email    string //email
	State    int16  //状态 0：禁用 1：审核 2：激活 3：
	IsAdmin  bool   //管理员权限
}

func NewUser() *User {
	return &User{}
}

//用户认证为船员
type Seaman struct {
	User
	/*以下为海员信息*/
}

//用户认证为船东
type Shipowner struct {
	User
	/*企业信息*/
	Edomain string //企业域名
}

//用户认证为船员管理公司
type CrewManage struct {
	User
	/*企业信息*/
	Edomain string //企业域名
}
