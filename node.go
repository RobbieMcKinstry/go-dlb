package dlb

// A Node is a member of the trie,
// containing a character.
// The bitmap is used to store whether it has a
// child with the given pattern for O(1) checking.
// The NodeStore contains its references to the next
// nodes in the trie.
type Node struct {
	b byte
	NodeStore
}

// NewNode returns a new node with this byte
// pattern and no children
func NewNode(b byte) *Node {
	return &Node{
		b:         b,
		NodeStore: NewUnorderedStore(),
	}
}

// Add will recursively add this bytestring to the trie.
func (n *Node) Add(b []byte) {
	var (
		first     = b[0]
		rest      = b[1:]
		childNode = n.Fetch(first)
	)

	if childNode == nil {
		// Add a new leaf to this level.
		next := NewNode(first)
		next.Add(rest)
		n.Store(next)
		return
	}

	// Otherwise, recurse with the rest of the string.
	childNode.Add(rest)
}
