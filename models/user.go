package models

import (
	"fmt";
	_"github.com\axtaxie\beego\orm";
	_"github.com\lib\pq"
)

type users struct {
    Id          int `orm:"pk"`
    age         int
    first_name  string
    last_name   string
    email       string
}

func Main(){
	fmt.Println("oi")

}