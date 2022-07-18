package demo2

type Wallet struct {
	MemberId     string  `json:"member_id" db:"member_id"`
	CurrencyType int     `json:"currency" db:"currency"` //币种  1 php 菲律宾披索  2. HKD 港币 3.RMB 人民币 4.KRW 韩元
	Balance      float32 `json:"balance" db:"balance"`
	CreatedAt    int64   `json:"created_at" db:"created_at"` //开户时间
	UpdatedAt    int64   `json:"updated_at" db:"updated_at"` //最后登录时间
}
