// binary_tree 二叉树
package main

import (
	"fmt"
)

type Node struct {
	data  string
	left  *Node
	right *Node
}

func main() {
	nodeG := Node{data: "g", left: nil, right: nil}
	nodeF := Node{data: "f", left: &nodeG, right: nil}
	nodeE := Node{data: "e", left: nil, right: nil}
	nodeD := Node{data: "d", left: &nodeE, right: nil}
	nodeC := Node{data: "c", left: nil, right: nil}
	nodeB := Node{data: "b", left: &nodeD, right: &nodeF}
	nodeA := Node{data: "a", left: &nodeB, right: &nodeC}

	// fmt.Println(breadthFirstSearch(nodeA))

	preOrderRecursive(nodeA)
}

func breadthFirstSearch(node Node) []string {
	var result []string
	var nodes []Node = []Node{node}

	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		result = append(result, node.data)
		if node.left != nil {
			nodes = append(nodes, *node.left)
		}
		if node.right != nil {
			nodes = append(nodes, *node.right)
		}
	}
	return result
}

func preOrderRecursive(node Node) {
	fmt.Println(node.data)
	if node.left != nil {

		preOrderRecursive(*node.left)
	}

	// 在这里输出就是中序
	if node.right != nil {

		preOrderRecursive(*node.right)
	}
	// 在这里输出是后序

}
