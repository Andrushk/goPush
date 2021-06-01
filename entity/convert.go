package entity

import (
"encoding/json"
"io"
)

func FromByte(reader io.Reader, entity interface{}) {
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(entity); err != nil {
		panic("error converting data!")
	}
}

func ToByte(v interface{}) (result []byte) {
	result, err := json.Marshal(v)
	if err != nil {
		panic("error converting data!")
	}
	return
}