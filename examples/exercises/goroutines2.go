package examples

import "fmt"

func Goroutines2() {

	chanWithSelect2()
}

func chanWithSelect2() {
	code := 19991
	char := rune(code)
	s := []rune{'A', '世'}

	fmt.Printf("o código %d representa caracter %c \n", code, char)
	fmt.Println(s)

}
