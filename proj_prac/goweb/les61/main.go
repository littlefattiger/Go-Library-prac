package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

// func f1(w http.ResponseWriter, r *http.Request) {

// 	k := func(name string) (string, error) {
// 		return name + " young and cool", nil

// 	}
// 	t, err := template.New("f.tmpl").Funcs(template.FuncMap{"kua": k}).ParseFiles("./f.tmpl")
// 	if err != nil {
// 		fmt.Printf("parse template failed, err:%v", err)
// 		return
// 	}
// 	name := "apple"
// 	t.Execute(w, name)

// }

func tmplDemo(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	user := "小王子"
	tmpl.Execute(w, user)
}

func main() {
	// http.HandleFunc("/", f1)
	http.HandleFunc("/tmpl", tmplDemo)

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Http Server start failed, err:%v", err)
		return
	}
}
