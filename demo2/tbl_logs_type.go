package demo2

// LogsTypeModel 操作日志的类型
type LogsTypeModel struct {
	Id       string `json:"id" db:"id"`
	Type     string `json:"type" db:"type"`         // 类型
	Desc     string `json:"desc" db:"desc"`         // 类型描述
	Category string `json:"category" db:"category"` // 类型所属分类
}
