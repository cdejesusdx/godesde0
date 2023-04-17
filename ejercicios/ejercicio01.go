package ejercicios

import "strconv"

func ConvertToNumber(val string) (int, string) {
	number, error := strconv.Atoi(val)

	if error != nil {
		return 0, "Error dureng conversion"
	} else if number > 100 {
		return number, "Es mayor a 100"
	} else {
		return number, "Es menor a 100"
	}
}
