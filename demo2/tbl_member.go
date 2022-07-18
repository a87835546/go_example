package demo2

// Member CMS 管理人员
type Member struct {
	MemberID        int64   `json:"member_id" db:"member_id"`
	SuperMemberID   int64   `json:"super_member_id" db:"super_member_id"`    //代理 id
	CurrencyType    int     `json:"currency" db:"currency"`                  //币种  1 php 菲律宾披索  2. HKD 港币 3.RMB 人民币 4.KRW 韩元
	ShareRatio      float32 `json:"share_ratio" db:"share_ratio"`            //占成比例  0.5 如果占成比 == 0 那么代理的佣金比不能大于我自己的佣金比
	XiMaRatio       float32 `json:"xi_ma_ratio" db:"xi_ma_ratio"`            // 洗码比
	MemberAccount   string  `json:"member_account" db:"member_account"`      //会员号
	MemberName      string  `json:"member_name" db:"member_name"`            //用户名
	MemberPwd       string  `json:"member_pwd" db:"member_pwd"`              //用户密码
	Mobile          string  `json:"mobile" db:"mobile"`                      //电话号码
	MobileBackup    string  `json:"mobile_backup" db:"mobile_backup"`        //备用电话
	AccountType     int     `json:"account_type" db:"account_type"`          //账号类型  1下线代理 2游戏会员  -- 即直属会员
	IsBanAccount    bool    `json:"is_ban_account" db:"ban_account"`         // 是否禁用账号 default false
	IsBanBetAccount bool    `json:"is_ban_bet_account" db:"ban_bet_account"` // 是否禁用游戏，但是还是可以登录
	CommissionRatio float32 `json:"commission_ratio" db:"commission_ratio"`  //上级代理给我的佣金率
	Limit           int     `json:"betting_limit" db:"betting_limit"`        // 限红设置 [id,id,id]
	CreatedAt       int64   `json:"created_at" db:"created_at"`              //开户时间
	UpdatedAt       int64   `json:"updated_at" db:"updated_at"`              //最后登录时间
	PermissionType  int     `json:"permission_type" db:"permission_type"`    //权限 1浏览 2编辑
}
