package main

import (
	"fmt"
	"test/db"
	"test/demo1"
	"test/iris"
)

func main() {
	demo1.InterfaceDemo()
	demo1.TestFunc()
	demo1.Test()
	//db.QueryData(1)
	db.Query()
	db.Insert()
	db.Update()
	db.Delete()
	demo1.ReadFile()
	fmt.Print(db.Print("11"))
	iris.Service()
}
