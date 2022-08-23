package parse

type Node struct {
	Kind  NodeKind
	Key   *Node // NdKVの時にNdStringが入る
	Value *Node // NdKVの時に色々な種類のノードが入る

	Number float64 // NdNumberの時にデータが入る
	Str    string  // NdStringの時にデータが入る

	Children []*Node // NdObjectの時はNdKVが、NdArrayの時は色々な種類のノードが入る
}

func NewNode(kind NodeKind, key, value *Node, num float64, str string, children []*Node) *Node {
	return &Node{
		Kind:     kind,
		Key:      key,
		Value:    value,
		Number:   num,
		Str:      str,
		Children: children,
	}
}

func NewKVNode(key, value *Node) *Node {
	return NewNode(NdKV, key, value, 0, "", nil)
}

func NewStringNode(str string) *Node {
	return NewNode(NdString, nil, nil, 0, str, nil)
}

func NewNumberNode(num float64) *Node {
	return NewNode(NdNumber, nil, nil, num, "", nil)
}

func NewTrueNode() *Node {
	return NewNode(NdTrue, nil, nil, 0, "", nil)
}

func NewFalseNode() *Node {
	return NewNode(NdFalse, nil, nil, 0, "", nil)
}

func NewNullNode() *Node {
	return NewNode(NdNULL, nil, nil, 0, "", nil)
}

func NewObjectNode(kvPairs []*Node) *Node {
	return NewNode(NdObject, nil, nil, 0, "", kvPairs)
}

func NewArrayNode(children []*Node) *Node {
	return NewNode(NdArray, nil, nil, 0, "", children)
}
