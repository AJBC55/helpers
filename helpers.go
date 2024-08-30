package helpers

// helper Map function to map an array and performs a transformation on that array
// either changing it's type or somthing else
func Map[I any, O any](inputArr []I, transformer func(I) O) []O {
	outputArr := make([]O, len(inputArr))
	for i := range inputArr {
		outputArr[i] = transformer(inputArr[i])
	}
	return outputArr
}
