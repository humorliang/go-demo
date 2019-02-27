package ops

import "blog/dbops/gmysql/db"

//ID查询
func UserIsExsit(id int) bool {
	db.Con.QueryRow()
}
