package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ArgumentsOnStack(lines []string) int {

	regexpRbpLoadHigher := regexp.MustCompile(`\[rbp \+ ([0-9]+)\]$`)

	offsets := make(map[int]bool)

	for _, l := range lines {

		if match := regexpRbpLoadHigher.FindStringSubmatch(l); len(match) > 1 {
			offset, _ := strconv.Atoi(match[1])
			if _, found := offsets[offset]; !found {
				offsets[offset] = true
			}
		}
	}

	fmt.Println(offsets)

	return len(offsets)
}

func GetGolangArgs(proto string) int {

	regexpFunc := regexp.MustCompile(`func .*\((.*)\)`)

	if match := regexpFunc.FindStringSubmatch(proto); len(match) > 1 {
		return len(strings.Split(match[1], ","))
	} else {
		panic(fmt.Sprintf("Bad function prototype: %s", proto))
	}
}
