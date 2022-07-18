package demo2

// BettingLimit 投注限红
type BettingLimit struct {
	Id            int64 `json:"id" db:"id"`
	MinAmount     int   `json:"min_amount" db:"min_amount"` //最小金额
	MaxAmount     int   `json:"max_amount" db:"max_amount"` // 最大金额
	BetResultType int   `json:"bet_type" db:"bet_type"`     //投注的类型， 1 庄/闲  2 和   3 对子
	Index         int8  `json:"index" db:"index"`           //三个选项 选项的索引
	CreatedAt     int64 `json:"created_at" db:"created_at"` //开户时间
	UpdatedAt     int64 `json:"updated_at" db:"updated_at"` //最后登录时间
}
