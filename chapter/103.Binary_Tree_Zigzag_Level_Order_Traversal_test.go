package chapter

import (
	"fmt"
	"testing"
)

func Test_zigzagLevelOrder(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	res := zigzagLevelOrder(root)
	fmt.Println(res)
}
