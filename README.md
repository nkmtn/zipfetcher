# zipfetcher

zipfetcher is a Golang library for fetching USA ZIP codes and some data about them from an official source - The USPS postal service.

zipfetcher allows you to check resource updates by date and receive data in the format of an array of ZipCode struct
```
type ZipCode struct {
    code       string // 5 digits zip code
    state      string // two-letter state designation
    city       string // city name
    localeName string // useful for filtering out military zones
}
```

## TODO: add usage example

## A little about the source:
The USPS updates xls table approximately every month: https://postalpro.usps.com/ZIP_Locale_Detail .
**However, they don't provide explanation for data.** 
The table contains 3 sheets with different column sets. None of them contains classification column. 
I guess the codes are allocated on the page as follows:
* "ZIP_DETAIL" sheet contains Standard and Post office box-only ZIPs
* "Unique_ZIP_DETAIL" - Unique ZIPs such as governmental agencies, universities, businesses, or buildings receiving sufficiently high volumes of mail and also Santa Claus code is here
* "Others" - Military and Post office box-only ZIPs

All sheets contain: ZIP code, state, city and locale name. For my purposes this is enough.
I use locale name to identify Military codes (APO or DPO value).


