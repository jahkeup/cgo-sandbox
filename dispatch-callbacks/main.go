package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("invoking callbacks")
	InvokeCallbacks(&Callbacks{
		One: callbackOne,
		Two: callbackTwo,
		Three: callbackThree,
	})

	fmt.Fprintln(log.Writer()) // separate output

	log.Println("invoking lambda callbacks")
	InvokeCallbacks(&Callbacks{
		One: func() {
			log.Println("lambdaOne called")
		},
		Two: func(x int) {
			log.Printf("lambdaTwo called with x:%d", x)
		},
		Three: func(x int, s string) {
			log.Printf("lambdaThree called with x:%d s:%q", x, s)
		},
	})
}

func callbackOne() {
	log.Println("callbackOne called")
}

func callbackTwo(x int) {
	log.Printf("callbackTwo called, x:%d", x)
}

func callbackThree(x int, s string) {
	log.Printf("callbackThree called, x:%d s:%q", x, s)
}
