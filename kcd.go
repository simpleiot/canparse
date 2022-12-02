package canparse

import (
	"encoding/xml"
)

type KcdDatabase struct {
	XMLName  xml.Name `xml:"NetworkDefinition"`
	KcdBusses []KcdBus `xml:"Bus"`
}

type KcdBus struct {
	XMLName xml.Name `xml:"Bus"`
	Name string `xml:"name,attr"`
	KcdMessages []KcdMessage `xml:"Message"`
}

type KcdMessage struct {
	XMLName xml.Name `xml:"Message"`
	Id string`xml:"id,attr"`
	Name string `xml:"name,attr"`
	Length int `xml:"length,attr"`
	Format string `xml:"format,attr"`
	KcdSignals []KcdSignal `xml:"Signal"`
}

type KcdSignal struct {
	XMLName xml.Name `xml:"Signal"`
	Name string `xml:"name,attr"`
	Offset int `xml:"offset,attr"` 			// KCD signal offset maps to signal start
	Length int `xml:"length,attr"`
	KcdValue KcdValue `xml:"Value"`
}

type KcdValue struct {
	XMLName xml.Name `xml:"Value"`
	Unit string `xml:"unit,attr"`
	Slope float64 `xml:"slope,attr"` 		// KCD value slope maps to signal scale
	Intercept float64 `xml:"intercept,attr"` 	// KCD value intercept maps to signal offset
}
