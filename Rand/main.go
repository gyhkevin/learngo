package main

import (
	"imooc.com/kevin/learngo/Rand/SampleRand"
	"fmt"
)

func main() {
	r1 := SampleRand.RandStringRunes(4)
	fmt.Println(r1)

	r2 := SampleRand.RandStringBytes(4)
	fmt.Println(r2)

	r3 := SampleRand.RandStringBytesRmndr(4)
	fmt.Println(r3)

	r4 := SampleRand.RandStringBytesMask(4)
	fmt.Println(r4)

	r5 := SampleRand.RandStringBytesMaskImpr(4)
	fmt.Println(r5)
}
