package main

import (
	"DumbTree/dumbError"
	"DumbTree/handle"
	"net/http"
	_ "net/http/pprof"
)

func main() {

	//主路由
	http.HandleFunc("/", dumbError.ErrWrapper(handle.MainHandle))

	//后台管理API
	http.HandleFunc("/api/", dumbError.ErrWrapper(handle.ApiHandle))

	//静态页面和资源
	http.HandleFunc("/vue/", dumbError.ErrWrapper(handle.VueHandle))
	http.HandleFunc("/js/", dumbError.ErrWrapper(handle.VueHandle))
	http.HandleFunc("/css/", dumbError.ErrWrapper(handle.VueHandle))
	http.HandleFunc("/img/", dumbError.ErrWrapper(handle.VueHandle))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
