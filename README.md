# Golang time/date Object

[![Build Status][build_status_image]][build_status] 
[![GoDoc][godoc_image]][godoc] 
[![Go Report Card][goreport_image]][goreport] 
[![Software License][license_image]](LICENSE.md)

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


[build_status]:       https://travis-ci.org/mantyr/newtime
[build_status_image]: https://travis-ci.org/mantyr/newtime.svg?branch=master

[godoc]:              http://godoc.org/github.com/mantyr/newtime
[godoc_image]:        https://godoc.org/github.com/mantyr/newtime?status.png

[goreport]:           https://goreportcard.com/report/github.com/mantyr/newtime
[goreport_image]:     https://goreportcard.com/badge/github.com/mantyr/newtime

[license_image]:      https://img.shields.io/badge/license-MIT-brightgreen.svg

