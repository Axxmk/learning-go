package array

import "fmt"

func FormattedPrint[T any](list []T, msg string) {
	fmt.Print(msg)

	last := len(list) - 1
	for i, el := range list {
		if i == last {
			fmt.Println(el)
		} else {
			fmt.Print(el, ", ")
		}
	}
}
