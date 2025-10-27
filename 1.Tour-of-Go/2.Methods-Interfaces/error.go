package main

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrEmpty = errors.New("empty input")

// 2) Error tipado (con datos)
type RangeError struct{ Min, Max int }
func (e RangeError) Error() string { return fmt.Sprintf("out of range [%d,%d]", e.Min, e.Max) }

func parsePositive(s string) (int, error) {
	if s == "" {
		return 0, ErrEmpty
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("parse %q: %w", s, err)
	}
	if n < 0 {
		return 0, RangeError{Min: 0, Max: 1<<31 - 1}
	}
	return n, nil
}

func main() {
	tests := []string{"", "abc", "-5", "7"}
	for _, t := range tests {
		n, err := parsePositive(t)
		if err != nil {
			switch {
			case errors.Is(err, ErrEmpty):
				fmt.Println("empty:", err)
			case errors.As(err, new(*RangeError)):
				fmt.Println("range:", err)
			default:
				fmt.Println("other:", err)
			}
			continue
		}
		fmt.Println("ok:", n)
	}
}
