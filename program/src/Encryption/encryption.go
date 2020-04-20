package Encryption

import (
	"math/rand"
)

type Table struct {
	// Encode is ascii
	// Decode is byte
	encode [256]byte
	decode [256]byte
}

func NewEncryptionTable() *Table {
	var values [256]byte
	keys := generateKeys()
	for index, key := range keys {
		keys[index] = key
		values[key] = byte(index)
	}
	return &Table{keys, values}
}
func NewEmptyEncryptionTable() *Table {
	var values [256]byte
	var keys [256]byte
	for i := 0; i < 256; i++ {
		values[i] = byte(0)
		keys[i] = byte(0)
	}
	return &Table{keys, values}

}

// key 0 - 255
// value : random ascii number
// key = value is not allowed
func generateKeys() [256]byte {
	randomArray := rand.Perm(256)
	var keys [256]byte
	for index, value := range randomArray {
		if index == value {
			return generateKeys()
		} else {
			keys[index] = byte(value)
		}
	}
	return keys
}
func (t *Table) Encode(keys []byte) []byte {
	// Encode this byte array need jiacheng to change
	values := make([]byte, len(keys))
	for index, value := range keys {
		values[index] = t.encode[value]
	}
	return values
}

func (t *Table) Decode(values []byte) []byte {
	//Decode this byte array need jiacheng to change
	keys := make([]byte, len(values))
	for index, value := range values {
		keys[index] = t.decode[value]
	}
	return keys
}

func (t *Table) GetEncodeArr() []byte {
	return t.encode[:]
}

func (t *Table) GetDecodeArr() []byte {
	return t.decode[:]
}
func (t *Table) SetEncodeArr(copy []byte) {
	for i := 0; i < 256; i++ {
		t.encode[i] = copy[i]
	}
}
func (t *Table) SetDecodeArr(copy []byte) {
	for i := 0; i < 256; i++ {
		t.decode[i] = copy[i]
	}
}
