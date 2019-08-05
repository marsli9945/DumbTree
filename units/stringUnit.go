package units

/**
字符串转数组
*/
func StringToArr(str string, pix string) []string {
	s := ""
	var arr []string

	for i, ch := range str {
		if ch == 47 {
			if s != "" {
				arr = append(arr, s)
			}
			s = ""
		} else {
			s += string(ch)
		}

		if i == len(str)-1 {
			arr = append(arr, s)
		}
	}

	return arr
}
