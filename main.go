package main

import "fmt"

func Space() {
	fmt.Println()
}

func Spaces() {
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func TernaryIF(condition bool, trueVal string, falseVal string) string {
	if !condition {
		return falseVal
	}
	return trueVal
}
