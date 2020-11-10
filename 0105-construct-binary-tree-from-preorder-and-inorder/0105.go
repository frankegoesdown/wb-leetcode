package main

import "fmt"

func main() {
	fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	fmt.Printf("preorder: %v \n", preorder)
	fmt.Printf("inorder: %v \n", inorder)
	if len(preorder) == 0 {
		return nil
	}

	index := search(preorder[0], inorder)
	node := &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:1+index], inorder[:index]),
		Right: buildTree(preorder[1+index:], inorder[index+1:]),
	}

	return node
}

func search(val int, array []int) int {
	fmt.Println("val:", val)
	fmt.Println("array:", array)
	for i := 0; i < len(array); i++ {
		if val == array[i] {
			return i
		}
	}

	return 0
}
