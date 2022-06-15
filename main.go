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
	db.QueryData(1)
	db.Query()
	db.Insert()
	db.Update()
	db.Delete()

	fmt.Print(db.Print("11"))
	iris.Service()
}
