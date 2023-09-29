package zipfetcher

import (
	"github.com/itchyny/timefmt-go"
	"os"
)

type ZipFetcher struct {
	downloader *Downloader
	parser     *Parser
}

func Create() *ZipFetcher {
	return &ZipFetcher{
		downloader: CreateDownloader(),
		parser:     CreateParser(),
	}
}

// GetAllZips return all zips
func (zf *ZipFetcher) GetAllZips() ([]ZipCode, error) {
	err := zf.downloader.parseSourcePage()

	defer os.RemoveAll(zf.downloader.XlsPath)
	err = zf.downloader.downloadXls()
	if err != nil {
		return []ZipCode{}, err
	}

	return zf.parser.ExtractZipsInfo(zf.downloader.XlsPath)
}

// CheckIfModifiedSince checking if data was modified after the date
func (zf *ZipFetcher) CheckIfModifiedSince(date string) (bool, error) {
	err := zf.downloader.parseSourcePage()
	if err != nil {
		return false, err
	}

	clientDate, err := timefmt.Parse(date, "%Y-%m-%d")
	if err != nil {
		return false, err
	}
	return clientDate.Before(zf.downloader.PageUpdateDate), nil
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
