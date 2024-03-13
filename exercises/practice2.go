package exercises

import (
	"fmt"
)

func TopKFrequent(nums []int, k int) []int {
	// map to count the frequency of each number
	// freqMap := make(map[int]int) // Map to count the frequency of each number.
	// for _, num := range nums {
	// 	freqMap[num]++ // Increment the count for the number.
	// }
	// create a queue
	// add each number from the map to the queue
	// A := 100
	// B := &A // 0x1400000e098
	// C := *B // 100
	// D := B  // 0x1400000e098
	// fmt.Print(C, D)

	// go printA()
	// printB()
	// printC()
	// printD()
	// for i := 0; i < 10; i++ {
	// 	fmt.Print(i)
	// }

	// f("direct")
	// go f("goroutine")
	// go going("going")
	// time.Sleep(time.Second) // pause synchronous execution for 1 second to wait for asynchronous goroutines

	// go fmt.Print("a") // asynchronous print a
	// fmt.Print("b")  // synchrounously print b
	// time.Sleep(time.Second) // pause synchronous execution for 1 second to wait for asynchronous goroutines
	//prints ba

	// create a channel called "messages" that moves a string
	// messages := make(chan string, 1) // 1 = maximum number of inputs into the channel
	// // call a goroutine that sends a "ping" string to the channel
	// go func() {
	// 	messages <- "ping"
	// }()
	// // receive the output of the channel in a variable called "msg"
	// go func() {
	// 	output := <-messages
	// 	fmt.Println(output) // "ping"
	// }()

	// output := <-messages
	// fmt.Println(output)

	// done := make(chan bool, 1)
	// // call the function above as a goroutine
	// go worker(done)
	// // receive the value from the channel
	// // without this, the sender will not send its message
	// <-done

	// 'i' is an interface that holds a string. The interface could potentially hold any type
	// var i interface{} = "hello"
	// value, ok := i.(string) // This is a valid type assertion.
	// if ok {
	// 	output := fmt.Sprintf("output: %s", value)
	// 	fmt.Print(output)
	// }

	// // example 1: non-slices are stored by value so when you update the value of a, b is not changed
	// a := 1
	// b := a
	// a = 2
	// fmt.Print(a, b) // 2, 1

	// // example 2: when you set f to the reference of e, then update e, it will update f as well
	// e := 1
	// f := &e
	// e = 2
	// fmt.Print(e, *f) // 2, 2

	// // example 3: since slices are stored by reference which is the same as example 2,
	// //  when you update slice 1, slice 3 is updaetd as well
	// slice1 := []int{1, 2}
	// slice2 := []int{3, 4}
	// slice3 := slice1
	// copy(slice1, slice2)
	// fmt.Println(slice1, slice2, slice3) // [3, 4] [3, 4] [3, 4]

	// variatic("hello", "1", "2", "c")
	// return []int{}

	// func variatic(input1 string, input2 ...string) {
	// 	fmt.Print(input1, input2) // prints hello [1, 2, 3]
	// }

	// go fmt.Print("a") // asynchronous print a
	// fmt.Print("b")  // synchrounously print b
	// time.Sleep(time.Second) // pause synchronous execution for 1 second to wait for asynchronous goroutines
	//prints ba

	// create a channel called "messages" that moves a string
	// messages := make(chan string, 1) // 1 = maximum number of inputs into the channel
	// // call a goroutine that sends a "ping" string to the channel
	// go func() {
	// 	messages <- "ping"
	// }()
	// // receive the output of the channel in a variable called "msg"
	// go func() {
	// 	output := <-messages
	// 	fmt.Println(output) // "ping"
	// }()

	// output := <-messages
	// fmt.Println(output)

	// done := make(chan bool, 1)
	// // call the function above as a goroutine
	// go worker(done)
	// // receive the value from the channel
	// // without this, the sender will not send its message
	// <-done

	// 'i' is an interface that holds a string. The interface could potentially hold any type
	// var i interface{} = "hello"
	// value, ok := i.(string) // This is a valid type assertion.
	// if ok {
	// 	output := fmt.Sprintf("output: %s", value)
	// 	fmt.Print(output)
	// }

	// // example 1: non-slices are stored by value so when you update the value of a, b is not changed
	// a := 1
	// b := a
	// a = 2
	// fmt.Print(a, b) // 2, 1

	// // example 2: when you set f to the reference of e, then update e, it will update f as well
	// e := 1
	// f := &e
	// e = 2
	// fmt.Print(e, *f) // 2, 2

	// // example 3: since slices are stored by reference which is the same as example 2,
	// //  when you update slice 1, slice 3 is updaetd as well
	// slice1 := []int{1, 2}
	// slice2 := []int{3, 4}
	// slice3 := slice1
	// copy(slice1, slice2)
	// fmt.Println(slice1, slice2, slice3) // [3, 4] [3, 4] [3, 4]

	// variatic("hello", "1", "2", "c")

	runeMap := make(map[rune]int)

	// runeMap2 := map[rune]int{}

	for _, letter := range "hello" {
		runeMap[letter] = 1
	}
	return []int{}
}

// func worker(done chan bool) {
// 	fmt.Print("working...")
// 	time.Sleep(time.Second)
// 	fmt.Println("done")
// 	done <- true
// }

// func printA() {
// 	fmt.Print("a")
// }

// func printB() {
// 	fmt.Print("b")
// }
// func printC() {
// 	fmt.Print("c")
// }
// func printD() {
// 	fmt.Print("d")
// }

// func f(from string) {
// 	for i := 0; i < 3; i++ {
// 		fmt.Println(from, ":", i)
// 	}
// }
// func going(msg string) {
// 	fmt.Println(msg)
// }

func TopKFrequentTester() bool {
	fmt.Print(TopKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)) //true

	return true
}
