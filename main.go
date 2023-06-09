package main

import (
	// "Golang_web/models"
	_ "Golang_web/routers"
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var DBARGS = struct {
	Driver string
	Source string
	Debug  string
}{
	os.Getenv("ORM_DRIVER"),
	os.Getenv("ORM_SOURCE"),
	os.Getenv("ORM_DEBUG"),
}

func init() {
	Debug, _ := orm.StrTo(DBARGS.Debug).Bool()

	fmt.Printf("Debug=%v\n", Debug)
	if Debug {
		fmt.Printf("ORM config=%v\n", DBARGS)
	}
	if DBARGS.Driver == "" || DBARGS.Source == "" {
		fmt.Println("Please set ORM_DRIVER/ORM_SOURCE")
		os.Exit(2)
	}

	// set default database
	orm.RegisterDataBase("default", DBARGS.Driver, DBARGS.Source, 30)
}

func main() {
	beego.Run()
	// user, err := models.GetUserById(1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(user)
	
}
