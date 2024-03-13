package exercises

import (
	"fmt"
	"time"
)

func Practice() {
	go fmt.Print("a")
	fmt.Print("b")
	fmt.Print("b")
	go fmt.Print("c")
	fmt.Print("b")
	fmt.Print("b")
	time.Sleep(time.Second)
	// bbbbac or bbbbca
}

func PracticeTester() bool {
	Practice()

	return true
}
