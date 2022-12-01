package canparse

type ByteOrder string

const (
	LittleEndian ByteOrder = "LittleEndian"
	BigEndian ByteOrder = "BigEndian"
)

// Signal represents a meaningful part of a CAN message.
type Signal struct {
	Name string
	Start int
	Length int
	ByteOrder ByteOrder
	IsSigned bool
	Scale float64
	Offset float64
	Minimum float64
	Maximum float64
	Unit string
}
