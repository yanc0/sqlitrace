# sqlitrace

sqlite trace callback experiment with Golang

## Description

Create a library callable by SQLite in Go (extension) and create a small app that load this extension a prints performance stats every seconds.

## Setup

* Download [sqlean/stats.so](https://github.com/nalgeon/sqlean) extension
* Install libsqlite3-dev

## Build

    go build -buildmode=c-shared -o trace.so
    go run example/main.go

## Links

- https://github.com/nalgeon/sqlean
- https://pkg.go.dev/cmd/cgo
- https://www.sqlite.org/c3ref/trace_v2.html
