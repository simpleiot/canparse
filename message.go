package canparse

// Message represents a CAN bus frame that has an id and data, chich is split into signals.
type Message struct {
	Id uint64
	Name string
	Length int
	Signals []Signal
}
