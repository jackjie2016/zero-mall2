package main

import "fmt"

func main() {

	s := make([]int, 1023)

	for i := 0; i < 65; i++ {
		s = append(s, 1)
		fmt.Printf("s:len:%d,cap：%d\n", len(s), cap(s))

	}

}
