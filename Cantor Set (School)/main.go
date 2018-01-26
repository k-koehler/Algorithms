/*
Kevin Koehler
kak750
11163209
CMPT 360 -- A3
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
	"strings"
)

//processes the info to readable form for processCantorSet
func parseInput(strs []string) []string {
	for i := 0; i < len(strs); i++ {
		strs[i] = strings.Replace(strs[i], "0.", "", -1)
		strs[i] = strings.Replace(strs[i], "1.", "", -1)
	}
	return strs
}

//does powers of ints
func powInt(base int, exponent int) int {
	for i := 1; i < exponent; i++ {
		base = base * base
	}
	return base
}

//this func should be O(log_3(n)) where n is number of digits
func isCantor(s string, level int) string {
	var numerator, denominator int
	fmt.Sscanf(s, "%d", &numerator)
	denominator = powInt(10, len(s))
	//check if first digit is a 0 or 2 in ternary expansion
	if powInt(3, level)*numerator < denominator || powInt(3, level)*numerator > powInt(2, level)*denominator {
		if level == len(s) {
			return "MEMBER"
		} else {
			return isCantor(s, level+1)
		}
	} else {
		return "NON-MEMBER"
	}
}

//iterates through the usable list
func processCantorSet(strs []string) []string {
	retList := []string{}
	for i := 0; i < len(strs); i++ {
		if strs[i] == "0" { //trivial case 0.0 or 1.0
			retList = append(retList, "MEMBER")
		} else {
			retList = append(retList, isCantor(strs[i], 1))
		}
	}
	return retList
}

func main() {

	//read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	strs := []string{}
	for scanner.Scan() {
		strs = append(strs, scanner.Text())
	}

	//process the info
	usableInfo := parseInput(strs)
	output := processCantorSet(usableInfo)

	//output everything
	for i := 0; i < len(output); i++ {
		fmt.Println(output[i])
	}
}
