# zipfetcher

>Is a Golang library for fetching USA ZIP codes and some data about them from an official source - The USPS postal service.

Allows you to check resource updates by date and receive data in the format of an array of ZipCode struct
```
type ZipCode struct {
    Code       string // 5 digits zip code
    State      string // two-letter state designation
    City       string // city name
    LocaleName string // useful for filtering out military zones
}
```

## Usage examples
Running examples, both of them use USPS as a source:
* first way, used default provider
```
    zf := zipfetcher.Create()
	zips, err := zf.GetAllZips()
	if err != nil {
		// ...
	}
	fmt.Println(zips)
```
* second way, explicitly declare the provider
```
    zf := zipfetcher.Create(zipfetcher.WithProvider(zipfetcher.CreateUspsProvider()))
	zips, err := zf.GetAllZips()
	if err != nil {
		// ...
	}
	fmt.Println(zips)
```

## A little about the USPS:
The USPS updates xls table approximately every month: https://postalpro.usps.com/ZIP_Locale_Detail .
**However, they don't provide explanation for data.** 
The table contains 3 sheets with different column sets. None of them contains classification column. 
I guess the codes are allocated on the page as follows:
* "ZIP_DETAIL" sheet contains Standard and Post office box-only ZIPs
* "Unique_ZIP_DETAIL" - Unique ZIPs such as governmental agencies, universities, businesses, or buildings receiving sufficiently high volumes of mail and also Santa Claus code is here
* "Others" - Military and Post office box-only ZIPs

All sheets contain: ZIP code, state, city and locale name. For my purposes this is enough.
I use locale name to identify Military codes (APO or DPO value).


