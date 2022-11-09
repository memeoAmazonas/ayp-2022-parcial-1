package internal

func Exercise_3(data []string) string {

	response := ""
	for _, item := range data {
		if !isNumber(item) {
			response += item
		}
	}
	return response
}

func isNumber(num string) bool {
	numList := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	for _, x := range numList {
		if x == num {
			return true
		}
	}
	return false
}
