package main

import (
	"fmt"
	"test/db"
	"test/demo1"
	"test/iris"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server Petstore server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /v2
func main() {
	demo1.InterfaceDemo()
	demo1.TestFunc()
	demo1.Test()
	demo1.Example()
	//db.QueryData(1)
	db.Query()
	db.Insert()
	db.Update()
	db.Delete()
	demo1.ReadFile()
	demo1.GoroutineTest()
	fmt.Print(db.Print("11"))
	iris.Service()
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover ---->>>> %v \n", r)
		}
	}()
}
