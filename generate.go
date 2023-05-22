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
