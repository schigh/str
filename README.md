[![Build Status](https://travis-ci.org/schigh/str.svg?branch=master)](https://travis-ci.org/schigh/str)
[![GoDoc](https://godoc.org/github.com/schigh/str?status.svg)](https://godoc.org/github.com/schigh/str)

# str
The missing Golang Strings Library

##### Yo

This is a work in progress.  Here are the proposed bits thusfar:

* Pad - left or right pad a string with another string (done)
* Substring - get a substring based on start index and length (done)
* MD5 - get MD5 string hash of a string (done)
* SHA1 - get SHA1 string hash of a string (done)
* SHA256 - get SHA256 string hash of a string (done)
* CRC32 - Get CRC32 checksum of a string (done)
* CRC64 - get CRC64 checksum of a string
* Wrap - get a non-breaking word-wrapped string of a long string
* Currency - get a currency-formatted string
* HMAC - get an HMAC hash of a string (done)
* Token - generate a token based on preferences

Some of these utilities will require pre-configuration via options, 
e.g. HMAC options, TokenOptions, etc.


Once these are all added in, I'm sure I'll add more

