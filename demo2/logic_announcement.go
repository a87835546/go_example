package demo2

import (
	"fmt"
	"test/db"
)

type AnnouncementService struct {
}

func AnnounceInsert(mode AnnouncementMode) {
	sql, _, err := db.G.Insert("tbl_announcement").Rows(mode).ToSQL()
	if err != nil {
		fmt.Errorf("sql err -->>> %v \n", err.Error())
	}
	_, err = db.Db1.Exec(sql)
	fmt.Printf("sql --->>> >%v\n", sql)
	if err != nil {
		fmt.Errorf("exec sql err -->>> %v \n", err.Error())
	}
}
func AnnounceList() (list []AnnouncementMode, err error) {
	rows, err := db.Db1.Queryx("SELECT * FROM tbl_announcement")
	//err = db.Db1.Select(&list, "SELECT * FROM tbl_announcement")
	if err != nil {
		fmt.Errorf("exec sql err -->>> %v \n", err.Error())
		return list, err
	}
	p := AnnouncementMode{}
	//p := map[string]any{}
	for rows.Next() {
		err := rows.StructScan(p)
		fmt.Printf("res -->>> %v \n", p)
		list = append(list, p)
		if err != nil {
			fmt.Errorf("model pareser err -->>> %v \n", err.Error())
		}
	}
	return list, nil
}
