package canparse

import (
	"encoding/xml"
)

type Bus struct {
	XMLName xml.Name `xml:"Bus"`
	Name string `xml:"name,attr"`
	Messages []Message `xml:"Message"`
}
