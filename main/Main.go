package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	i1 := myInt(1)
	i2 := i1.add(2)

	fmt.Println(i1, i2)

	testStruct()
	testDefer()

	printDeferNumber()
}

func testReader() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name:")

	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("Found an error : %s\n", err)
	} else {
		input = input[:len(input)-1]
		fmt.Printf("Hello,%s!\n", input)
	}

}

type myInt int

func (i *myInt) add(another int) myInt {
	*i = *i + myInt(another)
	return *i
}

type student struct {
	name string
	age  int
}

func testStruct() {
	xiaoming := student{name: "Xiaoming", age: 2}
	xiaoming2 := student{"Xiaoming", 2}
	fmt.Print(xiaoming == xiaoming2)

}

func testDefer() {
	defer fmt.Println("defer something")
	fmt.Println("print something")
}

func printDeferNumber() {
	for i := 0; i < 5; i++ {

		defer func(n int) {
			fmt.Printf("%d", n)
		}(i)
	}
}
