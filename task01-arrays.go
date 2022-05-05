package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var sum float32 = 0
	for _, value := range input {
		sum += value
	}
	result = sum / float32(len(input))
	return result
}
