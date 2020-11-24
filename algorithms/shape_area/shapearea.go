package shapearea

func shapeArea(n int) int {
	if n == 1 {
		return n
	}
	area := 1

	for i := 2; i <= n; i++ {
		temp := ((i - 1) * 4) + area
	}
}
