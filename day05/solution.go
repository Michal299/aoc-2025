package day05

import (
	"strconv"
	"strings"
)

type Node struct {
	start int
	end   int

	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func Part1(lines []string) int {

	tree := Tree{root: nil}

	productsToValidateId := 0

	for i, line := range lines {
		if line == "" {
			productsToValidateId = i
			break
		}

		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		tree.insertRange(start, end)
	}

	result := 0
	for _, productId := range lines[productsToValidateId:] {
		productIdVal, _ := strconv.Atoi(productId)
		if tree.exists(productIdVal) {
			result++
		}
	}
	return result
}

func Part2(lines []string) int {
	tree := Tree{root: nil}

	for _, line := range lines {
		if line == "" {
			break
		}

		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		tree.insertRange(start, end)
	}
	return tree.countAllAvailable()
}

func (t *Tree) insertRange(start, end int) {
	t.root = t.root.insertRange(start, end)
}

func (n *Node) insertRange(start, end int) *Node {
	if n == nil {
		return &Node{start: start, end: end}
	}

	// range already inside given node
	if n.start <= start && n.end >= end {
		return n
	}

	// if range is overlapping
	if (start <= n.start && end >= n.start) || (start <= n.end && end >= n.end) || (start <= n.start && end >= n.end) {
		newStart := min(n.start, start)
		newEnd := max(n.end, end)
		n.start = newStart
		n.end = newEnd

		// merge childrens if possible
		return n.mergeChildrens()
	}

	// range is completely on left side
	if end < n.start {
		n.left = n.left.insertRange(start, end)
	}

	// range is completely on right side
	if start > n.end {
		n.right = n.right.insertRange(start, end)
	}
	return n
}

func (n *Node) mergeChildrens() *Node {
	newNode := &Node{start: n.start, end: n.end}

	newNode.insertNode(n.left)
	newNode.insertNode(n.right)

	return newNode
}

func (n *Node) insertNode(node *Node) {
	if node == nil {
		return
	}
	n.insertRange(node.start, node.end)
	n.insertNode(node.left)
	n.insertNode(node.right)
}

func (t *Tree) exists(value int) bool {
	return t.root.exists(value)
}

func (n *Node) exists(value int) bool {
	if n == nil {
		return false
	}

	if value < n.start {
		return n.left.exists(value)
	} else if value > n.end {
		return n.right.exists(value)
	}
	return value >= n.start && value <= n.end
}

func (t *Tree) countAllAvailable() int {
	return t.root.countAllAvailable()
}

func (n *Node) countAllAvailable() int {
	if n == nil {
		return 0
	}

	return n.lenght() + n.left.countAllAvailable() + n.right.countAllAvailable()
}

func (n *Node) lenght() int {
	return n.end - n.start + 1
}
