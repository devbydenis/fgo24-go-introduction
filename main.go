package main

import (
	"fmt"
)

func areaCircle(r float32) float32 {
	return r * r * 3.14159265358979323846
}
func circumferenceCircle(r float32) float32 {
	return 2 * r * 3.14159265358979323846
}

func main() {
	fmt.Println(areaCircle(5))
	fmt.Println(circumferenceCircle(5))
}