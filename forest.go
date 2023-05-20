package pforest

type Token rune

type Node struct {
	head     Token
	count    int
	children map[Token]*Node
}

type Forest struct {
	parent   Node
	maxDepth int
}

func New(maxDepth int) Forest {
	return Forest{
		parent:   Node{children: make(map[Token]*Node)},
		maxDepth: maxDepth,
	}
}

func (f Forest) Learn(text string) {
	for k := 0; k < len(text); k++ {
		span := text[k:min(k+f.maxDepth, len(text))]
		f.add(span)
	}
}

func (f Forest) add(span string) {
	node := &f.parent
	for _, c := range span {
		token := Token(c)
		child, ok := node.children[token]
		if !ok {
			child = &Node{
				head:     token,
				count:    0,
				children: make(map[Token]*Node),
			}
			node.children[token] = child
		}
		child.count++
		node = child
	}
}
