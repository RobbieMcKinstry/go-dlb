package dlb

// An UnorderedStore has a linear time lookup by using
// a slice to store the nodes.
type UnorderedStore struct {
	nodes []*Node
}

// NewUnorderedStore returns a new instance of UnorderedStore
func NewUnorderedStore() *UnorderedStore {
	return &UnorderedStore{
		nodes: make([]*Node, 0, RadixSize),
	}
}

// Fetch will return the node in the store containing this byte,
// if it exists, and nil otherwise.
func (store *UnorderedStore) Fetch(b byte) *Node {
	return store.nodes[b]
}

// Add will add this node to the store for later retrieval.
func (store *UnorderedStore) Store(n *Node) {
	store.nodes[n.b] = n
}

var _ NodeStore = NewUnorderedStore()
