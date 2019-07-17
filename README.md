# gosdk

Manages a symlink to a specified version of a go get downloaded Go SDK.

You can download Go SDKs using commands such as `go get golang.org/dl/go1.12.6` followed by `go1.12.6 download` to
download and install new versions of Go. This tool makes it easy to manage a symlink to a specific downloaded version of
the SDKs without having to use the `go1.12.6` command directly or preconfiguring the environment.

## Prerequisites

Go is required. See https://golang.org/dl/ for installation instructions.

`$HOME/.local/bin` should be added to your PATH if it isn't already as `gosdk` will put the symlink here.

## Installation

Run `go get -u github.com/taisph/gosdk/cmd/gosdk` to install the latest version of the command.

## Usage

List currently installed Go SDKs using `gosdk list` or just `gosdk`.

Example:

```
$ gosdk list
go1.12.6
go1.11.11
go1.11.10
go1.11.9
go1.11.6
go1.11.5
```

Change the Go SDK using `gosdk set go1.12.6`.

Example:

```
$ go version
go version go1.11.10 linux/amd64
$ gosdk set go1.12.6
$ go version
go version go1.12.6 linux/amd64
```
