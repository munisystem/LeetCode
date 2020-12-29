package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 9,
	}
	fmt.Println(pseudoPalindromicPaths(root))
}

func pseudoPalindromicPaths(root *TreeNode) int {
	return dfs(root, []int{})
}

func dfs(node *TreeNode, nums []int) int {
	nums = append(nums, node.Val)
	if node.Left == nil && node.Right == nil {
		m := make(map[int]int, 10)
		for i := 0; i < len(m); i++ {
			m[i] = 0
		}
		for i := 0; i < len(nums); i++ {
			m[nums[i]]++
		}
		odds := 0
		for _, v := range m {
			if v%2 != 0 {
				odds++
			}
		}
		if (len(nums)%2 == 0 && odds == 0) || (len(nums)%2 == 1 && odds == 1) {
			return 1
		}
		return 0
	}
	n := 0
	if node.Left != nil {
		n = n + dfs(node.Left, nums)
	}
	if node.Right != nil {
		n = n + dfs(node.Right, nums)
	}
	return n
}
