package main

import "fmt"

func main() {
	i := 10
	// j := 4
	if i >= 5 {
		fmt.Println(" 5 이상. ")
	}
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

// 세미콜론은 컴파일러에서 붙혀줌. 한줄주석
/*
 여러줄 주석

 gofmt filename
 gofmt -w filename
*/
