package examples

import (
	"fmt"
	"strconv"
	"time"
)

func Goroutines() {

	//chan1()
	//chanWithBuffer()
	chanWithSelect()

}

func chanWithSelect() {
	// 	c1, c2 := make(chan string), make(chan string)
	//
	// 	go func() {
	// 		for {
	// 			c1 <- "1"
	// 			time.Sleep(time.Millisecond * 500)
	// 		}
	//
	// 	}()
	//
	// 	go func() {
	// 		for {
	// 			c2 <- "2"
	// 			time.Sleep(time.Second * 3)
	// 		}
	// 	}()
	//
	// 	for {
	// 		select {
	// 		case m1 := <-c1:
	// 			fmt.Println(m1)
	// 		case m2 := <-c2:
	// 			fmt.Println(m2)
	// 		}
	// 	}

	c1, c2 := make(chan string), make(chan string)

	go func() {
		for {
			c1 <- "1"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {

			c2 <- "2"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case m1 := <-c1:
			fmt.Println(m1)
		case m2 := <-c2:
			fmt.Println(m2)
		}

	}

}

func chanWithBuffer() {
	ch := make(chan string, 6) // diz o buffer do canal

	ch <- "a"
	ch <- "b"
	ch <- "c"
	ch <- "d"
	ch <- "e"
	ch <- "fc"

	close(ch)

	for v := range ch {

		fmt.Println(v)
	}
}

func chan1() {
	ch := make(chan string)

	write := func(s string, ch chan string) {
		for i := 0; i < 5; i++ {
			ch <- s + " " + strconv.Itoa(i+1)
			time.Sleep(time.Second)
		}

		close(ch)
	}

	go write("Teste", ch)

	fmt.Println("Fim!")

	for v := range ch {
		fmt.Println(v)
	}

}
