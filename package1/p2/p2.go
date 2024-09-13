package p2

func HalfDaSlice(input []int) []int {

	var halfnum int = len(input) / 2

	input = input[:halfnum-1]

	return input
}
