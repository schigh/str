[![Build Status](https://travis-ci.org/schigh/str.svg?branch=master)](https://travis-ci.org/schigh/str)
[![GoDoc](https://godoc.org/github.com/schigh/str?status.svg)](https://godoc.org/github.com/schigh/str)


# str
The missing Golang Strings Library

#### Overview

Many of these functions serve as Go analogues to some of the better PHP string 
functions (yes, they do exist).

Some things of note:
* All functions output a single string.
    * If an error is encountered, or if a function cannot do its job, it outputs an empty string.
* There are no dependencies other than elements from the standard library


#### Usage

```go
import "github.com/schigh/str"
```
##### Pad
Pads a string left or right, similar to PHP's `str_pad`
```go
sku := str.Pad("1234", "0", 8, str.PadLeft)
// 00001234
```
Your pad string may be longer than one character
```go
id := str.Pad("1234", "ABC", 12, str.PadRight)
// 1234ABCABCABC
```
Note that for multi-character pads, if the padding doesn't land cleanly at the desired length, `str.Pad` will overshoot the desired length until the padding is complete.


##### Substring
Gets a substring from a string.
```go
month := str.Substring("2017-04-24", 5, 2)
// 04

day := str.Substring("2017-04-24", -2, 2)
// 24
```
You can use a negative value for the start, which, like PHP's [`substr`](http://us3.php.net/manual/en/function.substr.php) function, will pull from the back of the source string.

##### Common Hash Functions
The following hash functions return the hex-encoded hash of an input string
```go
data := "This is a string"

md5 := str.MD5(data)
// 1710528bf976601a5d203cbc289e1a76

sha1 := str.SHA1(data)
// bd82fb0e81ee9f15f5929e0564093bc9f8015f1d

sha256 := str.SHA256(data)
// fe29bfd7e19d2a352cd9bfaa21176d521b874969dd57e8c72c8668fc8fd8f4fc
```

##### HMAC


* CRC32 - Get CRC32 checksum of a string (done)
* CRC64 - get CRC64 checksum of a string (done)
* Wrap - get a non-breaking word-wrapped string of a long string (done)
* Currency - get a currency-formatted string
* HMAC - get an HMAC hash of a string (done)
* Token - generate a token based on preferences

Some of these utilities will require pre-configuration via options, 
e.g. HMAC options, TokenOptions, etc.


Once these are all added in, I'm sure I'll add more

