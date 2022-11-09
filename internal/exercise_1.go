package internal

import "fmt"

//vale 3 puntos
/*func Exercise(values []string) {
	for i := -5; j > len(values); {

		if i%2 == 0 {
			i += 2
		} else {
			j++
		}
		fmt.Println("el valor es ", values[i])
	}
}*/
func Response(values []string) {
	for i := -5; i < len(values); {

		if i%2 == 0 {
			i += 2
		} else {
			i++
		}
		//validar que sea un numero dentro del rango de indices, con un if
		if i >= 0 && i < len(values) {
			fmt.Println("el valor es ", values[i])
		}
	}
}
