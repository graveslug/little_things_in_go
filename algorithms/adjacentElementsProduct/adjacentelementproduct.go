package adjacentelementproduct

func adjacentElementsProduct(inputArray []int) int {
	holder := inputArray[0] * inputArray[1]
	highNum := inputArray[0] * inputArray[1]

	for i := 0; i < len(inputArray)-1; i++ {
		holder = inputArray[i] * inputArray[i+1]
		if holder > highNum {
			highNum = holder
		}
	}
	return highNum
}
