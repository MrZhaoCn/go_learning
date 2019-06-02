package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
//type DbWorker struct {
//	username string
//	password string
//	dbName string
//}
func main()  {
	dbw := DbWorker{
		username:"root",
		password:"06081991lp",
		dbName:"mygodatabase_one",
	}
	Dsn := dbw.username + ":" + dbw.password +  "@tcp(127.0.0.1:3306)/" + dbw.dbName
	db, err := sql.Open("mysql",Dsn)
	if err != nil {
		panic(err)
		return
	}
	_, err = db.Exec(`INSERT INTO godemo_one (id,name, phone) VALUES (2,"MrZhaoCn", "15068189675")`)
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	var name string
	err = db.QueryRow("select name from godemo_one where id = ?", 1).Scan(&name)
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	fmt.Println(name)
	defer db.Close()
}


