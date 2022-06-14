package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

const (
	dbUser     string = "admin"
	dbPassword string = "Sugarotect2022**"
	dbHost     string = "test.cjyntu0au13f.ap-southeast-1.rds.amazonaws.com"
	dbPort     int    = 3306
	dbName     string = "TEST"
)

type User struct {
	Id       int
	Age      int
	UserName string    `db:"user_name"`
	CreateAt time.Time `db:"create_at"`
}

//var user User
var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
var _db *gorm.DB

func init() {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if nil == err {
		fmt.Println("data base connect success")
	} else {
		fmt.Println("data base connect fail", err)
	}
	//db.Debug().Select("SELECT * FROM user")
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(time.Hour)

	_db = db
}

func QueryData(id int) (re User) {
	user := User{Id: id}
	fmt.Println("starting ......")
	//err := _db.First(&user)
	res := _db.Find(&user)
	if err := res.Error; err != nil {
		panic("no have data" + res.Error.Error())
	} else {
		fmt.Println("select data ", res)
		return user
	}
}

func InsertData(age int, username string) (res User) {
	user := User{0, age,
		username, time.Now()}
	_res := _db.Create(&user)
	if err := _res.Error; nil != err {
		panic("no have data" + _res.Error.Error())
	} else {
		fmt.Println("insert data res --->>>", user)
		return user
	}

}
