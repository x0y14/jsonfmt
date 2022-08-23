package parse

type NodeKind int

const (
	NdILLEGAL NodeKind = iota

	NdKV // key and value

	NdString // "..." only value
	NdNumber // 123
	NdTrue   // true
	NdFalse  // false
	NdNULL   // null
	NdObject // {...}
	NdArray  // [...]
)
