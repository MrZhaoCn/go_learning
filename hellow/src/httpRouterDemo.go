package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func hello(writer http.ResponseWriter, request *http.Request, p httprouter.Params)  {
	fmt.Fprintf(writer, "%s",p.ByName("name"))
}
func main()  {
	mux := httprouter.New()
	mux.GET("/hellow/:name",hello)
	server :=http.Server{Addr:"127.0.0.1:8081", Handler:mux}
	server.ListenAndServe()
}