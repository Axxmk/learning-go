package input

import "fmt"

func GetInput[T any](args ...*T) error {
	for _, a := range args {
		if _, err := fmt.Scan(a); err != nil {
			return err
		}
	}
	return nil
}
