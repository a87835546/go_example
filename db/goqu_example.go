package db

import (
	"encoding/json"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	jsoniter "github.com/json-iterator/go"
	"log"
	"reflect"
	"test/model"
	"time"
)

var Db *sqlx.DB
var G = goqu.Dialect("mysql")

func init() {

	go func() {
		_dsn := "root:123456abc@tcp(ec2-13-215-251-145.ap-southeast-1.compute.amazonaws.com:3306)/demo?charset=utf8mb4&parseTime=True"
		td, err := sqlx.Connect("mysql", _dsn)
		if nil != err {
			log.Printf("connect DB failed,err:%v\n", err)
			return
		}
		Db = td
		Db.SetMaxOpenConns(20)
		Db.SetMaxIdleConns(10)
	}()
	//s := goqu.S("TEST")
	//t := s.Table("users")
	//sql, _, _ := goqu.From(t).Select(t.Col("user_name")).ToSQL()
	//log.Println("sql", sql)
	//Print("db inited")
}
func Query() {

	log.Println("query  db", Db)
	var users []User
	//users := make([]User, 0)
	//ds := goqu.From("users")
	//sql, _, _ := goqu.From("users").Select(goqu.C("id").Eq(1)).ToSQL()
	//sql, agrs, err := g.From("users").Where(goqu.C("user_name").Eq("zhansan")).ToSQL()
	ex := goqu.Ex{
		"user_name": "zhansan",
	}
	sql, agrs, err := G.From("users").Where(ex).ToSQL()
	log.Println("A ", agrs)
	//sql, _, _ := goqu.From("users").ToSQL()
	//sql, _, _ := ds.Limit(1).ToSQL()

	log.Println(sql)
	//res, err := Db.Exec(sql)
	//err = db.Select(&s, "SELECT * FROM  users where user_name = 'zhansan' ")
	log.Println("res", users)
	if nil != err {
		log.Printf("query failed %v\n", err)
	} else {
		log.Printf("query res -->>%v\n", users)
	}
	defer func() {
		if r := recover(); r != nil {

		} else {
			err = Db.Select(&users, sql)

		}
	}()
}

func QueryById(id int) []User {
	var users []User
	ex := goqu.Ex{
		"id": id,
	}
	sql, _, err := G.From("users").Where(ex).ToSQL()
	err = Db.Select(&users, sql)
	log.Println("res", users)
	if nil != err {
		log.Printf("query failed %v\n", err)
	} else {
		log.Printf("query res -->>%v\n", users)
	}
	return users
}

func QueryAllUsers() []User {
	var users []User
	sql, _, _ := G.From("users").Select().ToSQL()
	res, err := Db.Queryx(sql)
	if nil != err {
		log.Printf("err --->>>> %s", err.Error())
	}
	for res.Next() {
		var p User
		err = res.StructScan(&p)
		users = append(users, p)
	}
	return users
}

func Insert() {
	//sql, _, _ := goqu.Insert("users").Rows(goqu.Record{"user_name": "lisi", "age": 12, "create_at": time.Now()}).ToSQL()
	d := goqu.Record{
		"user_name": "lisi",
		"age":       12,
		"create_at": time.Now().Unix(),
	}
	sql, _, _ := G.From("users").Insert().Rows(d).ToSQL()

	log.Print("insert sql --->>>", sql)
	res, err := Db.Exec(sql)
	if nil != err {
		log.Printf("insert failed %v\n", err)
	} else {
		log.Printf("insert res -->>%v\n", res)
	}
}

func Update() {
	sql, _, _ := G.From("users").Update().Set(goqu.Record{"user_name": "lisi"}).Where(goqu.C("id").Eq(1)).ToSQL()
	log.Print("update sql --->>>", sql)
	res, err := Db.Exec(sql)
	if nil != err {
		log.Printf("update failed %v\n", err)
	} else {
		log.Printf("update res -->>%v\n", res)
	}
}

func UpdateUserTitle(title string, id int) bool {
	sql := fmt.Sprintf("UPDATE users set title = %s WHERE id = %d", title, id)
	log.Print("update sql --->>>", sql)
	res, err := Db.Exec(sql)

	if nil != err {
		log.Printf("update failed %v\n", err)
		return false
	} else {
		log.Printf("update res -->>%v\n", res)
		return true
	}

}

// FetchTitleById  查询中文的title/
func FetchTitleById(id int) string {
	sql := fmt.Sprintf("SELECT json_extract(title,'$.title_CN') as title FROM  users WHERE id = %d", id)
	log.Print("update sql --->>>", sql)
	//res, err := Db.Exec(sql)
	res, err := Db.Query(sql)
	var p string
	if nil != err {
		log.Printf("update failed %v\n", err)
	} else {
		for res.Next() {
			err = res.Scan(&p)
		}
		log.Printf("update success %v\n", p)
	}
	return p
}

// UpdateTitleById  更新中文的title/
func UpdateTitleById(id int, cn string, i int) bool {
	selectSql := fmt.Sprintf("SELECT title FROM users where id = %d", id)
	res, err := Db.Query(selectSql)
	var title string
	for res.Next() {
		err = res.Scan(&title)
	}
	lan := model.Language{}
	err = json.Unmarshal([]byte(title), &lan)
	switch i {
	case 0:
		lan.CN = cn
		break
	case 1:
		lan.EN = cn
	case 2:
		lan.KR = cn
	}
	if err != nil {
		log.Printf("json 解析 struct 异常 --->>%s", err.Error())
	}
	var stringLanguage []byte
	if stringLanguage, err = json.Marshal(lan); err != nil {
		log.Printf("struct 解析 string 异常 --->>%s", err.Error())
		return false
	}
	if res1, err1 := jsoniter.MarshalToString(string(stringLanguage)); err1 == nil {
		log.Printf("title --->>>>>%s  %s  lan --->>>%s", res1, reflect.TypeOf(res1), lan)
		return UpdateUserTitle(res1, id)
	}
	return false
}

func Delete() {
	sql, _, _ := G.From("users").Delete().Where(goqu.C("id").Eq(3)).ToSQL()
	log.Print("delete sql --->>>", sql)
	res, err := Db.Exec(sql)
	if nil != err {
		log.Printf("delete failed %v\n", err)
	} else {
		log.Printf("delete res -->>%v\n", res)
	}
}

func Print(name string) string {
	if nil == Db.DB {
		return "db dont init"
	}
	res, err := Db.Exec("SELECT * FROM users")
	if nil != err {
		log.Printf("query failed %v\n", err)
	} else {
		log.Printf("query res -->>%v\n", &res)
	}
	log.Println("test", name)
	return ""
}
