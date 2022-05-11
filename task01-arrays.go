package homework

func average(input [5]float32) (result float32) {
	var sum float32
	length := len(input)
	for i := 0; i < length; i++ {
		sum += input[i]
	}
	return sum / float32(length)
}
