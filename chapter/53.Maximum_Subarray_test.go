package chapter

import (
	"fmt"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	// nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums := []int{-2, 1}
	res := maxSubArray(nums)
	fmt.Println(res)
}
