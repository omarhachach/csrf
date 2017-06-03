## CSRF - A toolset for CSRF middleware
[![Travis branch](https://img.shields.io/travis/rust-lang/rust/master.svg?style=flat-square)](https://travis-ci.org/omar-h/csrf)
![GitHub tag](https://img.shields.io/github/release/omar-h/csrf.svg?style=flat-square)
[![Report Card](https://img.shields.io/badge/report%20card-a%2B-c0392b.svg?style=flat-square)](https://goreportcard.com/report/github.com/omar-h/csrf)
![Powered By](https://img.shields.io/badge/powered%20by-go-blue.svg?style=flat-square)
[![Docs](https://img.shields.io/badge/docs-reference-9b59b6.svg?style=flat-square)](https://godoc.org/github.com/omar-h/csrf)
[![License](https://img.shields.io/badge/license-MIT%20License-1abc9c.svg?style=flat-square)](https://github.com/omar-h/csrf/blob/master/LICENSE.txt)

Logic behind CSRF token creation and validation.

Inspired by [pillarjs/csrf](https://github.com/pillarjs/csrf). Read [Understanding-CSRF](https://github.com/pillarjs/understanding-csrf) for more information on CSRF.

### Install
```Bash
$ go get -u github.com/omar-h/csrf
```

## Example
This is an example of how to initiliaze and use the package:
```Go
package main

import (
        "fmt"
        
        "github.com/omar-h/csrf"
)

func main() {
        const secret = "erHUnxuhBMRIsVB1LfqmiWCgB83ZEerH"
        CSRF := csrf.New(csrf.Options{
                // Secret should persist over program restart.
                Secret: secret,
                SaltLen: 16,
        })
        
        salt := CSRF.GenerateSalt()
        token := CSRF.GenerateToken(salt)
        
        // Print the secret, a random salt and the token generated from them.
        fmt.Println("Secret: ", secret)
        fmt.Println("Salt: ", salt)
        fmt.Println("Token: ", token)
        
        // Returns true
        CSRF.Verify(token)
}
```

## License
CSRF is licensed under the [MIT License](https://github.com/omar-h/csrf/blob/master/LICENSE.txt).
