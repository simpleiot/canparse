package canparse

// Message represents a CAN bus frame that has an id and data which is split
// into signals.
type Message struct {
	Id      uint32
	Name    string
	Length  int
	Signals []Signal
}
