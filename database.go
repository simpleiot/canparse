package canparse

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"go.einride.tech/can/pkg/dbc"
)

// Database represents a generic CAN database with multiple CAN busses.
type Database struct {
	Busses []Bus
}

// Clean empties the database
func (db *Database) Clean() {
	db.Busses = []Bus{}
}

// Read takes the path to a .dbc or .kcd CAN database file and parses it. It then
// appends the information to the generic Database.
func (db *Database) Read(filePathDb string) error {
	file, err := os.Open(filePathDb)
	if err != nil {
		return errors.Wrap(err, "Canparse, error opening file at supplied path")
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return errors.Wrap(err, "Canparse, error reading file contents")
	}

	err = db.ReadBytes(bytes, filePathDb)
	if err != nil {
		return err
	}
	return nil
}

// ReadBytes takes bytes containing data in .dbc or .kcd format and parses it. It
// then appends the information to the generic Database.
func (db *Database) ReadBytes(bytes []byte, filePathDb string) error {
	log.Println(filepath.Ext(filePathDb))
	switch filepath.Ext(filePathDb) {
	case ".dbc":
		err := db.ReadDbc(bytes, filePathDb)
		if err != nil {
			return err
		}
	case ".kcd":
		err := db.ReadKcd(bytes)
		if err != nil {
			return err
		}
	default:
		return errors.New(fmt.Sprintf("canparse: invalid CAN database file extension: %v", filePathDb))
	}
	return nil
}

// ReadDbc takes bytes containing data in .dbc format and parses it. It
// then appends the information to the generic Database.
func (db *Database) ReadDbc(bytes []byte, filePathDb string) error {
	parser := dbc.NewParser(filePathDb, bytes)
	err := parser.Parse()
	if err != nil {
		return errors.Wrap(err, "Canparse, error parsing DBC data")
	}
	bus := &Bus{
		Name: filePathDb,
	}
	defs := parser.Defs()
	for _, d := range defs {
		keyword := string(bytes[d.Position().Offset : d.Position().Offset+3])
		log.Println(d.Position(), keyword)

		// Code to detect dbc.MessageDef keyword and somehow parse the rest of the
		// data as that???
		if keyword != string(dbc.KeywordMessage) {
			continue
		}
		//_ = d.Name

		msg := Message{}
		bus.Messages = append(bus.Messages, msg)
	}

	return nil
}

// ReadKcd takes bytes containing data in .kcd format and parses it. It
// then appends the information to the generic Database.
func (db *Database) ReadKcd(bytes []byte) error {

	kcdDb := &KcdDatabase{}
	err := xml.Unmarshal(bytes, kcdDb)
	if err != nil {
		return errors.Wrap(err, "Canparse, error parsing KCD data")
	}

	// Append parsed KCD data to master Database
	for _, b := range kcdDb.KcdBusses {
		parsedMessages := []Message{}
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
			parsedMessages = append(parsedMessages, parsedMsg)
		}
		db.Busses = append(db.Busses, Bus{Name: b.Name, Messages: parsedMessages})
	}

	// Now the database should be populated with the data of the KCD file
	return nil
}
