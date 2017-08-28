package model

import "github.com/astaxie/beego/orm"

func init() {
    orm.RegisterDriver("postgres", orm.DRPostgres)
    
	orm.RegisterDataBase("default", "postgres", "host=localhost port=5432 user=postgres password=adm123 dbname=default sslmode=disable")
	//orm.RegisterModel(new(users))
    //orm.RunSyncdb("default", true, true)
}
