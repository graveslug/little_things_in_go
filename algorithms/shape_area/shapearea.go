package shapearea

//Iterative solution
func shapeArea(n int) int {
	if n == 1 {
		return n
	}

	area := 1

	for i := 2; i <= n; i++ {
		temp := ((i - 1) * 4) + area
		area = temp
	}
	return
}

//Recursive solution
func shapeArea2(n int) int {
	if n == 1 {
		return n
	}

	return shapeArea(n-1) + (4 * (n - 1))
}
