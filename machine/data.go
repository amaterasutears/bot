package machine

// Data represents a single data item.
type Data struct {
	Key   string
	Value any
}

// DataMap holds a collection of data items keyed by string.
type DataMap map[string]any
