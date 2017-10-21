# Golang time/date Object

[![Build Status](https://travis-ci.org/mantyr/newtime.svg?branch=master)][build_status] 
[![GoDoc](https://godoc.org/github.com/mantyr/newtime?status.png)][godoc] 
[![Go Report Card](https://goreportcard.com/badge/github.com/mantyr/newtime)][goreport] 
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](LICENSE.md)

This stable version

## Installation

    $ go get github.com/mantyr/newtime

## Example

```GO
package main

import (
    "github.com/mantyr/newtime"
    "fmt"
)

func main() {
    date := newtime.NewTime()
    fmt.Println(date) // "YYYY-mm-dd HH:ii:ss", this current time, "2016-04-01 04:06:36"

    fmt.Println(date.Format("yyyy/_long_month_/dd")) // "YYYY/Month/dd", example "2016/April/01"

    date.Parse("2016-04-01T00:12:36", "yyyy-mm-ddTHH:ii:ss")
    fmt.Println(date.Format("YYYY-mm-dd HH:ii:ss"))  // "2016-04-01 00:12:36"

    // More examples in time_test.go
}
```

## Author

[Oleg Shevelev][mantyr]

[mantyr]: https://github.com/mantyr


[build_status]: https://travis-ci.org/mantyr/newtime
[godoc]:        http://godoc.org/github.com/mantyr/newtime
[goreport]:     https://goreportcard.com/report/github.com/mantyr/newtime

