package helpers_test

import (
	"testing"

	"github.com/AJBC55/helpers"
)

func addOne(i int) int {
	return i + 1
}

func TestMap(t *testing.T) {
	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = i
	}
	addedNums := helpers.Map(nums, addOne)
	for i := 1; i < 11; i++ {
		if addedNums[i-1] != i {
			t.FailNow()
		}
	}

}
