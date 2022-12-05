package canparse

import (
	"errors"
	"log"

	"go.einride.tech/can"
)

func DecodeMessage(frame can.Frame, db *Database) (DecodedMsg, error) {
	var message Message
	for _, bus := range db.Busses {
		for _, msg := range bus.Messages {
			log.Println("canparse:", msg.Id, frame.ID)
			if msg.Id == frame.ID {
				message = msg
			}
		}
	}
	decoded := DecodedMsg{
		Name:    message.Name,
		Id:      message.Id,
		Signals: make([]DecodedSig, len(message.Signals)),
	}
	if decoded.Id == 0 {
		return decoded, errors.New("canparse: no matching ID in database")
	}
	for i, sig := range message.Signals {
		decoded.Signals[i].Name = sig.Name
		decoded.Signals[i].Unit = sig.Unit
		decoded.Signals[i].Value =
			frame.Data.UnsignedBitsLittleEndian(uint8(sig.Start), uint8(sig.Length))
	}
	return decoded, nil
}

type DecodedMsg struct {
	Name    string
	Id      uint32
	Signals []DecodedSig
}

type DecodedSig struct {
	Name  string
	Unit  string
	Value uint64
}
