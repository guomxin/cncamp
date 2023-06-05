package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http.ListenAndServe(":8088", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	queryParams := r.URL.Query()
	// 	data, _ := json.Marshal(queryParams)
	// 	fmt.Println(string(data))
	// 	w.Write([]byte("你好"))
	// }))

	m := http.NewServeMux()
	// m.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Visit /"))
	// }))
	m.Handle("/aaa", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		name := queryParams.Get("name")
		w.Write([]byte(fmt.Sprintf("Visit /aaa, name=%s", name)))
	}))
	m.Handle("/bbb", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Visit /bbb"))
	}))

	http.ListenAndServe(":8088", m)
}
