package iris

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
)

type Result struct {
	message string
	code    int
	data    any
}
type LoginVo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Service() {
	app := iris.New()
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
			code:    200,
			message: "test4",
			data:    context.PostValue("username"),
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
				"data":    context.PostValue("username"),
			})
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

	example := app.Party("example", ExampleHandle)
	example.Get("/test", ExampleRequest)
	home := app.Party("home", HomeHandle)
	home.Get("/banner", GetBanner)
	home.Get("", GetHome)
	user := app.Party("user", UserHandle)
	user.Get("", GetUser)
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
