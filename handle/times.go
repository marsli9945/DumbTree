package handle

import (
	"net/http"
)

func TimeTest(writer http.ResponseWriter, request *http.Request) {
	http.Error(writer, "time test", 305)
	//fmt.Println(writer,"this is times")
}
