package main

import (
	"test/db"
	"test/demo1"
	"test/iris"
)

func main() {
	demo1.TestFunc()
	//db.InsertData()
	db.QueryData()
	db.Query()
	db.Insert()
	db.Update()
	db.Delete()

	//fmt.Print(db.Print("11"))
	//db.InsertData()
	iris.Service()
}
