package examples

import (
	"fmt"
	"os"
	"strings"
)

func SumDigits2() int {

	file, err := os.ReadFile("SumDigits2.txt")
	if err != nil {
		fmt.Printf("Erro ao ler o arquivo: %v", err)
		return 0
	}

	rows := strings.Split(string(file), "\n")

	digits := ""
	var listDigits []int
	ret := 0

	nums := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for _, v := range rows { //"oneygffugietwodjfashl"
		digits = ""
		for key, val := range nums { // "one": 1
			if strings.Contains(v, key) {
				digits += string(rune(val))
				fmt.Println("index:  ", strings.Index(v, key), " - ", key)
			}
		}
		aux := digits[0] + digits[len(digits)-1]
		listDigits = append(listDigits, int(aux))
		fmt.Println(listDigits)
	}

	for _, value := range listDigits {
		ret += value
	}

	return ret
}
