package demo2

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"test/db"
)

type CurrencyService struct {
}

func InsertCurrency(record goqu.Record) error {
	sql, _, err := db.G.Insert("tbl_currency").Rows(record).ToSQL()
	if err != nil {
		fmt.Errorf("sql err -- >>> >%v \n", err.Error())
		return err
	} else {
		fmt.Printf("sql ---->>> %v\n", sql)
		_, err = db.Db.Exec(sql)
		return err
	}
}
func QueryCurrencyList() (v []Currency, err error) {
	res, err := db.Db.Queryx("SELECT * FROM tbl_currency")
	for res.Next() {
		var p Currency
		err = res.StructScan(&p)
		v = append(v, p)
	}
	if err != nil {
		fmt.Errorf("err --->>>> %v \n", err.Error())
	}
	return v, err
}
