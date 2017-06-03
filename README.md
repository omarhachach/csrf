## CSRF - A toolset for CSRF middleware
[![Build Status](https://travis-ci.org/omar-h/csrf.svg?branch=master)](https://travis-ci.org/omar-h/csrf)
[![Go Report Card](https://goreportcard.com/badge/github.com/omar-h/csrf)](https://goreportcard.com/report/github.com/omar-h/csrf)
[![GoDoc](https://godoc.org/github.com/omar-h/csrf?status.svg)](https://godoc.org/github.com/omar-h/csrf)


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
