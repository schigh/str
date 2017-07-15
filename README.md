[![Build Status](https://travis-ci.org/schigh/str.svg?branch=master)](https://travis-ci.org/schigh/str)
[![GoDoc](https://godoc.org/github.com/schigh/str?status.svg)](https://godoc.org/github.com/schigh/str)
[![Go Report Card](https://goreportcard.com/badge/github.com/schigh/str)](https://goreportcard.com/report/github.com/schigh/str)


# str
The missing Golang Strings Library

str makes some common string tasks a little simpler, so you can get on with writing your applications.

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

###### Pad
Pads a string with another string.  This behavior is similar to PHP's [`str_pad`](http://php.net/manual/en/function.str-pad.php)

```go
package main

import (
    "fmt"
    "github.com/schigh/str"
)

var partialNum = "1234"

func main() {
    fmt.Println(str.Pad(partialNum, "0", 8, str.PadLeft))
    // prints 00001234
    
    fmt.Println(str.Pad(partialNum, "0", 8, str.PadRight))
    // prints 12340000
    
    fmt.Println(str.Pad(partialNum, "abcd", 8, str.PadLeft))
    // prints abcd1234
    
    fmt.Println(str.Pad(partialNum, "abcd", 10, str.PadLeft))
    // prints abcdabcd1234
    
    // The pad will be repeated at least until the length is satisfied.  
    // In some cases, like above, it will overshoot.  
    // It will NEVER undershoot 
}
```

###### Substring
Gets a substring from a string.  This behavior is similar to PHP's [`substr`](http://php.net/manual/en/function.substr.php)

```go
package main

import (
	"fmt"
	"github.com/schigh/str"
)

var slug = "1234567890"

func main() {
	fmt.Println(str.Substring(slug, 0, 3))
	// prints 123
	
	fmt.Println(str.Substring(slug, 1, 3))
	// prints 234
	
	fmt.Println(str.Substring(slug, -3, 3))
	// prints 890
	
	fmt.Println(str.Substring(slug, -1, 3))
	// nonsense call, this prints 0
	
	fmt.Println(str.Substring(slug, -11, 3))
	// out of range, prints empty string
}
```

###### MD5
Returns the MD5 hash (hex-encoded string) of a supplied string.

```go
package main

import (
	"fmt"
	"github.com/schigh/str"
)

func main() {
    fmt.Println(str.MD5("我爱“走”"))
    // prints 502d97cf15ded2b612eb7caf580440a5
}
```

###### SHA1
Returns the SHA1 hash (hex-encoded string) of a supplied string.

```go
package main

import (
	"fmt"
	"github.com/schigh/str"
)

func main() {
    fmt.Println(str.SHA1("我爱“走”"))
    // prints ca23c160bcf89334a83bcad027c136381bf8cc2b
}
```

###### SHA256
Returns the SHA256 hash (hex-encoded string) of a supplied string.

```go
package main

import (
	"fmt"
	"github.com/schigh/str"
)

func main() {
    fmt.Println(str.SHA256("我爱“走”"))
    // prints 6ce570942bcdd21116e930a66dd59cfbf23b075d1796d87e93951de43fabae4b
}
```

###### HMAC
Returns the HMAC digest (hex-encoded) of a string with supplied options

_Note_:

If you don't set the default HMAC signing key, one will be generated for you, based on the hardware
address of the first network interface the program finds.

This is **not** meant to replace the security of a strong signing key, but it is better than nothing.

```go
package main

import (
	"fmt"
	"github.com/schigh/str"
)

var superSecretKey = "ILIKETURTLES!!!1"

func main() {
	options := &str.HMACOptions{
		KeyData: []byte(superSecretKey),
		DigestType: str.HMACDigestTypeSHA256, // see hmac.go for digest types
	}
	
	// This will set the HMAC options globally
	str.SetHMACOptions(options)
	fmt.Println(str.HMAC("我爱“走”"))
	// prints 271b03817aa1bea17b1ef447eec434751fb71d5b6273324fc181c10b18e61463
	
	// you can also just call HMACWithOptions directly
	fmt.Println(str.HMACWithOptions("我爱“走”", options))
	// prints 271b03817aa1bea17b1ef447eec434751fb71d5b6273324fc181c10b18e61463
}
```

###### CRC32
Returns the Cyclic Redundancy Check (CRC) 32-bit checksum of a string.

```go
package main

import (
	"fmt"
	"github.com/schigh/str"
)

func main() {
	options := &str.CRC32Options{
		PolynomialType: str.CRC32TypeK, // see crc32.go for polynomial types
	}
	
	// This will set the CRC32 options globally
	str.SetCRC32Options(options)
	fmt.Println(str.CRC32("我爱“走”"))
	// prints 2845165967
	
	// you can also just call HMACWithOptions directly
	fmt.Println(str.CRC32WithOptions("我爱“走”", options))
	// prints 2845165967
}
```

###### CRC64
Returns the Cyclic Redundancy Check (CRC) 64-bit checksum of a string.

```go
package main

import (
	"fmt"
	"github.com/schigh/str"
)

func main() {
	options := &str.CRC64Options{
		PolynomialType: str.CRC64TypeISO, // see crc64.go for polynomial types
	}
	
	// This will set the CRC64 options globally
	str.SetCRC64Options(options)
	fmt.Println(str.CRC64("我爱“走”"))
	// prints 17267613643596219338
	
	// you can also just call HMACWithOptions directly
	fmt.Println(str.CRC64WithOptions("我爱“走”", options))
	// prints 17267613643596219338
}
```

###### Wrap
Returns a string that is wrapped according to a specified behavior.  The wrap is delimited by a specified string (default is '\n')

```go
package main

import (
	"fmt"
	"github.com/schigh/str"
)

var longSentence = `This is a string that could very likely be broken up into more than one line.  Maybe we should do that.  Just a thought.`

func main() {
	options := &str.WrapOptions{
		Width: 24,
		LineBreak: "\n",
		Behavior: str.WrapAfterWord, // wrap after a word if it extends past line length
	}
	
	// set wordwrap options globally
	str.SetWrapOptions(options)
	
	fmt.Println(str.Wrap(longSentence))
	/*
	prints:
	This is a string that could
	very likely be broken up
	into more than one line.
	Maybe we should do that.
	Just a thought.
	*/
	
	options = &str.WrapOptions{
		Width: 24,
		LineBreak: "\n",
		Behavior: str.WrapBeforeWord, // wrap before the word if it will extend past the line length
	}
	
	fmt.Println(str.WrapWithOptions(longSentence, options))
	/*
	prints:
	This is a string that
	could very likely be
	broken up into more than
	one line. Maybe we
	should do that. Just a
	thought.
	*/
	
	options = &str.WrapOptions{
		Width: 24,
		LineBreak: "\n",
		Behavior: str.WrapLiteral, // wrap exactly at the line length, do not trim whitespace (this is good for base64-encoded lines e.g. RFC 4648)
	}
	
	fmt.Println(str.WrapWithOptions(longSentence, options))
	/*
	prints
	This is a string that co
	uld very likely be broke
	n up into more than one 
	line. Maybe we should do
	 that. Just a thought.
	*/
}
```

###### Token
Returns a token based generated from user options

```go
package main

import (
	"fmt"
	"github.com/schigh/str"
)

func main() {
	fmt.Println(str.Token())
	// prints default format token that matches ^[a-zA-Z0-9]{20}$
	// ex: 7wyxEyzgvsrEGdHcpbiU
	
	options := &str.TokenOptions{
		Length: 32,
		Charset: `10`,
		Prefix: ``,
	}
	
	str.SetTokenOptions(options)
	fmt.Println(str.Token())
	// prints something like 00011110000110010100011010011111
	
	options = &str.TokenOptions{
		Length: 6,
		Charset: `abcdef1234567890`,
		Prefix: `#`,
	}
	fmt.Println(str.TokenWithOptions(options))
	// prints something like #cb4db
}
```