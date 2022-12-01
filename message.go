package canparse

import (
	"encoding/xml"
)

// Message represents a CAN bus frame that has an id and data, chich is split into signals.
type Message struct {
	XMLName xml.Name `xml:"Message"`
	Id string`xml:"id,attr"`
	Name string `xml:"name,attr"`
	Length int `xml:"length,attr"`
	Signals []Signal `xml:"Signal"`
}
