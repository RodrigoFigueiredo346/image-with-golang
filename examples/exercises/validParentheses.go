package examples

func IsValid(s string) bool {

	var stack []rune

	if len(s) == 1 {
		return false
	}
	if len(s)%2 == 1 {
		return false
	}

	for _, v := range s {
		switch v {
		case '{', '[', '(':
			stack = append(stack, v)
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		}

	}

	return len(stack) == 0
}
