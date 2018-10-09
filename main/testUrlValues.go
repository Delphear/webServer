package main

import (
	"fmt"
	"net/url"
)

func main(){
	v := url.Values{}
	v.Set("name","hhaha")
	v.Add("value","https:")
	v.Add("value","www.")
	v.Add("value","baidu.com")
	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("value"))
	fmt.Println(v["value"])
}
