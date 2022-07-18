package demo2

type Currency struct {
	Id        string  `json:"id" db:"id"`
	Type      int     `json:"type" db:"type"`             //币种  1 php 菲律宾披索  2. HKD 港币 3.RMB 人民币 4.KRW 韩元 5 USDT
	Ratio     float32 `json:"ratio" db:"ratio"`           //兑换披索的汇率
	Desc      string  `json:"desc" db:"desc"`             //币种的描述
	CreatedAt int64   `json:"created_at" db:"created_at"` //开户时间
	UpdatedAt int64   `json:"updated_at" db:"updated_at"` //最后登录时间
}
