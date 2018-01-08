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
