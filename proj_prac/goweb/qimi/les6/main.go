package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func f1(w http.ResponseWriter, r *http.Request) {

	k := func(name string) (string, error) {
		return name + " young and cool", nil

	}
	t, err := template.New("f.tmpl").Funcs(template.FuncMap{"kua": k}).ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v", err)
		return
	}
	name := "apple"
	t.Execute(w, name)

}

func main() {
	http.HandleFunc("/", f1)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Http Server start failed, err:%v", err)
		return
	}
}
