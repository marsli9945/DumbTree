package handle

import (
	"fmt"
	"net/http"
)

func ApiHandle(writer http.ResponseWriter, request *http.Request) error {

	fmt.Println(request.URL.Path)
	return nil
}
