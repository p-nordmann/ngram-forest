package pforest

import "math/rand"

func (f Forest) Generate(n int) string {
	var tokens []Token
	var heads []*Node = []*Node{&f.parent}
	for k := 0; k < n; k++ {
		tokens, heads = updateTokenAndHeads(f, tokens, heads)
	}
	return string(tokens)
}

func updateTokenAndHeads(f Forest, tokens []Token, heads []*Node) ([]Token, []*Node) {
	// Select first head with children.
	var head *Node
	for _, node := range heads {
		head = node
		if len(head.children) > 0 {
			break
		}
	}

	// Prepare candidates and density for sampling next token.
	var density []int
	var candidates []Token
	for candidate, node := range head.children {
		candidates = append(candidates, candidate)
		density = append(density, node.count)
	}

	// Sample next token and update heads.
	token := candidates[sample(density)]
	tokens = append(tokens, token)
	var nextHeads []*Node
	for _, head := range heads {
		node, ok := head.children[token]
		if ok {
			nextHeads = append(nextHeads, node)
		}
	}
	nextHeads = append(nextHeads, &f.parent)

	return tokens, nextHeads
}

// Samples an index according to the density.
//
// TODO implement less naive approach
func sample(density []int) int {
	var total int
	var stairs []int
	for _, d := range density {
		stairs = append(stairs, total)
		total += d
	}
	dice := rand.Intn(total)
	for k, stair := range stairs {
		if dice < stair {
			return k - 1
		}
	}
	return len(stairs) - 1
}

// Predicts next token probability, up to a constant multiplicative factor.
func (f Forest) Predict(context string) map[Token]int {
	for length := f.maxDepth; length >= 0; length-- {
		if len(context) < length {
			continue
		}
		suffix := context[len(context)-length:]
		children, ok := f.findChildren(suffix)
		if ok && len(children) > 0 {
			density := make(map[Token]int)
			for token, child := range children {
				density[token] = child.count
			}
			return density
		}
	}
	return nil // Cannot actually happen.
}

func (f Forest) findChildren(suffix string) (map[Token]*Node, bool) {
	node := &f.parent
	for _, token := range suffix {
		child, ok := node.children[Token(token)]
		if !ok {
			return nil, false
		}
		node = child
	}
	return node.children, true
}
