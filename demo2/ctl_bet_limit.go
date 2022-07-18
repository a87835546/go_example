package demo2

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/kataras/iris/v12"
	"test/consts"
	"time"
)

type BetLimitCtl struct{}

type BettingLimitVo struct {
	Index         int `json:"index"`
	MinAmount     int `json:"min_amount" db:"min_amount"` //最小金额
	MaxAmount     int `json:"max_amount" db:"max_amount"` // 最大金额
	BetResultType int `json:"bet_type" db:"bet_type"`     //投注的类型， 1 庄/闲  2 和   3 对子
}

func (v *BetLimitCtl) Update(ctx iris.Context) {
	limit := BettingLimit{}
	ctx.ReadJSON(&limit)
	fmt.Printf("limit --->>>> %v \n", limit)
	BetLimitUpdate()

}
func (v *BetLimitCtl) List(ctx iris.Context) {
	if res, err := QueryAllList(); err == nil {
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

func (v *BetLimitCtl) Insert(ctx iris.Context) {
	limit := BettingLimitVo{}
	ctx.ReadJSON(&limit)

	fmt.Printf("covert struct to map --->>>> %v \n", CovertStructToMap(limit))
	fmt.Printf("limit --->>>> %v \n", limit)
	if err := BetLimitInsert(configInsertRd(limit)); err == nil {
		ctx.JSON(consts.Result{
			Message: "success insert",
			Code:    200,
			Data:    "",
		})
	} else {
		ctx.JSON(consts.Result{
			Message: "failed insert" + err.Error(),
			Code:    400,
			Data:    "",
		})
	}

}

func configInsertRd(v BettingLimitVo) goqu.Record {
	rd := goqu.Record{}
	rd["created_at"] = time.Now().Unix()
	rd["updated_at"] = time.Now().Unix()
	rd["bet_type"] = v.BetResultType
	rd["min_amount"] = v.MinAmount
	rd["max_amount"] = v.MaxAmount
	rd["index"] = v.Index
	return rd
}
