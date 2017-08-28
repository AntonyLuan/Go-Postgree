package main

import (
	_ "ProjetoWeb/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

// package main

// import (
// 	"fmt"

// 	"github.com/astaxie/beego/orm"
// 	_ "github.com/lib/pq"
// )

// type User struct {
// 	Id   int
// 	Name string
// }

// func init() {
// 	orm.RegisterDriver("postgres", orm.DRPostgres)

// 	orm.RegisterDataBase("default", "postgres", "host=localhost port=5432 user=postgres password=adm123 dbname=default sslmode=disable")
// }

// func main() {
// 	o := orm.NewOrm()
// 	o.Using("teste") // Using default, you can use other database
// 	orm.RegisterModel(new(User))

// 	user := new(User)
// 	user.Name = "slene"

// 	fmt.Println(o.Insert(user))
//}
