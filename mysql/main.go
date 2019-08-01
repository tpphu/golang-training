package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := sql.Open("mysql", "default:secret@/code_service")
	result, _ := db.Exec("insert into code(`min`,`max`) " +
		" select 21, 30 " +
		" from code " +
		" where 0 = (select count(id) " +
		" from code " +
		" where (21 <= `max`) AND 30 >= `min`) " +
		" LIMIT 1")

	rowAffect, _ := result.RowsAffected()
	fmt.Println(rowAffect)
}
