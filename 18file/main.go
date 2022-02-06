package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome ot files in golang")
	content := "This needs to go in a file - LearnCodeOnline.in\n"

	file, err := os.Create("./mylcogofile.txt")

	checkNilError(err)

	lenght, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("Lenght is: ", lenght)
	defer file.Close()
	readFile("./mylcogofile.txt")

}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)

	fmt.Println("Text data inside the file \n", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
