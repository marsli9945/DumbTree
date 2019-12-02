package main

import "fmt"

type node struct {
	val []interface{}
}

func main() {
	result := map[string]interface{}{}
	matches := [][]string{
		{"aa", "dd"},
		{"aa", "ee", "ff"},
	}
	v := "55"

	for _, match := range matches {
		val := result
		for i := 0; i < len(match); i++ {
			if i == len(match)-1 {
				val[match[i]] = v
			} else {
				fmt.Println(val[match[i]])
				if val[match[i]] == nil {
					val[match[i]] = map[string]interface{}{}
				}
				val = val[match[i]].(map[string]interface{})
			}
		}
		fmt.Println(val)
	}

	fmt.Println(result)
}
