package model

type Language struct {
	CN string `json:"cn" db:"cn"`
	EN string `json:"en" db:"en"`
	KR string `json:"kr" db:"kr"`
}

const (
	cn = iota
	en
	kr
)
