package handle

import (
	"DumbTree/dumbError"
	"DumbTree/units"
	"fmt"
	"net/http"
)

type Collection struct {
	path  []string
	param interface{}
	dns   string
}

func MainHandle(writer http.ResponseWriter, request *http.Request) error {

	if request.URL.Path == "/favicon.ico" {
		return nil
	}

	var coll Collection

	coll.path = units.StringToArr(request.URL.Path, "/")

	if len(coll.path) < 2 || len(coll.path) > 3 {
		return dumbError.UserError("非法路由")
	}

	param, err := units.Params(request)

	if err != nil {
		return err
	}

	coll.param = param

	fmt.Println(coll.param)

	//units.Foreach(coll.param, func(k string, v interface{}) {
	//	fmt.Println(k, v)
	//	units.Foreach(v, func(mk string, mv interface{}) {
	//		fmt.Println(mk,mv)
	//	})
	//	fmt.Println()
	//})

	return nil
}
