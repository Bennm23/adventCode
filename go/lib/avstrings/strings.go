package avstrings

func ParseTextInParens(str string) string {
	var s string

	marked := false
	for _, c := range str {
		if c == ')' {
			break;
		}
		if marked {
			s = s + string(c)
		}
		if c == '(' {
			marked = true
			continue
		}
	}
	
	return s
}