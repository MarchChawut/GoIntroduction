package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	//ดัก path
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home, %q", html.EscapeString(r.URL.Path))
	})

	//ดัก path แล้วดึงค่า username และ password ออกมาแสดง
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Login, %s, %s", r.URL.Query().Get("username"), r.URL.Query().Get("password"))
	})
	//http://localhost:85/login?username=asdf&&password=1234

	log.Fatal(http.ListenAndServe(":85", nil))
}

