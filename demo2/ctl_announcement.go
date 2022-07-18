package demo2

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"test/consts"
	"time"
)

type AnnouncementCtl struct {
}

func (v *AnnouncementCtl) InsertAnnounce(ctx iris.Context) {
	model := AnnouncementMode{}
	ctx.ReadJSON(&model)
	fmt.Printf("model --->> %v\n", model)
	model.CreatedAt = time.Now().Unix()
	model.UpdatedAt = time.Now().Unix()
	model.Index = time.Now().Second()
	model.PublishName = "zhansan"
	AnnounceInsert(model)
	ctx.JSON(consts.Result{
		"insert announce",
		200,
		"",
	})
}

func (v *AnnouncementCtl) QueryAnnounceList(ctx iris.Context) {
	if res, err := AnnounceList(); err == nil {
		ctx.JSON(consts.Result{
			Message: "query  announce successful",
			Code:    200,
			Data:    res,
		})
	} else {
		ctx.JSON(consts.Result{
			Message: "query  announce failed",
			Code:    400,
		})
	}
}
