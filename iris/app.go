package iris

import (
	"encoding/json"
	"fmt"
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"io"
	"os"
	_ "test/docs"
)

type Result struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    any    `json:"data"`
}
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

		result := Result{
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
			context.JSON(Result{
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
				context.JSON(Result{
					"upload file failed" + err.Error(),
					201,
					"11",
				})
				return
			}
			context.JSON(Result{
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
