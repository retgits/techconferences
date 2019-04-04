# techconferences

[![Godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/retgits/techconferences)

[confs.tech](https://confs.tech/) is an awesome open-source and crowd-sourced conference website that helps you find your next tech conference. Sometimes you just need to get some of the data in your Go app, without getting using a browser.

## Usage

Get today's holidays

```go
import "github.com/retgits/techconferences"

conferences, err := techconferences.GetConferences(techconferences.DevOps, 2019)
if err != nil {
    fmt.Printf("Oh noes, an error occured: %s", err.Error())
}

for idx := range conferences {
    fmt.Printf("Let's go to %s\n", conferences[idx].Name)
}
```

## License

See the [LICENSE](./LICENSE) file in the repository

## Acknowledgements

A most sincere thanks to the team of [confs.tech](https://confs.tech/), for building and maintaining such a large dataset of conferences!

_This package is not endorsed by confs.tech_
