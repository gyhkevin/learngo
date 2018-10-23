package main

import (
	"io/ioutil"
	"fmt"
)

func grade(score int) string  {
	g := ""
	switch {
	case score<0 || score > 100:
		panic(fmt.Sprintf("unsupported value %d", score))
	case score<60:
		g = "E"
	case score<70:
		g = "D"
	case score<80:
		g = "C"
	case score<90:
		g = "B"
	case score<=100:
		g = "A"
	}
	return g
}

func main() {
	const filename = "abc.txt"
	//pwd,_ := os.Getwd()
	/*contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n",contents)
	}*/
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n",contents)
	}

	fmt.Println(
		grade(80),
		grade(99),
		grade(20),
		grade(67),
		grade(77),
		grade(100),
	)
}
