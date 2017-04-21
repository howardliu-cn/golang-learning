package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func main() {
	db, err := sql.Open("sqlite3", "db/sqlite/test.db")
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) VALUES (?, ?, ?)")
	checkErr(err)

	res, err := stmt.Exec("u1", "p1", time.Now().Format("2006-01-02 15:04:05"))
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	stmt, err = db.Prepare("UPDATE userinfo SET username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("u2", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, username, department, created)
	}

	stmt, err = db.Prepare("DELETE FROM userinfo WHERE uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	defer db.Close()
	defer stmt.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
