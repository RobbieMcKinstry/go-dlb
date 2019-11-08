package dlb

const (
	// RadixSize represents the branching factor
	// of the trie. Should this value increase,
	// you would also have to update the index type
	// from "byte" to whatever holds this larger value.
	RadixSize = 256
)

// A Trie is an implementation of a DLB with a radix branching
// factor of RadixSize.
type Trie struct {
	// The root node is the only node
	// in the trie without a byte stored on it.
	root Node
	// Empty is true if the empty string was added to the DLB.
	empty bool
}

// NewDLBTrie returns a new, empty trie.
func NewDLBTrie() *Trie {
	return &Trie{
		empty: false,
		root:  *NewNode(byte(0)),
	}
}

// Add will intern this string within the trie.
func (trie *Trie) Add(s string) {
	if s == "" {
		trie.empty = true
	}

	asBytes := []byte(s)
	trie.root.Add(asBytes)
}
