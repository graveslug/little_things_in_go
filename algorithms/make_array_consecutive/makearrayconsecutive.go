package makearrayconsecutive

func makeArrayConsecutive(inputArray []int) int {
	max := inputArray[0]
	min := inputArray[0]

	for i := 0; i < len(inputArray); i++ {
		if max < inputArray[i] {
			max = inputArray[i]
		}
		if min > inputArray[i] {
			min = inputArray[i]
		}
	}
	return max - min + 1 - len(inputArray)
}

func makeArrayConsecutive2(input []int) int {
	max := input[0]
	min := input[0]

	for _, n = range input {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	return max - min + 1 - len(input)
}
