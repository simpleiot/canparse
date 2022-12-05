package canparse

// Bus represents a generic CAN bus with a name and multiple CAN messages.
type Bus struct {
	Name     string
	Messages []Message
}
