package iris

import (
	"github.com/kataras/iris/v12"
	"time"
)

func HomeHandle(ctx iris.Context) {
	ctx.Next()
}
func GetBanner(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": "iris 路由的基本使用",
		"path":    ctx.Path(),
		"time":    time.Now(),
		"data":    []string{"1", "2", "3"},
	})
}

func GetHome(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": "iris 路由的基本使用",
		"path":    ctx.Path(),
		"time":    time.Now(),
		"data":    []string{},
	})
}
