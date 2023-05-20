package pforest

func increaseNode(n1, n2 *Node) {
	n1.count += n2.count
	for token, child2 := range n2.children {
		child1, ok := n1.children[token]
		if !ok {
			child1 = &Node{
				head:     token,
				count:    0,
				children: make(map[Token]*Node),
			}
			n1.children[token] = child1
		}
		increaseNode(child1, child2)
	}
}

func Sum(f1, f2 Forest) Forest {
	f := New(max(f1.maxDepth, f2.maxDepth))
	increaseNode(&f.parent, &f1.parent)
	increaseNode(&f.parent, &f2.parent)
	return f
}

func joinNodes(n0, n1 *Node) {
	for token, child1 := range n1.children {
		child0, ok := n0.children[token]
		if !ok {
			child0 = &Node{
				head:     token,
				count:    0,
				children: make(map[Token]*Node),
			}
			n0.children[token] = child0
		}
		child0.count = max(child0.count, child1.count)
		joinNodes(child0, child1)
	}
}

func Union(f1, f2 Forest) Forest {
	f := New(max(f1.maxDepth, f2.maxDepth))
	joinNodes(&f.parent, &f1.parent)
	joinNodes(&f.parent, &f2.parent)
	return f
}

func intersectNodes(n0, n1, n2 *Node) {
	for token, child1 := range n1.children {
		child2, ok := n2.children[token]
		if !ok {
			continue
		}
		n0.children[token] = &Node{
			head:     token,
			count:    min(child1.count, child2.count),
			children: make(map[Token]*Node),
		}
		intersectNodes(n0.children[token], child1, child2)
	}
}

func Intersection(f1, f2 Forest) Forest {
	f := New(min(f1.maxDepth, f2.maxDepth))
	intersectNodes(&f.parent, &f1.parent, &f2.parent)
	return f
}
