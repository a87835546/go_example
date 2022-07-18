package demo2

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"reflect"
	"test/consts"
	"time"
)

type UserMember struct{}
type UserMemberPin struct {
	MemberID        int64   `json:"member_id" db:"member_id"`
	SuperMemberID   int64   `json:"super_member_id" db:"super_member_id"`   //代理 id
	CurrencyType    int     `json:"currency" db:"currency"`                 //币种  1 php 菲律宾披索  2. HKD 港币 3.RMB 人民币 4.KRW 韩元
	ShareRatio      float32 `json:"share_ratio" db:"share_ratio"`           //占成比例  0.5 如果占成比 == 0 那么代理的佣金比不能大于我自己的佣金比
	XiMaRatio       float32 `json:"xi_ma_ratio" db:"xi_ma_ratio"`           // 洗码比
	MemberAccount   string  `json:"member_account" db:"member_account"`     //会员号
	MemberName      string  `json:"member_name" db:"member_name"`           //用户名
	NickName        string  `json:"member_name" db:"member_name"`           //用户名
	MemberPwd       string  `json:"member_pwd" db:"member_pwd"`             //用户密码
	Mobile          string  `json:"mobile" db:"mobile"`                     //电话号码
	MobileBackup    string  `json:"mobile_backup" db:"mobile_backup"`       //备用电话
	AccountType     int     `json:"account_type" db:"account_type"`         //账号类型  1下线代理 2游戏会员  -- 即直属会员
	IsBanAccount    bool    `json:"ban_account" db:"ban_account"`           // 是否禁用账号 default false
	IsBanBetAccount bool    `json:"ban_account" db:"ban_bet_account"`       // 是否禁用游戏，但是还是可以登录
	CommissionRatio float32 `json:"commission_ratio" db:"commission_ratio"` //上级代理给我的佣金率
	Limit           int     `json:"betting_limit" db:"betting_limit"`       // 限红设置 index
	PermissionType  int     `json:"permission_type" db:"permission_type"`   //权限 1浏览 2编辑
}

func (v *UserMember) QueryList(ctx iris.Context) {
	if res, err := QueryMembers(); err == nil {
		ctx.JSON(consts.Result{
			Message: "err --->> ",
			Code:    200,
			Data:    res,
		})
	} else {
		ctx.JSON(consts.Result{
			Message: "err --->> " + err.Error(),
			Code:    400,
			Data:    res,
		})
	}

}
func (v *UserMember) Insert(ctx iris.Context) {
	p := UserMemberPin{}
	ctx.ReadJSON(&p)
	p.MemberID = time.Now().Unix()

	value := make(map[string]any)
	value = CovertStructToMap(p)
	value["created_at"] = time.Now().Unix()
	value["updated_at"] = time.Now().Unix()
	err := InsertMember(value)
	if err != nil {
		fmt.Errorf("err --- >> >%v \n", err.Error())
		ctx.JSON(consts.Result{
			Message: "err --->> " + err.Error(),
			Code:    400,
			Data:    nil,
		})
	} else {
		ctx.JSON(consts.Result{
			Message: "success --->> ",
			Code:    200,
			Data:    nil,
		})
	}

}
func (v *UserMember) Update(ctx iris.Context) {
	p := UserMemberPin{}
	ctx.ReadJSON(&p)
	value := make(map[string]any)
	value = CovertStructToMap(p)
	value["updated_at"] = time.Now().Unix()
	err := UpdateMember(value)
	if err != nil {
		fmt.Errorf("err --- >> >%v \n", err.Error())
		ctx.JSON(consts.Result{
			Message: "update err --->> " + err.Error(),
			Code:    400,
			Data:    nil,
		})
	} else {
		ctx.JSON(consts.Result{
			Message: "update success --->> ",
			Code:    200,
			Data:    nil,
		})
	}
}

func CovertStructToMap(a any) (res map[string]any) {
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	m := make(map[string]any)
	for i := 0; i < t.NumField(); i++ {
		tagName := t.Field(i).Tag.Get("json")
		fmt.Printf("%s --- %v\n", tagName, v.Field(i).Interface())
		m[tagName] = v.Field(i).Interface()
		if v.Field(i).Kind() == reflect.Slice {
			t2 := v.Field(i).Type()
			v2 := v.Field(i)
			fmt.Printf("t2 --->>> %s", t2.String())
			//m[t.Field(i).Name] = v.Field(i).Interface()
			tagName = t.Field(i).Tag.Get("json")
			m[tagName] = v2.Field(i).Interface()
			for j := 0; j < v2.Len(); j++ {
				v3 := v2.Index(j)
				for k := 0; k < v3.NumField(); k++ {
					fmt.Printf("%s --- %s\n", v3.Type(), v3.Field(k).Interface())
					tagName = t.Field(i).Tag.Get("json")
					m[tagName] = v3.Field(k).Interface()

				}
			}
		}
	}
	return m
}
