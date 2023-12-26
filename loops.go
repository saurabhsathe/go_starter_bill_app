package main

import "fmt"

func main() {
	var names []string = []string{}
	names = append(names, "saurabh")
	names = append(names, "pratik")
	names = append(names, "harsh")

	for idx, val := range names {
		fmt.Println(idx, ",", val)

	}
	x := 0
	for x < len(names) {
		fmt.Println("loop2", names[x])
		x++
	}

	for x = 0; x < len(names); x++ {
		fmt.Println("loop3", names[x])

	}
}
