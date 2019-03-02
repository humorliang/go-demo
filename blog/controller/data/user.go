package data

//ll_admin数据表
type Admin struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Realname    string `json:"realname"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
	Description string `json:"description"`
	LoginTime   string `json:"login_time"`
	LoginCount  int    `json:"login_count"`
	Status      int    `json:"status"`
	ParentId    int    `json:"parent_id"`
	IsSuper     int    `json:"is_super"`
	IsDelete    int    `json:"is_delete"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}
