package main

import (
	"fmt"
)

func main() {
	s := "fsdf"
	for i, v := range s {
		fmt.Printf("%T %T", s[i], v) // uint8(byte) int32(rune)
	}

}
