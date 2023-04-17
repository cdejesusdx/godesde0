package interaciones

import (
	"fmt"
)

func InterarTo10() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func InterarToReverse() {
	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}
}

func InterarToBreak() {
	for i := 10; i >= 1; i-- {
		if i == 6 {
			break
		}
		fmt.Println(i)
	}
}

func InterarToContinue() {
	for i := 10; i >= 1; i-- {
		if i == 6 {
			continue
		}
		fmt.Println(i)
	}
}
