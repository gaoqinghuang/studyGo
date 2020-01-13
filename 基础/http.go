package main

import (
	"net/http"
)

func main() {
	//获取域名的
	//ctx.Request.Host
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8081", nil)

}
