package internal

func Exercise_2(init, end, incr int) int {
	count := 0
	//i>0 vale 5puntos
	for i := init; i <= end; i += incr {
		if i%2 == 0 {
			count++
		}
	}
	return count
}
