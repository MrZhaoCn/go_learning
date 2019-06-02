package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"path"
	"strconv"
)
var Db *sql.DB
type DbWorker struct {
	username string
	password string
	dbName string
}
type User struct {
	id int `json:"id"`
	name string `json:"name"`
	phone string `json:"phone"`
}

func main()  {
	dbw := DbWorker{
		username:"root",
		password:"06081991lp",
		dbName:"mygodatabase_one",
	}
	Dsn := dbw.username + ":" + dbw.password +  "@tcp(127.0.0.1:3306)/" + dbw.dbName
	var err error
	Db, err = sql.Open("mysql",Dsn)
	defer Db.Close()
	if err != nil {
		panic(err)
		return
	}
	server := http.Server{Addr:"127.0.0.1:8080"}
	http.HandleFunc("/user/",handleRequest)
	server.ListenAndServe()
}
func handleRequest(w http.ResponseWriter, r *http.Request)  {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w,r)
	case "POST":
		err = handlePost(w,r)
	case "DELETE":
		err = handleDelete(w,r)
	default:
		fmt.Printf("data error: %v\n", err)
	}
}
func handleGet (w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {return }
	user, err := Retrieve(id)
	if err != nil {return }
	output, err := json.MarshalIndent(&user,"","\t\t")
	if err != nil {return }
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	return
}
func handlePost (w http.ResponseWriter, r *http.Request) (err error) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		r.Body.Close()
		log.Fatal(err)
	}
	fmt.Println(user.name)
	err = user.create(user.name,user.phone)
	if err != nil {return }
	w.WriteHeader(200)
	return
}
func handleDelete (w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {return }
	user, err := Retrieve(id)
	if err != nil {return }
	err = user.delete()
	if err != nil {return }
	w.WriteHeader(200)
	return
}
func Retrieve(id int) (user User, err error)  {
	user = User{}
	err = Db.QueryRow("select id,name,phone from godemo_one where id = $1", id).Scan(&user.id,&user.name,&user.phone)
	if err != nil {
		fmt.Printf("select data error: %v\n", err)
		return
	}
	return user,err
}
func (user *User) create(name string, phone string) (err error) {
	_, err = Db.Exec(`insert into godemo_one (name, phone) values (name, phone)`)
	fmt.Println(name + phone)
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return err
	}
	return err
}
func (user *User) update() (err error) {
	_, err = Db.Exec(`update godemo_one set name = $2, phone = $3 where id=$1`,user.id,user.name,user.phone)
	if err != nil {
		fmt.Printf("update data error: %v\n", err)
		return err
	}
	return err
}
func (user *User) delete() (err error) {
	_, err = Db.Exec(`delete from godemo_one where id=$1`,user.id)
	if err != nil {
		fmt.Printf("delete data error: %v\n", err)
		return err
	}
	return err
}

