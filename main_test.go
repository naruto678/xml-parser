package main

import (
	"fmt"
	"testing"
	"unicode"
)

type Celsius float32
type Fahrenheit float32

func toFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit((9*c)/5 + 32)
}

func getMonth(n int) (month string) {
	switch n {
	case 1, 2, 12:
		month = "Winter"
	case 3, 4, 5:
		month = "Spring"
	case 6, 7, 8:
		month = "Summer"
	case 9, 10, 11:
		month = "Autumn"
	}
	return month
}

func TestFirstChallenge(t *testing.T) {
	celsius := 100
	fahrenheit := toFahrenheit(Celsius(celsius))
	if fahrenheit != 212 {
		t.Fatal("Fahrenheit function failed ")
	}
}

func TestPointer(t *testing.T) {
	val := 5
	fmt.Printf("An integer: %d . it's location in memory is %p\n", val, &val)
}

func TestSwitch(t *testing.T) {
	month_num := 3
	month := getMonth(month_num)
	if month != "Spring" {
		t.Fatal("Switch case test failed")
	}

}

func computeFactorial(n int) int {
	if n <= 1 {
		return 1
	} else {
		return n * computeFactorial(n-1)
	}
}

func TestFactorial(t *testing.T) {
	input := 5
	output := computeFactorial(input)
	if output != 120 {
		t.Fatal("Factorial failed ")
	}
}

func TestEven(t *testing.T) {
	input_arr := []int{1, 2, 3, 4, 5}
	filtered_arr := filter(input_arr, func(num int) bool {
		return num%2 == 0
	})
	for _, val := range filtered_arr {
		if val%2 != 0 {
			t.Fatal("Filter failed")
		}
	}
	fmt.Println(filtered_arr)

}

func filter(input_arr []int, criteria func(int) bool) []int {
	filtered_arr := make([]int, 0)
	for _, val := range input_arr {
		if criteria(val) {
			filtered_arr = append(filtered_arr, val)
		}
	}
	return filtered_arr
}

func TestAdhoc(t *testing.T) {
	val := "arnab"
	fmt.Println("This is the val ", val[2])
	fmt.Println("This is byte", []byte("123"))
	fmt.Println("this i sth etest", unicode.IsLetter('<'))

}
