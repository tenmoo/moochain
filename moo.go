package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Moo represents a data block in the moochain
type Moo struct {
	DataHash  string
	Timestamp string
	PrevHash  string
	Hash      string
}

// Validate Moo by checking and comparing the hash of the previous moo
func isMooValid(newMoo, oldMoo Moo) bool {
	if oldMoo.Hash != newMoo.PrevHash {
		return false
	}

	if calculateHash(newMoo) != newMoo.Hash {
		return false
	}

	return true
}

// Calculate SHA256 hash
func calculateHash(Moo Moo) string {
	record := Moo.DataHash + Moo.Timestamp + Moo.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// Create a new Moo using previous Moo's hash
func generateNewMoo(oldMoo Moo, data string) Moo {
	var newMoo Moo
	t := time.Now()
	newMoo.DataHash = data
	newMoo.Timestamp = t.String()
	newMoo.PrevHash = oldMoo.Hash
	newMoo.Hash = calculateHash(newMoo)
	return newMoo
}
