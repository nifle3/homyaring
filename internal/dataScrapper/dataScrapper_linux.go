//go:build linux
// +build linux

package datascrapper

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

var (
	coreProcStatRe = regexp.MustCompile(`^cpu\s+(\d\s)+`)
	digitRegex     = regexp.MustCompile(`\d+`)
)

func GetCPU() (CPUInfo, error) {
	return CPUInfo{}, nil
}

func GetMemory() {
	_, err := os.Stat("/")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func GetRAM() (RAMInfo, error) {
	file, err := os.OpenFile("/proc/meminfo", os.O_RDONLY, 0666)
	if err != nil {
		return RAMInfo{}, err
	}
	defer file.Close()

	result := RAMInfo{}
	scanner := bufio.NewScanner(file)

	result.Total, err = readRAMString(scanner)
	if err != nil {
		return result, err
	}

	result.Free, err = readRAMString(scanner)
	resutl.

	return result, err
}

func readRAMString(scanner *bufio.Scanner) (int, error) {
	scanner.Scan()
	bytes := scanner.Bytes()
	ram := digitRegex.Find(bytes)
	strRam := string(ram)
	if ram == nil {
		return 0, fmt.Errorf("cannot find digit in %s", strRam)
	}

	result, err := strconv.Atoi(strRam)
	return result, err
}

func GetUptime() (UptimeInfo, error) {
	fileData, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return UptimeInfo{}, err
	}

	var allTime float64
	var idleTime float64
	if _, err := fmt.Sscanf(string(fileData), "%f %f", &allTime, &idleTime); err != nil {
		return UptimeInfo{}, err
	}

	allTimeHours := parseHours(allTime)
	allIdleHours := parseHours(idleTime)

	result := UptimeInfo{
		IdleTime: allIdleHours,
		AllTime:  allTimeHours,
	}
	result.CalculateTotalTime()

	return result, nil
}

func parseHours(uptime float64) float64 {
	hours := uptime / 3600.0
	hoursTwoSymbol := math.Round(hours*100) / 100
	return hoursTwoSymbol
}
