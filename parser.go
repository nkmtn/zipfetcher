package zipfetcher

import (
	"fmt"
	"github.com/pbnjay/grate"
	"github.com/pbnjay/grate/xls"
)

type Parser struct{}

func CreateParser() *Parser {
	return &Parser{}
}

type ZipCode struct {
	code       string // 5 digits zip code
	state      string // two-letter state designation
	city       string
	localeName string // useful for identifying military
}

// ExtractZipsInfo to extract info from all xls sheets
func (p *Parser) ExtractZipsInfo(path string) ([]ZipCode, error) {
	source, err := xls.Open(path)
	if err != nil {
		return nil, fmt.Errorf("can't open the file: %s", err.Error())
	}

	l, err := source.List()
	if err != nil {
		return nil, fmt.Errorf("can't get sheets from %s: %s", path, err.Error())
	}

	m := make(map[string]ZipCode) // file contains duplicates, map is used for uniqueness
	for _, table := range l {
		err = extractZipsFromSheet(source, table, m)
		if err != nil {
			return nil, err
		}
	}
	return extractMapValues(m), nil
}

// extract info from one xls sheet list
func extractZipsFromSheet(source grate.Source, table string, m map[string]ZipCode) error {
	collection, err := source.Get(table)
	if err != nil {
		return fmt.Errorf("can't get table %s: %s", table, err.Error())
	}

	collection.Next() // titles row
	collection.Next() // the first row with data
	for !collection.IsEmpty() {
		if collection.Err() != nil {
			return fmt.Errorf("reading error in table %s: %s", table, err.Error())
		}
		if collection.Strings()[ZipColumn[table]] == "" { // ignore empty tail
			break
		}

		row := collection.Strings()
		m[row[ZipColumn[table]]] = ZipCode{
			code:       row[ZipColumn[table]],
			state:      row[StateColumn[table]],
			city:       row[CityColumn[table]],
			localeName: row[LocaleNameColumn[table]],
		}
		collection.Next()
	}
	return nil
}

func extractMapValues(m map[string]ZipCode) []ZipCode {
	zips := make([]ZipCode, 0, len(m))
	for _, z := range m {
		zips = append(zips, z)
	}
	return zips
}
