package demo2

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"test/db"
)

type BetLimit struct{}

func BetLimitUpdate() {

}

func QueryAllList() (v []BettingLimit, err error) {
	res, err := db.Db.Queryx("SELECT * FROM tbl_bet_limit")
	for res.Next() {
		var p BettingLimit
		err = res.StructScan(&p)
		v = append(v, p)
	}
	if err != nil {
		fmt.Errorf("err --->>>> %v \n", err.Error())
	}
	return v, err
}
func BetLimitInsert(record goqu.Record) (err error) {
	query, _, _ := db.G.From("tbl_bet_limit").Insert().Rows(record).ToSQL()

	//query, _, err := goqu.Insert("tbl_bet_limit").Rows(record).ToSQL()
	if err != nil {
		fmt.Errorf("config sql err -->>> %v", err.Error())
	}
	fmt.Printf("sql -- >>> %v \n", query)
	_, err = db.Db.Exec(query)
	if err != nil {
		fmt.Errorf("insert sql err -->>> %v", err.Error())
	}
	return
}
