package main

import (
	"fmt"
	"strings"
)

func main() {
	//stringOps()
	//mapOps()
}

func stringOps() {
	var str1 = "Hello"
	fmt.Println(len(str1))
	for i, elem := range str1 {
		fmt.Println(i, str1[i], elem)
	}
	for elem := range str1 {
		fmt.Println(elem)
	}
	for _, r := range str1 {
		c := string(r)
		fmt.Println(c)
	}
	str2 := "Hi! I am bored"
	strArr := strings.Split(str2, " ")
	for _, r := range strArr {
		c := string(r)
		fmt.Println(c)
	}

}

func mapOps() {

	var map1 = make(map[string]int64)
	fmt.Println(map1["a"])
	var ctr int64 = 34
	ctr++
	map1["a"] = map1["a"] + 1

	fmt.Println(map1["a"])
	map1["a"] = ctr
	fmt.Println(map1["a"])
}
