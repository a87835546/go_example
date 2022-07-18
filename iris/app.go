package iris

import (
	"encoding/json"
	"fmt"
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/websocket"
	"io"
	"os"
	"test/consts"
	"test/demo2"
	"test/websocket_demo"
	"time"
)

type LoginVo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Service() {
	app := iris.New()

	config1 := &swagger.Config{
		URL: "http://localhost:8080/swagger/doc.json", //The url pointing to API definition
	}

	// use swagger middleware to
	app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(config1, swaggerFiles.Handler))

	app.Use(myMiddleware)
	app.Handle("GET", "/test", func(context iris.Context) {
		context.JSON(iris.Map{"message": "test"})
	})
	app.Handle("GET", "/test1", func(context iris.Context) {
		context.JSON(iris.Map{"message": "test1"})

	})
	app.Get("/test2/{id}", func(context iris.Context) {
		context.JSON(iris.Map{"message": "test2", "id": context.Params().Get("id")})
	})
	app.Get("/test3", func(context iris.Context) {
		context.JSON(iris.Map{"message": "test3", "name": context.URLParam("name")})
	})
	app.Post("/test4", func(context iris.Context) {
		fmt.Println("post value", context.PostValue("username"))

		result := consts.Result{
			Code:    200,
			Message: "test4",
			Data:    context.PostValue("username"),
		}
		fmt.Printf(
			"res--->>>%s\n",
			result,
		)
		if data, err := json.Marshal(result); err != nil {
			fmt.Print("data", data)
		} else {

			//res := json.NewDecoder(bytes.NewReader(data))
			context.JSON(iris.Map{
				"code":    200,
				"message": "test4",
				"data":    context.PostValue("username") + context.RemoteAddr(),
			})
		}
	})
	app.Post("/upload", func(context iris.Context) {
		file, fileHeader, err := context.FormFile("file")
		if nil != err {
			context.JSON(consts.Result{
				"upload file failed" + err.Error(),
				201,
				"11",
			})
			return
		} else {
			defer file.Close()
			dest := "file path" + fileHeader.Filename
			dest1, _ := os.Create(dest)
			_, err = io.Copy(dest1, file)
			if nil != err {
				context.JSON(consts.Result{
					"upload file failed" + err.Error(),
					201,
					"11",
				})
				return
			}
			context.JSON(consts.Result{
				"upload file success",
				200,
				"11",
			})
			defer dest1.Close()
		}
	})
	app.Post("/postJson", func(ctx iris.Context) {
		login := LoginVo{}
		if err := ctx.ReadJSON(&login); err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("login --->>%s\n", login)
		}
		ctx.JSON(iris.Map{
			"code":    200,
			"message": "post request json",
			"data":    login,
		})
	})

	// iris 路由的基本使用情况
	/*
	   注解的基本使用
	   *
	*/

	example := app.Party("example", ExampleHandle)
	example.Get("/test", ExampleRequest)
	example.Get("/result", ResultCommonRequest)
	example.Get("/query", Query)
	home := app.Party("home", HomeHandle)
	home.Get("/banner", GetBanner)
	home.Get("", GetHome)
	user := app.Party("user", UserHandle)
	user.Get("", GetUser)
	user.Post("/register", RegisterUser)
	user.Post("/queryAll", QueryAllUsers)
	user.Get("/updateById", UpdateUser)
	user.Get("/queryById", QueryUsersById)
	user.Post("/updateTitleById", UpdateUserTitleByGoqu)
	user.Get("/getTitleById", FetchUserTitle)
	user.Get("/updateCNTitle", UpdateCNUserTitle)

	betCtl := demo2.BetLimitCtl{}
	app.PartyFunc("/config", func(r router.Party) {
		//电投限红配置
		r.Post("/limit/insert", betCtl.Insert) // 后台登录
		r.Post("/limit/update", betCtl.Update)
		r.Get("/limit/list", betCtl.List)
		r.Post("/limit/update", betCtl.Update)

		r.Post("/currency/insert", betCtl.Insert) // 后台登录
		r.Post("/currency/update", betCtl.Update)
		r.Get("/currency/list", betCtl.List)
		r.Post("/currency/update", betCtl.Update)
	})
	memberCtl := demo2.UserMember{}
	app.PartyFunc("/member", func(r router.Party) {
		//电投限红配置
		r.Post("/insert", memberCtl.Insert) // 后台登录
		r.Post("/update", memberCtl.Update)
		r.Get("/list", memberCtl.QueryList)
		r.Post("/update", memberCtl.Update)

	})
	currencyCtl := demo2.CtlCurrency{}

	app.PartyFunc("/currency", func(r router.Party) {
		//电投限红配置
		r.Post("/insert", currencyCtl.Insert) // 后台登录
		//r.Post("/update", currencyCtl.Update)
		r.Get("/list", currencyCtl.CurrencyList)
		//r.Post("/update", currencyCtl.Update)

	})

	announce := demo2.AnnouncementCtl{}
	app.PartyFunc("/announce", func(r router.Party) {
		//电投限红配置
		r.Post("/insert", announce.InsertAnnounce) // 后台登录
		r.Post("/update", announce.InsertAnnounce)
		r.Get("/list", announce.QueryAnnounceList)
		r.Post("/update", announce.InsertAnnounce)

	})

	app.Get("/msg", websocket.Handler(websocket_demo.InitWebsocket()))
	app.Get("/push", func(ctx iris.Context) {
		result := consts.Result{
			Message: "socket 消息回复 " + time.Now().GoString(),
			Code:    200,
			Data:    fmt.Sprintf("socket链接的个数 %d", len(websocket_demo.Cmap)),
		}
		message := websocket.Message{
			Body:     result.ToBytes(),
			IsNative: true,
		}
		for key, _ := range websocket_demo.Cmap {
			var conn = websocket_demo.Cmap.GetValue(key)
			conn.Write(message)
		}
		ctx.JSON(result)
	})
	config := iris.WithConfiguration(iris.Configuration{
		DisableStartupLog: true,

		EnableOptimizations: false,
		Charset:             "UTF-8",
	})
	err := app.Run(iris.Addr(":8080"), config)
	if nil == err {
		fmt.Print("start service success")
	} else {
		fmt.Println("err", err)
	}
}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
