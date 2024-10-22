package main

import (
	"fmt"

	datascrapper "github.com/nifle3/homyaring/internal/dataScrapper"
)

func main() {
	info, _ := datascrapper.GetUptime()
	fmt.Printf("%#v\n", info)

	ramInfo, err := datascrapper.GetRAM()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%#v", ramInfo)
	fmt.Println()
}
