package main

import "net/http"

func main() {

	http.HandleFunc("/test", handlefunc)

}

func handlefunc(r http.ResponseWriter, rq *http.Request) {

}