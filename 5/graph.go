package main

type Node[T comparable] struct {
	prev  []*Node[T]
	next  []*Node[T]
	value T
}

func (n *Node[T]) Successors() []T {
	result := []T{}
	for _, x := range n.next {
		result = append(result, x.value)
	}
	return result
}

func (n *Node[T]) Predecessors() []T {
	result := []T{}
	for _, x := range n.prev {
		result = append(result, x.value)
	}
	return result
}

type Graph[T comparable] struct {
	nodeIndex map[T]*Node[T]
	initNodes []*Node[T]
}

func (g *Graph[T]) Insert(prevValue T, nextValue T) {
	if g.nodeIndex == nil {
		g.nodeIndex = make(map[T]*Node[T])
	}
	if _, ok := g.nodeIndex[prevValue]; !ok {
		//fmt.Println("Creating node for", prevValue)
		g.nodeIndex[prevValue] = &Node[T]{value: prevValue, prev: []*Node[T]{}, next: []*Node[T]{}}
	}
	if _, ok := g.nodeIndex[nextValue]; !ok {
		//fmt.Println("Creating node for", nextValue)
		g.nodeIndex[nextValue] = &Node[T]{value: nextValue, prev: []*Node[T]{}, next: []*Node[T]{}}
	}

	n1 := g.nodeIndex[prevValue]
	n2 := g.nodeIndex[nextValue]
	n1.next = append(n1.next, n2)
	n2.prev = append(n2.prev, n1)
}

func (g *Graph[T]) Consolidate() {
	if g.initNodes == nil {
		g.initNodes = []*Node[T]{}
	}
	for _, n := range g.nodeIndex {
		if len(n.prev) == 0 {
			g.initNodes = append(g.initNodes, n)
		}
	}
}

func (g Graph[T]) Successors(page T) []T {
	return g.nodeIndex[page].Successors()
}

func (g Graph[T]) Predecessors(page T) []T {
	return g.nodeIndex[page].Predecessors()
}

func (g *Graph[T]) Sort(sequence []T, less func(*Graph[T], T, T) bool) []T {
	if len(sequence) < 2 {
		return sequence
	}

	pivot := sequence[0]
	left, right := []T{}, []T{}

	for _, x := range sequence[1:] {
		if less(g, x, pivot) {
			left = append(left, x)
		} else {
			right = append(right, x)
		}
	}

	return append(append(g.Sort(left, less), pivot), g.Sort(right, less)...)
}

// Custom 'less than' function for use with 'Graph.Sort()'
func isPrevious(g *Graph[int], i int, j int) bool {
	if contains(i, g.Predecessors(g.nodeIndex[j].value)) {
		return true
	} else {
		return false
	}
}

func contains(item int, sequence []int) bool {
	for _, x := range sequence {
		if item == x {
			return true
		}
	}
	return false
}
