package handle

import (
	"net/http"
	"strings"
)

/**
处理静态资源文件
*/
func VueHandle(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path
	if strings.Index(request.URL.Path, "/vue/") == 0 {
		path = request.URL.Path[len("/vue/"):]
		if path == "" {
			path = "/index.html"
		}
	}

	path = "vue/dist" + path

	http.ServeFile(writer, request, path)

	return nil
}
