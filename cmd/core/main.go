package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

func getTotalCPU() {
	result, err := os.ReadFile("/proc/stat")
	if err != nil {
		fmt.Printf("err: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println(string(result))
}

func getTotalMemory() {
	_, err := os.Stat("/")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func getTotalUser() {

}

func getProcessCount() {

}

type UptimeInfo struct {
	idleTime time.Duration
	allTime  time.Duration
}

func showUptime() (UptimeInfo, error) {
	fileData, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return UptimeInfo{}, err
	}

	var allTime float64
	var idleTime float64
	if _, err := fmt.Sscanf(string(fileData), "%f %f", &allTime, &idleTime); err != nil {
		return UptimeInfo{}, err
	}

	allTimeDuration := ParseTimeDuration(allTime)
	fmt.Println(allTimeDuration)
	doo := ParseTimeDuration(idleTime)
	fmt.Println(doo)

	return UptimeInfo{}, nil
}

func ParseTimeDuration(uptime float64) float64 {
	hours := uptime / 180.0
	hoursTwoSymbol := math.Round(hours*100) / 100
	return hoursTwoSymbol
}

func main() {
	getTotalMemory()
	showUptime()
}
