package handle

import (
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

	fmt.Println(request.Method)

	params := units.Params(request)

	units.Foreach(params, func(k string, v interface{}) {
		fmt.Println(k, v)
	})

	//fmt.Println(units.Get(request))
	//fmt.Println(units.Post(request))
	//fmt.Println(units.Post(request))
	//param := units.Post(request)
	//fmt.Println(param["id"])
	//fmt.Println(param["name"].(string))
	//fmt.Println(param["group"])
	//group := param["group"].(map[string]interface{})
	//fmt.Println(group["aa"])
	//fmt.Println(group["bb"])
	//fmt.Println(group["bb"].(map[string]interface{})["cc"].(float64))

	//fmt.Println(request.Header)
	//request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	//fmt.Println(request.URL.Query())
	//fmt.Println(request.ParseForm())

	//fmt.Println(request.Method)
	//fmt.Println(request.URL.Path)
	return nil
}
