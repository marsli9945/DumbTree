package units

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func get(r *http.Request) interface{} {
	r.ParseForm()
	return r.Form
}

func post(r *http.Request) interface{} {
	var result map[string]interface{}
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		json.Unmarshal(body, &result)
	}
	return result
}

func Params(r *http.Request) interface{} {
	if strings.Index(r.Header.Get("Content-type"), "application/json") > -1 {
		if r.Method == "POST" || r.Method == "PUT" {
			return post(r)
		}
	}
	return get(r)
}

func Foreach(i interface{}, f func(k string, v interface{})) {
	m, ok := i.(map[string]interface{})
	if ok {
		for mk, mv := range m {
			f(mk, mv)
		}
	}
}
