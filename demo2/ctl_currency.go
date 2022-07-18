package demo2

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/kataras/iris/v12"
	"test/consts"
	"time"
)

type CtlCurrency struct{}
type CtlCurrencyVo struct {
	Name  string  `json:"name"`
	Ratio float32 `json:"ratio"`
	Type  int     `json:"type"`
}

func (v *CtlCurrency) Insert(ctx iris.Context) {
	vo := CtlCurrencyVo{}
	ctx.ReadJSON(&vo)
	fmt.Printf("insert currency parameters  --->>> %v \n", vo)
	if err := InsertCurrency(configCurrencyModel(vo)); err == nil {
		ctx.JSON(consts.Result{
			Message: "insert currency success",
			Code:    200,
			Data:    "",
		})
	} else {
		ctx.JSON(consts.Result{
			Message: "insert currency failed" + err.Error(),
			Code:    400,
			Data:    "",
		})
	}
}

func (v *CtlCurrency) CurrencyList(ctx iris.Context) {
	if res, err := QueryCurrencyList(); err == nil {
		ctx.JSON(consts.Result{
			Message: "success query list",
			Code:    200,
			Data:    res,
		})
	} else {
		ctx.JSON(consts.Result{
			Message: "failed query list" + err.Error(),
			Code:    400,
			Data:    res,
		})
	}
}

func configCurrencyModel(vo CtlCurrencyVo) goqu.Record {
	record := goqu.Record{}
	record["desc"] = vo.Name
	record["ratio"] = vo.Ratio
	record["type"] = vo.Type
	record["created_at"] = time.Now().Unix()
	record["updated_at"] = time.Now().Unix()
	return record
}
