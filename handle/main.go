package handle

import (
	"DumbTree/controller/main"
	"DumbTree/dumbError"
	"DumbTree/units"
	"fmt"
	"net/http"
)

func MainHandle(writer http.ResponseWriter, request *http.Request) error {

	if request.URL.Path == "/favicon.ico" {
		return nil
	}

	pathArr := units.StringToArr(request.URL.Path, "/")

	fmt.Println(pathArr)

	if len(pathArr) < 2 || len(pathArr) > 3 {
		return dumbError.UserError("非法路由")
	}

	switch request.Method {
	case "get":
		controller.Get()
	case "post":
		controller.Insert()

	}

	//fmt.Println(request.Method)
	//fmt.Println(request.URL.Path)
	return nil
}
