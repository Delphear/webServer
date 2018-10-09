package main

import (
	"fmt"
	"github.com/Delphear/webServer/common"
	"net/url"
)

func main() {
	v := url.Values{}
	v.Set("name", "hhaha")
	v.Add("value", "https:")
	v.Add("value", "www.")
	v.Add("value", "baidu.com")
	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("value"))
	fmt.Println(v["value"])

	targetUrl := "http://127.0.0.1:9090/upload"
	fileName := "README.md"
	common.PostFile(fileName, targetUrl)
}
