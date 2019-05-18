package main

import (
	"fmt"
	"net/http"
)
type HelloHandle struct {

}
type WorldHandle struct {

}
func (h * HelloHandle) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hellow")
}
func (h * WorldHandle) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "World")
}
func main()  {
	hello := HelloHandle{}
	world := WorldHandle{}
	server :=http.Server{Addr:"127.0.0.1:8081"}
	http.Handle("/hello",&hello)
	http.Handle("/world",&world)
	server.ListenAndServe()
}
