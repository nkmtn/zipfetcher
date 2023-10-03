package zipfetcher

import "time"

type ZipProvider interface {
	GetLastModificationDate() (time.Time, error)
	GetZips() ([]ZipCode, error)
}

type ZipCode struct {
	Code       string // 5 digits zip Code
	State      string // two-letter State designation
	City       string
	LocaleName string // useful for identifying military
}
