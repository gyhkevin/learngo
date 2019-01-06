package main

import (
	"time"
	"fmt"
)

func main()  {
	//var a [10]int
	for i := 0; i < 1000; i++ {
		go func(j int) { // race condition
			for  {
				fmt.Printf("Hello from" + "gorountine %d\n", i)
				//a[j]++
				//runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Minute)
	//fmt.Println(a)
}
