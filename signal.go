package canparse

import (
	"encoding/xml"
)

type ByteOrder string

const (
	LittleEndian ByteOrder = "LittleEndian"
	BigEndian ByteOrder = "BigEndian"
)

// Signal represents a meaningful part of a CAN message.
type Signal struct {
	XMLName xml.Name `xml:"Signal"`
	Name string `xml:"name,attr"`
	Start int
	Length int
	ByteOrder ByteOrder
	IsSigned bool
	Scale float64
	Offset float64 `xml:"offset,attr"`
	Minimum float64
	Maximum float64
	Unit string
}
