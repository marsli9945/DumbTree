package units

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

const jsonRe = `([^\[\]]+)`

func get(r *http.Request) (interface{}, error) {

	err := r.ParseForm()
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(jsonRe)
	result := map[string]interface{}{}

	for k, v := range r.Form {
		matches := re.FindAllStringSubmatch(k, -1)
		matchesLen := len(matches)
		val := result

		if matchesLen < 1 {
			result[k] = v[0]
		} else {
			for i := 0; i < matchesLen; i++ {
				if i == matchesLen-1 {
					val[matches[i][1]] = v[0]
				} else {
					if val[matches[i][1]] == nil {
						val[matches[i][1]] = map[string]interface{}{}
					}
					val = val[matches[i][1]].(map[string]interface{})
				}
			}
		}
	}

	return result, nil
}

func post(r *http.Request) (interface{}, error) {
	result := map[string]interface{}{}
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		err := json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func Params(r *http.Request) (interface{}, error) {
	if strings.Index(r.Header.Get("Content-type"), "application/json") > -1 {
		if r.Method == "POST" || r.Method == "PUT" {
			return post(r)
		}
	}
	return get(r)
}

func Foreach(i interface{}, f func(k string, v interface{})) {
	if m, ok := i.(map[string]interface{}); ok {
		for mk, mv := range m {
			f(mk, mv)
		}
	}
}
