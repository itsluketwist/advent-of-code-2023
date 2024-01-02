package utils

import "fmt"

func PrintList[T any](list []T) {
	for _, item := range list {
		fmt.Println(item)
	}
}
