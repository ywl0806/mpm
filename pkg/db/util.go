package db

import (
	"bytes"
	"encoding/gob"
	"log"
)

func Serialize(data interface{}) []byte {
	buf := bytes.NewBuffer(nil)
	gob.Register(data)
	err := gob.NewEncoder(buf).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

func Deserialize(data []byte, v interface{}) error {
	gob.Register(v)
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(v)
}
