package demo2

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type TouchAt struct {
	CreatedAt int64 `json:"created_at" db:"created_at"`
	UpdatedAt int64 `json:"updated_at" db:"updated_at"`
}
type Language struct {
	CN string `json:"cn" db:"cn"`
	KR string `json:"kr" db:"kr"`
	EN string `json:"en" db:"en"`
}

func (l *Language) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		json.Unmarshal(v, &l)
		return nil
	case string:
		json.Unmarshal([]byte(v), &l)
		return nil
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}

}

func (l Language) Value() (driver.Value, error) {
	return json.Marshal(&l)

}
