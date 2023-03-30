package demo2

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"test/db"
)

type MemberService struct {
}
type m = map[string]any

func InsertMember(record m) (err error) {
	sql, _, err := db.G.Insert("tbl_member").Rows(record).ToSQL()
	fmt.Printf("sql --->>>> %v \n\n", sql)
	if err == nil {
		_, err = db.Db.Exec(sql)
		if err != nil {
			fmt.Errorf("err -->>> %v \n", err.Error())
		}
	} else {
		fmt.Errorf("insert memeber err   -->>>> %v \n", err.Error())
	}
	return err
}
func UpdateMember(record m) (err error) {
	ex := goqu.Ex{
		"member_id": record["member_id"],
	}
	sql, _, err := db.G.Update("tbl_member").Set(record).Where(ex).ToSQL()
	fmt.Printf("sql --->>>> %v \n\n", sql)
	if err == nil {
		_, err = db.Db.Exec(sql)
		if err != nil {
			fmt.Errorf("err -->>> %v \n", err.Error())
		}
	} else {
		fmt.Errorf("insert memeber err   -->>>> %v \n", err.Error())
	}
	return err
}
func QueryMembers() (members []Member, err error) {
	res, err := db.Db.Queryx("SELECT * FROM tbl_member")
	if err != nil {
		fmt.Errorf("err --->>> %v \n", err.Error())
	}
	for res.Next() {
		var p Member
		err = res.StructScan(&p)
		members = append(members, p)
		if err != nil {
			fmt.Errorf("parser dataset error --->>> " + err.Error())
		}
	}
	fmt.Printf("list ---- >>>> %v \n", members)
	return members, err
}
