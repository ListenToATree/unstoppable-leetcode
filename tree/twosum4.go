package tree

import (
	"unstoppable-leetcode/shared"
)

type BSTNode struct {
	Val   int
	Left  *BSTNode
	Right *BSTNode
}

// Time: O(n) Space: O(n)
// using set
func findTarget(root *BSTNode, k int) bool {
	set := shared.NewSet[int]()
	return find(root, k, set)
}

func find(root *BSTNode, k int, set *shared.Set[int]) bool {
	if root == nil {
		return false
	}
	if set.Has(k - root.Val) {
		return true
	}
	set.Add(root.Val)
	return find(root.Left, k, set) || find(root.Right, k, set)
}
