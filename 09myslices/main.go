package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Println("FruitList is:", fruitList)
	fmt.Printf("type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println("FruitList updated is:", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 445
	highScores[3] = 867
	fmt.Println(highScores)

	highScores = append(highScores, 555, 666, 777)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove a value form slices bases on index

	var courses = []string{"reactjs", "Javascript", "swift", "Python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
