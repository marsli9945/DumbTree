package dumbError

import (
	"log"
	"net/http"
	"os"
)

type UserError string

func (e UserError) Message() string {
	return string(e)
}

func (e UserError) Error() string {
	return e.Message()
}

type appHandle func(writer http.ResponseWriter, request *http.Request) error

/**
错误统一拦截处理模块
*/
func ErrWrapper(handle appHandle) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handle(writer, request)
		if err != nil {

			log.Printf("Error Handling request:%s", err.Error())

			//user error
			if userErr, ok := err.(UserError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			//system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden

			default:
				code = http.StatusInternalServerError
			}

			http.Error(writer, http.StatusText(code), code)
		}
	}
}
