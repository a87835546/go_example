package db

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"time"
)

var Db *sqlx.DB

func init() {

	_dsn := "admin:Sugarotect2022**@tcp(test.cjyntu0au13f.ap-southeast-1.rds.amazonaws.com:3306)/TEST?charset=utf8mb4&parseTime=True"
	td, err := sqlx.Connect("mysql", _dsn)
	if nil != err {
		fmt.Printf("connect DB failed,err:%v\n", err)
		return
	}
	Db = td

	fmt.Println("db", Db)
	Db.SetMaxOpenConns(20)
	Db.SetMaxIdleConns(10)
	//s := goqu.S("TEST")
	//t := s.Table("users")
	//sql, _, _ := goqu.From(t).Select(t.Col("user_name")).ToSQL()
	//fmt.Println("sql", sql)
	//Print("db inited")
}
func Query() {
	g := goqu.Dialect("mysql")

	fmt.Println("query  db", Db)
	user := User{}
	//users := make([]User, 0)
	//ds := goqu.From("users")
	//sql, _, _ := goqu.From("users").Select(goqu.C("id").Eq(1)).ToSQL()
	sql, agrs, err := g.From("users").Where(goqu.C("user_name").Eq("zhansan")).ToSQL()
	fmt.Println("A ", agrs)
	//sql, _, _ := goqu.From("users").ToSQL()
	//sql, _, _ := ds.Limit(1).ToSQL()

	var s []User
	fmt.Println(sql)
	err = Db.Select(&s, sql)
	//err = db.Select(&s, "SELECT * FROM  users where user_name = 'zhansan' ")
	if nil != err {
		fmt.Printf("query failed %v\n", err)
	} else {
		fmt.Printf("query res -->>%v\n", user)
	}
}

func Insert() {
	//sql, _, _ := goqu.Insert("users").Rows(goqu.Record{"user_name": "lisi", "age": 12, "create_at": time.Now()}).ToSQL()
	sql, _, _ := goqu.From("users").Insert().Rows(goqu.Record{"user_name": "lisi", "age": 12, "create_at": time.Now()}).ToSQL()
	fmt.Print("insert sql --->>>", sql)
	res, err := Db.Exec(sql)
	if nil != err {
		fmt.Printf("insert failed %v\n", err)
	} else {
		fmt.Printf("insert res -->>%v\n", res)
	}
}

func Update() {
	sql, _, _ := goqu.From("users").Update().Set(goqu.Record{"user_name": "lisi"}).Where(goqu.C("id").Eq(1)).ToSQL()
	fmt.Print("update sql --->>>", sql)
	res, err := Db.Exec(sql)
	if nil != err {
		fmt.Printf("update failed %v\n", err)
	} else {
		fmt.Printf("update res -->>%v\n", res)
	}
}

func Delete() {
	sql, _, _ := goqu.From("users").Delete().Where(goqu.C("id").Eq(3)).ToSQL()
	fmt.Print("delete sql --->>>", sql)
	res, err := Db.Exec(sql)
	if nil != err {
		fmt.Printf("delete failed %v\n", err)
	} else {
		fmt.Printf("delete res -->>%v\n", res)
	}
}

func Print(name string) string {
	if nil == Db.DB {
		return "db dont init"
	}
	res, err := Db.Exec("SELECT * FROM users")
	if nil != err {
		fmt.Printf("query failed %v\n", err)
	} else {
		fmt.Printf("query res -->>%v\n", &res)
	}
	fmt.Println("test", name)
	return ""
}
