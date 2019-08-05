package handle

import (
	"DumbTree/dumbError"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

func FileHandle(writer http.ResponseWriter, request *http.Request) error {

	if request.URL.Path == "/favicon.ico" {
		return nil
	}

	if strings.Index(request.URL.Path, prefix) != 0 {
		return dumbError.UserError("path must start with " + prefix)
	}

	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)

	return nil
}
