package iris

import (
	"github.com/kataras/iris/v12"
	"time"
)

func ExampleHandle(ctx iris.Context) {
	ctx.Next()
}

func ExampleRequest(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"time": time.Now(),
		"path": ctx.Path(),
	})
}
