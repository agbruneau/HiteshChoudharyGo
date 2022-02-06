package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var signals = []string{} // Stack the listed endpoint after calling them on the web

var wg sync.WaitGroup // usually this data are pointers
var mut sync.Mutex    // usually this data are pointers

func main() {
	fmt.Println("Go routine Introduction")

	go greeter("Hello")
	greeter("world")

	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://ibm.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatuscode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

func getStatuscode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPD in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s \n", res.StatusCode, endpoint)
	}
}
