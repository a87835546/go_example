package iris

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/kataras/iris/v12"
	"test/db"
	"time"
)

type RegisterVo struct {
	Name     string  `json:"username"`
	Password string  `json:"password"`
	Type     int     `json:"type"`
	Desc     db.Desc `json:"desc"`
}

type UpdateVo struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
}

func UserHandle(ctx iris.Context) {
	ctx.Next()
}

func GetUser(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": "iris 路由的基本使用",
		"path":    ctx.Path(),
		"time":    time.Now(),
		"data":    db.QueryAll(),
	})
}

// RegisterUser ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  RegisterVo
// @Failure      400  {object}  RegisterVo
// @Failure      404  {object}	RegisterVo
// @Failure      500  {object}	demo1.Service
// @Router       /accounts/{id} [get]
func RegisterUser(ctx iris.Context) {
	register := RegisterVo{}
	if err := ctx.ReadJSON(&register); nil != err {
		ctx.JSON(Result{
			"parameter error-->>" + err.Error(),
			500,
			ctx.RemoteAddr(),
		})
	} else {
		fmt.Printf("register post data ---->>>>%s\n", register)
		user := db.InsertData(10, register.Name)
		ctx.JSON(iris.Map{
			"message": "iris 路由的基本使用",
			"path":    ctx.Path(),
			"time":    time.Now(),
			"data":    user,
		})
	}
}

func QueryAllUsers(ctx iris.Context) {
	ctx.JSON(Result{
		"query all users data",
		200,
		db.QueryAllUsers(),
	})
}
func QueryUsersById(ctx iris.Context) {
	id, _ := ctx.URLParamInt("id")
	ctx.JSON(Result{
		"query all users data",
		200,
		db.QueryById(id),
	})
}

func UpdateUser(ctx iris.Context) {
	age, _ := ctx.PostValueInt64("age")
	id, _ := ctx.PostValueInt64("id")
	ctx.JSON(Result{
		"query all users data",
		200,
		db.UpdateUserById(age, id),
	})
}
func UpdateUserTitle(ctx iris.Context) {
	vo := UpdateVo{}
	var err error
	if err = ctx.ReadJSON(&vo); err == nil {
		desc := map[string]interface{}{}
		fmt.Println("desc ---- >>>> %a  title--->>>>", desc, vo.Title)
		if err = json.Unmarshal([]byte(vo.Title), &desc); err == nil {

			ctx.JSON(Result{
				"update user data",
				200,
				db.UpdateUserTitleById(desc, vo.Id),
			})
		}
	}
	ctx.JSON(Result{
		"parameter error--->>>" + err.Error(),
		500,
		nil,
	})

}

func UpdateUserTitleByGoqu(ctx iris.Context) {
	vo := UpdateVo{}
	var err error
	if err = ctx.ReadJSON(&vo); err == nil {
		if res, err1 := jsoniter.MarshalToString(vo.Title); err1 == nil {
			db.UpdateUserTitle(res, int(vo.Id))
			ctx.JSON(Result{
				"update user data",
				200,
				time.Now(),
			})
			return
		} else {
			ctx.JSON(Result{
				"parameter error--->>>" + err1.Error(),
				501,
				nil,
			})
			return
		}

	}

	ctx.JSON(Result{
		"parameter error--->>>" + err.Error(),
		500,
		nil,
	})

}

func FetchUserTitle(ctx iris.Context) {
	id, _ := ctx.URLParamInt("id")

	ctx.JSON(Result{
		"select user title by id",
		200,
		db.FetchTitleById(id),
	})
}

func UpdateCNUserTitle(ctx iris.Context) {
	id, _ := ctx.URLParamInt("id")
	title := ctx.URLParam("cn")
	index, _ := ctx.URLParamInt("type")

	ctx.JSON(Result{
		"select user title by id",
		200,
		db.UpdateTitleById(id, title, index),
	})
}
