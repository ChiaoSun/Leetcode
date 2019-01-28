package main

import "fmt"
import "strings"

func main() {
	fmt.Println(detectCapitalUse("USA"))
}

func detectCapitalUse(word string) bool {
	arr := []byte(word)
	count := 0
	for key := range arr {
		temp := string(arr[key])
		i := strings.Compare(temp, strings.ToLower(temp))
		if i != 0 {
			count++
		}
	}
	if count > 0 {
		return false
	}
	return true
}
