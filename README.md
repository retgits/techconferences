# techconferences

[![Go Report Card](https://goreportcard.com/badge/github.com/retgits/techconferences?style=flat-square)](https://goreportcard.com/report/github.com/retgits/techconferences)
[![Godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/retgits/techconferences)
![GitHub](https://img.shields.io/github/license/retgits/techconferences?style=flat-square)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/retgits/techconferences?sort=semver&style=flat-square)

> A Go library for confs.tech, an open-source and crowd-sourced conference website to find your next conference

[confs.tech](https://confs.tech/) is an awesome open-source and crowd-sourced conference website that helps you find your next tech conference. Sometimes you just need to get some of the data in your Go app, without getting using a browser.

## Prerequisites

[Go (at least Go 1.12)](https://golang.org/dl/)

## Installation

Using `go get`

```bash
go get github.com/retgits/techconferences
```

## Usage

Get today's holidays

```go
import tc "github.com/retgits/techconferences/v2"

conferences, err := tc.GetConferences(tc.DevOps, 2019)
if err != nil {
    fmt.Printf("Oh noes, an error occured: %s", err.Error())
}

for idx := range conferences {
    fmt.Printf("Let's go to %s\n", conferences[idx].Name)
}
```

## Acknowledgements

A most sincere thanks to the team of [confs.tech](https://confs.tech/), for building and maintaining such a large dataset of conferences!

_This package is not endorsed by confs.tech_

## Contributing

[Pull requests](https://github.com/retgits/techconferences/pulls) are welcome. For major changes, please open [an issue](https://github.com/retgits/techconferences/issues) first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

See the [LICENSE](./LICENSE) file in the repository
