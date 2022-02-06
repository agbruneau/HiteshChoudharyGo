// test git
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//const url = "https://google.com"
const url = "https://lco.dev"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)

	checkNilError(err)

	fmt.Printf("Response is of tye : %T", response)

	defer response.Body.Close() // caller's responsability to close response

	databytes, err := ioutil.ReadAll(response.Body)

	checkNilError(err)

	content := string(databytes)
	fmt.Printf("Content is:\n%s\n", content)

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
