package zipfetcher

// ZipColumn used to allocate which column in each page contains zips
var ZipColumn = map[string]int{
	"ZIP_DETAIL":        4,
	"Unique_ZIP_DETAIL": 4,
	"Other":             5,
}

// StateColumn used to allocate which column in each page contains states
// Pay attention military zones data don't contain real physical location
var StateColumn = map[string]int{
	"ZIP_DETAIL":        8,
	"Unique_ZIP_DETAIL": 8,
	"Other":             12,
}

// CityColumn used to allocate which column in each page contains city
// Pay attention military zones data don't contain real physical location
var CityColumn = map[string]int{
	"ZIP_DETAIL":        7,
	"Unique_ZIP_DETAIL": 7,
	"Other":             11,
}

// LocaleNameColumn used to allocate which column in each page contains locale name
var LocaleNameColumn = map[string]int{
	"ZIP_DETAIL":        5,
	"Unique_ZIP_DETAIL": 5,
	"Other":             7,
}
