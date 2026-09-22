package main

import "fmt"

func main() {
	grades := []int{5, 4, 3, 5, 4, 2, 5, 3, 4, 5}

	fives := 0
	lowGrades := 0

	/*for _, grade := range grades {
	  switch {
	  case grade == 5:
	      fives++
	  case grade <= 3:
	      lowGrades++
	  } */

	for _, grade := range grades {
		if grade == 5 {
			fives++
		} else if grade <= 3 {
			lowGrades++
		}
	}

	fmt.Printf("Пятерок: %d\n", fives)
	fmt.Printf("Троек и ниже: %d\n", lowGrades)
}
