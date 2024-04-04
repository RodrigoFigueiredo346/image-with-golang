package exercises

func LongestCommonPrefix(s []string) string {

	shortest := shortestString(s)
	aux := ""
	ret := ""
	isPrefix := true

	for iS, x := range shortest {
		isPrefix = true
		for _, v := range s {
			if string(x) == string(v[iS]) {
				aux = string(x)
			} else {
				isPrefix = false
			}
		}
		if isPrefix {
			ret += aux
		} else {
			return ret
		}

	}
	return ret
}

func shortestString(str []string) string {
	shortest := str[0]
	for i := range str {
		if len(shortest) > len(str[i]) {
			shortest = str[i]
		}
	}
	return shortest
}
