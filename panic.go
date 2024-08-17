package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("upps error")
	}
	fmt.Println("Run App")
}

func main() {
	runApp(true)
	fmt.Println("End Main")
}
