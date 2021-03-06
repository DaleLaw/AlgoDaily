package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// depth first search
// Time		O(n)
// Space	O(height of the tree) the call stack
// beats 14.29%
// 10sep2018
func closestValue(root *TreeNode, target float64) int {
	closest := root.Val
	var dfs func(cur *TreeNode)
	dfs = func(cur *TreeNode) {
		diff := math.Abs(target - float64(closest))
		cur_diff := math.Abs(target - float64(cur.Val))
		if cur_diff < diff {
			closest = cur.Val
		}
		if cur.Left != nil {
			dfs(cur.Left)
		}
		if cur.Right != nil {
			dfs(cur.Right)
		}
	}

	dfs(root)

	return closest
}

// recursive depth first search with return value
func closestValue1(root *TreeNode, target float64) int {
	cur_diff := math.Abs(target - float64(root.Val))
	left := math.MaxInt64
	right := math.MaxInt64
	if root.Left != nil {
		left = closestValue(root.Left, target)
	}
	left_diff := math.Abs(target - float64(left))
	if root.Right != nil {
		right = closestValue(root.Right, target)
	}
	right_diff := math.Abs(target - float64(right))
	if cur_diff < left_diff && cur_diff < right_diff {
		return root.Val
	} else if left_diff < cur_diff && left_diff < right_diff {
		return left
	} else if right_diff < cur_diff && right_diff < left_diff {
		return right
	}
	return root.Val
}

// iterative dfs
func closestValue2(root *TreeNode, target float64) int {

	result := math.MaxInt64

	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		var pop *TreeNode
		pop, stack = stack[len(stack)-1], stack[:len(stack)-1]

		if math.Abs(float64(pop.Val)-target) < math.Abs(float64(result)-target) {
			result = pop.Val
		}

		if pop.Left != nil {
			stack = append(stack, pop.Left)
		}
		if pop.Right != nil {
			stack = append(stack, pop.Right)
		}
	}

	return result
}

// BST search
// Time		O(logn)
// Space	O(height of the tree) the call stack
// beats 100%
// 20jan2019
func closestValue3(root *TreeNode, target float64) int {
	closest := root.Val
	var dfs func(cur *TreeNode)
	dfs = func(cur *TreeNode) {
		diff := math.Abs(target - float64(closest))
		cur_diff := math.Abs(target - float64(cur.Val))
		if cur_diff < diff {
			closest = cur.Val
		}
		if target < float64(cur.Val) && cur.Left != nil {
			dfs(cur.Left)
		} else if target > float64(cur.Val) && cur.Right != nil {
			dfs(cur.Right)
		}
	}
	dfs(root)
	return closest
}

func main() {
}
