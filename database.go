package canparse

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"github.com/pkg/errors"
)

type Database struct {
	XMLName  xml.Name `xml:"NetworkDefinition"`
	Busses []Bus `xml:"Bus"`
}

func ReadKcd(filePathKcd string, db *Database) error {
	kcdFile, err := os.Open(filePathKcd)
	if err != nil {
		return errors.Wrap(err, "Canparse, error opening file at supplied path:")
	}
	defer kcdFile.Close()
	
	kcdData, err := ioutil.ReadAll(kcdFile)
	if err != nil {
		return errors.Wrap(err, "Canparse, error reading file contents:")
	}

	err = xml.Unmarshal(kcdData, db)
	if err != nil {
		return errors.Wrap(err, "Canparse, error parsing KCD data:")
	}

	// Now the database should be populated with the data of the KCD file
	return nil
}
