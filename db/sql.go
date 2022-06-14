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

func QueryData() {
	user := User{}
	fmt.Println("starting ......")
	err := _db.First(&user)
	fmt.Println("err", err)
	fmt.Println("select data ", user)
}

func InsertData() {
	//user := User{
	//	Age:      18,
	//	UserName: "zhansan",
	//}
	//_db.CreateInBatches(user, 100)
	_db.Model(&User{}).Create(map[string]interface{}{
		"Age":      18,
		"Username": "zhansan",
		"CreateAt": time.Now(),
	})
}
