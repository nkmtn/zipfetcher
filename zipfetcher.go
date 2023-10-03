package zipfetcher

import (
	"github.com/itchyny/timefmt-go"
)

type ZipFetcher struct {
	provider ZipProvider
}

// default provider - UspsProvider
func Create(provider ...func(*ZipFetcher)) *ZipFetcher {
	zf := &ZipFetcher{
		provider: CreateUspsProvider(),
	}
	for _, p := range provider {
		p(zf)
	}
	return zf
}

func WithProvider(provider ZipProvider) func(*ZipFetcher) {
	return func(zf *ZipFetcher) {
		zf.provider = provider
	}
}

// GetAllZips return all zips
func (zf *ZipFetcher) GetAllZips() ([]ZipCode, error) {
	return zf.provider.GetZips()
}

// CheckIfModifiedSince checking if data was modified after the date
//
// date format: yyyy-mm-dd
func (zf *ZipFetcher) CheckIfModifiedSince(date string) (bool, error) {
	modificationDate, err := zf.provider.GetLastModificationDate()
	if err != nil {
		return false, err
	}

	clientDate, err := timefmt.Parse(date, "%Y-%m-%d")
	if err != nil {
		return false, err
	}
	return clientDate.Before(modificationDate), nil
}

// GetAllZipsIfModifiedSince return all zips if data was modified after the date.
//
// If not modified, would return empty list without error
func (zf *ZipFetcher) GetAllZipsIfModifiedSince(date string) ([]ZipCode, error) {
	isModified, err := zf.CheckIfModifiedSince(date)
	if err != nil {
		return []ZipCode{}, err
	}
	if isModified {
		return zf.GetAllZips()
	}
	return []ZipCode{}, nil
}
