//go:build windows
// +build windows

package datascrapper

import "fmt"

func GetUptime() (UptimeInfo, error) {
	fmt.Println("Hello world")
	return UptimeInfo{}, nil
}

func GetRAM() (RAMInfo, error) {
	return RAMInfo{}, nil
}
