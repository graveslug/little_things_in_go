package main

func increasingSequence(sequence []int) bool {
	c1 := 0
	c2 := 0
	for i := 0; i < len(sequence)-1; i++ {
		if sequence[i] >= sequence[i+1] {
			c1++
		}
	}
	for j := 0; j < len(sequence)-2; j++ {
		if sequence[j] >= sequence[j+2] {
			c2++
		}
	}
	return (c1 <= 1) && (c2 <= 1)
}

func main() {
	input := []int{1, 2, 5, 4}              //true
	input2 := []int{1, 1}                   //true
	input3 := []int{1, 2, 1, 2}             //false
	input4 := []int{3, 6, 5, 8, 10, 20, 15} //false
	increasingSequence(input)
	increasingSequence(input2)
	increasingSequence(input3)
	increasingSequence(input4)
}
