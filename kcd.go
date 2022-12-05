package canparse

import (
	"encoding/xml"
)

// KcdDatabase mirrors the generic Database type with added XML decoding tags for the
// KCD format.
type KcdDatabase struct {
	XMLName   xml.Name `xml:"NetworkDefinition"`
	KcdBusses []KcdBus `xml:"Bus"`
}

// KcdBus mirrors the generic Bus type with added XML decoding tags for the KCD
// format.
type KcdBus struct {
	XMLName     xml.Name     `xml:"Bus"`
	Name        string       `xml:"name,attr"`
	KcdMessages []KcdMessage `xml:"Message"`
}

// KcdMessage mostly mirrors the generic Message type with added XML decoding tags
// for the KCD format. The Format field is part of the KCD standard but is not yet
// included in the generic Message type.
type KcdMessage struct {
	XMLName    xml.Name    `xml:"Message"`
	Id         string      `xml:"id,attr"`
	Name       string      `xml:"name,attr"`
	Length     int         `xml:"length,attr"`
	Format     string      `xml:"format,attr"`
	KcdSignals []KcdSignal `xml:"Signal"`
}

// KcdSignal mostly mirrors the generic Signal type with added XML decoding tags
// for the KCD format. It contains a sub struct Value in adherence with the KCD
// standard.
type KcdSignal struct {
	XMLName  xml.Name `xml:"Signal"`
	Name     string   `xml:"name,attr"`
	Offset   int      `xml:"offset,attr"` // KCD signal offset maps to signal start
	Length   int      `xml:"length,attr"`
	KcdValue KcdValue `xml:"Value"`
}

// KcdValue adhering to the KCD standard.
type KcdValue struct {
	XMLName   xml.Name `xml:"Value"`
	Unit      string   `xml:"unit,attr"`
	Slope     float64  `xml:"slope,attr"`     // KCD value slope maps to signal scale
	Intercept float64  `xml:"intercept,attr"` // KCD value intercept maps to signal offset
}
