package iris

import (
	"github.com/kataras/iris/v12"
	"test/db"
	"time"
)

type RegisterVo struct {
	Name     string `json:"username"`
	Password string `json:"password"`
	Type     int    `json:"type"`
}

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

func RegisterUser(ctx iris.Context) {
	register := RegisterVo{}
	if err := ctx.ReadJSON(&register); nil != err {
		ctx.JSON(Result{
			"parameter error" + err.Error(),
			500,
			nil,
		})
	} else {
		user := db.InsertData(10, register.Name)
		ctx.JSON(iris.Map{
			"message": "iris 路由的基本使用",
			"path":    ctx.Path(),
			"time":    time.Now(),
			"data":    user,
		})
	}
}
