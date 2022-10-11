package main

import "fmt"

func main() {
	var arr = [5]string{"I", "am", "stupid", "and", "weak"}

	for index, value := range arr {
		if value == "stupid" {
			arr[index] = "smart"
		}
		if value == "weak" {
			arr[index] = "strong"
		}
	}
	fmt.Println(arr)
}
