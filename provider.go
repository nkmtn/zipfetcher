package zipfetcher

import "time"

type ZipProvider interface {
	GetLastModificationDate() (time.Time, error)
	GetZips() ([]ZipCode, error)
}

type ZipCode struct {
	code       string // 5 digits zip code
	state      string // two-letter state designation
	city       string
	localeName string // useful for identifying military
}
