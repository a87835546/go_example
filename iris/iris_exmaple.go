package iris

import (
	"github.com/kataras/iris/v12"
	"test/db"
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

func ResultCommonRequest(ctx iris.Context) {
	ctx.JSON(Result{
		ctx.Path(),
		200,
		time.Now(),
	})
}

func Query(ctx iris.Context) {
	if id, err := ctx.URLParamInt("id"); err == nil {
		ctx.JSON(Result{
			ctx.Path(),
			200,
			db.QueryData(id),
		})
	} else {
		ctx.JSON(Result{
			ctx.Path(),
			500,
			"parameter is error " + err.Error(),
		})
	}

}
