package main

import (
	"DumbTree/handle"
	"DumbTree/router"
	"DumbTree/router/middleware"
	"net/http"
)

func main() {
	router := router.NewRouter()
	router.Use(middleware.Times)
	router.Add("/time", handle.TimeTest)

	router.Use(middleware.Test1)
	router.Use(middleware.Test2)
	router.Add("/test", handle.TimeTest)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//
	//})
	//
	//http.Handle("/", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
	//
	//}))

	//arr := []string{}
	//
	//arr = append(arr, "aa")
	//arr = append(arr, "bb")
	//
	//for i := len(arr) - 1; i >= 0; i-- {
	//	fmt.Println(arr[i])
	//}
}
