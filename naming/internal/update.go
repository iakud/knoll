package internal

type Operation uint8

const (
	// Add indicates a new address is added.
	Add Operation = iota
	// Delete indicates an existing address is deleted.
	Delete
)

type Update struct {
	Op       Operation
	Addr     string
	Metadata any
}
