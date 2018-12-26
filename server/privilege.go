package main

const (
	Everyone          int = 1 << 0                            //访客             001
	Authenticated         = 1 << 1                            //已登陆           010
	AdministratorOnly     = 1 << 2                            //管理员           100
	Administrator         = AdministratorOnly | Authenticated //管理员也要已登陆 110
)
