package main

import "net/http"

func main() {

	http.HandleFunc("/test", handlefunc)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func handlefunc(rs http.ResponseWriter, rq *http.Request) {

}
