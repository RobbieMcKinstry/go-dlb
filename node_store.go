package dlb

// A NodeStore is anything that can fetch a node
// from it. It should be implemented as an unordered
// slice of pointers for linear search, or a binary search
// tree or red-black tree for very fast search.
type NodeStore interface {
	Fetch(byte) *Node
	Store(*Node)
}
