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

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	msg := "小王子"
	tmpl.Execute(w, msg)
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	msg := "小王子"
	tmpl.ExecuteTemplate(w, "index2.tmpl", msg)
}

func index2(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/base.tmpl", "./templates/index2.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	msg := "小王子"
	tmpl.ExecuteTemplate(w, "home2.tmpl", msg)
}

func home2(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/base.tmpl", "./templates/home2.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	msg := "小王子"
	tmpl.Execute(w, msg)
}
func main() {
	// http.HandleFunc("/", f1)
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Http Server start failed, err:%v", err)
		return
	}
}
