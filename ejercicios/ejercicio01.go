package ejercicios

import "strconv"

func ConvertToNumber(val string) (int, string) {
	number, err := strconv.Atoi(val)

	if err != nil {
		return 0, "Error dureng conversion" + err.Error()
	} else if number > 100 {
		return number, "Es mayor a 100"
	} else {
		return number, "Es menor a 100"
	}
}
