package helpers

// helper Map function to map an array of one type to another type
func Map[I, O any](inputArr []I, transformer func(I) O) []O {
	outputArr := make([]O, len(inputArr))
	for i := range inputArr {
		outputArr[i] = transformer(inputArr[i])
	}
	return outputArr
}
