package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)
type Post struct {
	User string
	Threads []string
}
func writeExample (writer http.ResponseWriter, request *http.Request ) {
	str:=`<html>
<head><title><GoWeb</title></head>
<body><h1>Hellow</h1></body>
</html>`
  writer.Write([]byte(str))
}

func writeHeaderExample (writer http.ResponseWriter, request *http.Request ) {
	writer.WriteHeader(501)
	fmt.Fprintf(writer, "no such server")
}
func headerExample (writer http.ResponseWriter, request *http.Request ) {
	writer.Header().Set("Location","https://www.baidu.com")
	writer.WriteHeader(302)
}

func jsonExample (writer http.ResponseWriter, request *http.Request ) {
	writer.Header().Set("Content-Type","application/json")
	post :=&Post{
		User:"MrZhaoCn",
		Threads: []string{"first","second"},
	}
	json,_ :=json.Marshal(post)
	writer.Write(json)
}
func main()  {
	server :=http.Server{Addr:"127.0.0.1:8081"}
	http.HandleFunc("/write",writeExample)
	http.HandleFunc("/writehead",writeHeaderExample)
	http.HandleFunc("/redirect",headerExample)
	http.HandleFunc("/json",jsonExample)
	server.ListenAndServe()
}