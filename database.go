package canparse

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Database represents a generic CAN database with multiple CAN busses.
type Database struct {
	Busses []Bus
}

// ReadKcd erases any information in the database and populates the database
// with the information parsed from the KCD file.
func (db *Database) ReadKcd(filePathKcd string) error {

	// Make sure the database is clean
	db.Busses = []Bus{}

	kcdFile, err := os.Open(filePathKcd)
	if err != nil {
		return errors.Wrap(err, "Canparse, error opening file at supplied path")
	}
	defer kcdFile.Close()

	kcdData, err := ioutil.ReadAll(kcdFile)
	if err != nil {
		return errors.Wrap(err, "Canparse, error reading file contents")
	}

	kcdDb := &KcdDatabase{}
	err = xml.Unmarshal(kcdData, kcdDb)
	if err != nil {
		return errors.Wrap(err, "Canparse, error parsing KCD data")
	}

	// Confert KCD database to generic database
	for i, b := range kcdDb.KcdBusses {
		db.Busses = append(db.Busses, Bus{})
		for _, m := range b.KcdMessages {

			// Copy over the message contents
			var parsedMsg Message
			parsedMsg.Name = m.Name
			// Parse the string representing a hexidecimal number to a uint64
			hexStr := strings.Replace(m.Id, "0x", "", -1)
			hexStr = strings.Replace(hexStr, "0X", "", -1)
			id, err := strconv.ParseUint(hexStr, 16, 32)
			if err != nil {
				return errors.Wrap(err, "Canparse, error parsing hexidecimal from KCD file")
			}
			parsedMsg.Id = uint32(id)
			// Copy over the signals of the message
			for _, s := range m.KcdSignals {

				// Copy the signal contents
				var parsedSig Signal
				parsedSig.Name = s.Name
				parsedSig.Start = s.Offset
				parsedSig.Length = s.Length
				parsedSig.Scale = s.KcdValue.Slope
				parsedSig.Offset = s.KcdValue.Intercept
				parsedSig.Unit = s.KcdValue.Unit
				parsedMsg.Signals = append(parsedMsg.Signals, parsedSig)
			}
			db.Busses[i].Messages = append(db.Busses[i].Messages, parsedMsg)
		}
	}

	// Now the database should be populated with the data of the KCD file
	return nil
}
