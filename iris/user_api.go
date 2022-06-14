package iris

import (
	"github.com/kataras/iris/v12"
	"time"
)

func UserHandle(ctx iris.Context) {
	ctx.Next()
}

func GetUser(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": "iris 路由的基本使用",
		"path":    ctx.Path(),
		"time":    time.Now(),
		"data":    []string{},
	})
}
