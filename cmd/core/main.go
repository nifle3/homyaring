package main

import (
	"fmt"
	"os"
)

func getTotalMemory() {
	result, err := os.ReadFile("/proc/stat")
	if err != nil {
		fmt.Printf("err: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println(string(result))
}

func main() {
	getTotalMemory()
}
