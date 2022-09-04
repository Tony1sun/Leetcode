package main

import (
	"fmt"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
}
