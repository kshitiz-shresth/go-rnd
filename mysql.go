package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"


db, err := sql.Open("mysql", "root:@/rtw_jolfas")

