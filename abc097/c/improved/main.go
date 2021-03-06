package main

import (
	"fmt"
	"math"
	"sort"
)

//コンテスト後に解説見たやつ

func main() {
	var s string
	var K int
	fmt.Scan(&s)
	fmt.Scan(&K)

	var substrs []string

	for i := 0; i < len(s); i++ {
		for j := int(math.Min(float64(i+5), float64(len(s)))); j > i; j-- {
			if !strSearch(substrs, s[i:j]) {
				substrs = append(substrs, s[i:j])
			}
		}
	}
	sort.Strings(substrs)

	fmt.Println(substrs[K-1])
}

func strSearch(a []string, b string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == b {
			return true
		}
	}
	return false
}
