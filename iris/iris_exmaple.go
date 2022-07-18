package iris

import (
	"github.com/kataras/iris/v12"
	"test/consts"
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
	ctx.JSON(consts.Result{
		Message: ctx.Path(),
		Code:    200,
		Data:    time.Now(),
	})
}

func Query(ctx iris.Context) {
	if id, err := ctx.URLParamInt("id"); err == nil {
		ctx.JSON(consts.Result{
			Message: ctx.Path(),
			Code:    200,
			Data:    db.QueryData(id),
		})
	} else {
		ctx.JSON(consts.Result{
			Message: ctx.Path(),
			Code:    500,
			Data:    "parameter is error " + err.Error(),
		})
	}

}
