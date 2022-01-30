# Go Cli tool

To run anywhere doing some simple stuffs

## Requirements

```
$ go version
go version go1.17.5 linux/amd64

$ goreleaser --version
goreleaser version 1.4.1
...
```

## Getting started

```
go run .
go build
./dice-cli

goreleaser init    # already done
goreleaser release --snapshot --rm-dist
```


## Release Assets

Each release will have assets attached (compressed binaries, checksums and source code).

The cli binaries will be available for platform.

Example for linux:

```
wget https://github.com/Duwab/go-cli-tool/releases/download/v0.0.1/go-cli-tool_0.0.1_Linux_x86_64.tar.gz
sha256sum go-cli-tool_0.0.1_Linux_x86_64.tar.gz

tar xzvf go-cli-tool_0.0.1_Linux_x86_64.tar.gz
./dice-cli
```

## Greetings

Thanks to Mason Egger (@DigitalOcean) for [this Live tutorial](https://www.youtube.com/watch?v=PP4CVvgLXrU)