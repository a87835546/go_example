package demo2

// LogsModel 操作日志数据库
type LogsModel struct {
	OperationAccount string `json:"operation_account" db:"operation_account"` //操作账号
	OperationType    int    `json:"operation_type" db:"type"`                 //操作类型  tbl_logs_type id，创建下线账号  创建子账号
	Account          string `json:"account" db:"account"`                     //对应的账号
	Ip               string `json:"ip" db:"ip"`
	Notice           string `json:"notice" db:"notice"` //备注
	TouchAt
}
