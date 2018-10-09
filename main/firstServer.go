package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func sayHelloName(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fmt.Println("values",r.PostForm)
	fmt.Println("form",r.Form)
	fmt.Println("path :",r.URL.Path)
	fmt.Println("scheme :",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v := range r.Form{
		fmt.Println("k :",k)
		fmt.Println("v :",strings.Join(v,""))
	}
	fmt.Fprintf(w,"hello")

}
func login(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fmt.Println("method",r.Method)
	if r.Method == "GET" {
		t,_ := template.ParseFiles("pages/login.html")
		t.Execute(w,nil)
	}else {
		username := r.Form["username"]
		password := r.Form["password"]
		fmt.Println("username : ",username)
		fmt.Println("password : ",password)
		if  username[0] == "hh" && password[0] == "aa"{
			t,_ := template.ParseFiles("pages/index.html")
			t.Execute(w,nil)
		}else {
			t,_ := template.ParseFiles("pages/error.html")
			t.Execute(w,nil)
		}
	}
	/**
	下拉菜单功能
	*/
	slice:=[]string{"apple","pear","banane"}
	for _, v := range slice {
		if v == r.Form.Get("fruit") {
			fmt.Println("fruit",v)
		}
	}
	/**
	单选按钮
	 */
	slice1 := []string{"1","2"}
	for _,v := range slice1{
		if v == r.Form.Get("gender") {
			fmt.Println("gender",v)
		}
	}
	/**
	复选框
	 */
	 slice2 := []string{"football","basketball","tennis"}
	 a := contains(slice2,r.Form["interest"])
	if a {
		fmt.Println("it is true")
	}else {
		fmt.Println("it is false")
	}

}

func main(){
	mux := &myMux{}
	err := http.ListenAndServe(":9090",mux)
	if err != nil {
		log.Fatal(err)
	}
}
type myMux struct {
}

func (p *myMux)ServeHTTP(w http.ResponseWriter,r *http.Request){
	if r.URL.Path == "/" {
		sayHelloName(w,r)
		return
	}else if r.URL.Path == "/login" {
		login(w,r)
		return
	}else if r.URL.Path == "/upload" {
		upload(w,r)
		return
	}
	http.NotFound(w,r)
	return
}
func upload(w http.ResponseWriter,r *http.Request){
	if r.Method == "GET" {
		t,_ := template.ParseFiles("pages/upload.html")
		t.Execute(w,nil)
	}else {
		r.ParseMultipartForm(32 << 20)
		file,handler,err := r.FormFile("uploadfile")
		if err != nil {
			log.Fatal(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w,"%v",handler.Header)
		f,err := os.Open("./pages/" + handler.Filename)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer f.Close()
		io.Copy(f,file)
	}
}

func contains(slice1,slice2 []string) bool{
	m := make(map[interface{}]int)
	err := true
	for _,v := range slice1 {
		m[v] ++
	}
	for _,v := range slice2 {
		if m[v] <= 0 {
			err = false
			return err
		}
	}
	return err
}