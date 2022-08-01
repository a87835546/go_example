package go_rpc_exmaple

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	config := iris.WithConfiguration(iris.Configuration{
		DisableStartupLog: true,

		EnableOptimizations: false,
		Charset:             "UTF-8",
	})
	err := app.Run(iris.Addr(":8081"), config)
	if err != nil {
		panic("start server err:" + err.Error())
	}
}
