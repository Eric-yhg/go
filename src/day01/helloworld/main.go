package main

import "fmt"

func main() {
	for i := 1; i < 13; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
