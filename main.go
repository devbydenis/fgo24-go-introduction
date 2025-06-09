package main

import (
	"fmt"
	"playground/circle"
	// "fgo24-go-introduction/circle/main.go"
)


func main() {
	// fmt.Println(circle.CircumferenceCircle(7))
	
	
	// scores1 := []int{10, 20, 30, 40, 50}
	// for i := 0; i < 5; i++ {
	// 	scores1 = append(scores1, 70 + i)
	// }
	// fmt.Println(scores1)
	// scores2 := append(scores1, 60)

	// fmt.Println(scores2)
	
	// skors := []int{50, 75, 66, 20, 32, 90}
	// var newScores [7] int
	// for x := range skors {
	// 	if (skors[x] == 20) {
	// 		newScores[x] = skors[x]
	// 	}
	// }
	// fmt.Println(newScores)
	// fgo24-go-append-slice
	// 
	// 
	
	fmt.Println(circle.AreaCircle(7))
	fmt.Println(circle.CircumferenceCircle(7))
}

