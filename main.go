package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"net/rpc"
	"test/db"
	"test/demo1"
	app "test/iris"

	"time"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server Petstore server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// World @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
type World struct {
}

func (w *World) HelloWorld(name string, resp *string) error {
	*resp = name + "hello\n"
	return nil
}

func (w *World) Add(arg string, res *string) error {
	*res = arg + "add method "
	return nil
}

type User struct {
	Name string
	Age  int32
	Time int64
}

func (w *World) GetUsers(user User, reply *[]User) error {
	*reply = append(*reply, user)
	return nil
}

func (w *World) GetUser(user User, reply *User) error {
	fmt.Printf("user -->>>> %v\n", user)
	user.Age = user.Age + 10
	user.Name = user.Name + "操作后的数据"
	user.Time = time.Now().UnixNano()
	*reply = user
	return nil
}
func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	_app := iris.New()
	app.XenditPay()
	err := rpc.RegisterName("hello", new(World)) //注册rpc服务   需要绑定结构体
	if err != nil {
		fmt.Println("fail in reg rpc")
		return
	}

	//lis, err := net.Listen("tcp", "localhost:9999") //开始监听
	//
	//if err != nil {
	//	fmt.Println("fail in listen")
	//}
	//defer lis.Close()
	//service := hprpc.NewFastHTTPService()
	//service.AddFunction("hello", hello)
	//_app.Any("/test2", func(ctx *iris.Context) {
	//	service.ServeFastHTTP(ctx.RequestCtx)
	//})
	//func() {
	//	go app.Service(_app)
	//
	//	for {
	//		conn, err := lis.Accept() //建立连接
	//		if err != nil {
	//			fmt.Printf("fail in accept")
	//		}
	//		defer conn.Close()
	//		rpc.ServeConn(conn) //将连接绑定到rpc服务
	//	}
	//}()
	go demo1.InterfaceDemo()
	go demo1.ReadKey()

	go func() {
		demo1.TestFunc()
		demo1.Test()
		demo1.Example()
		db.QueryData(1)
	}()
	go db.Query()
	go db.Insert()
	go db.Update()
	go db.Delete()
	demo1.ReadFile()
	demo1.GoroutineTest()
	demo1.HSetValue("object", "1")
	fmt.Print(db.Print("11"))
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover ---->>>> %v \n", r)
		}
	}()
	app.Service(_app)

}
