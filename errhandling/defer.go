package main

import (
	"fmt"
	"os"
	"bufio"
	"imooc.com/kevin/learngo/functional/fib"
)

func tryDefer()  {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("Printed to many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	defer write.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(write,f())
	}
}

func main()  {
	tryDefer()
	writeFile("fib.txt")
}