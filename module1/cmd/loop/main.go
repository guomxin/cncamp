package main

import "fmt"

func main() {
	arrs := [5]string{"I", "am", "stupid", "and", "weak"}
	fmt.Printf("Before change arrs=%v\n", arrs)
	for i := 1; i < len(arrs); i++ {
		if i == 2 {
			arrs[i] = "smart"
		} else if i == 4 {
			arrs[i] = "strong"
		}
	}
	fmt.Printf("After change arrs=%v\n", arrs)
}
