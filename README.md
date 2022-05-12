# Golang binding for ladder
## Prerequisites
* [Ladder](https://github.com/frannecki/ladder)

## Usage
```sh
cd server
swig -go -cgo -c++ -intgosize 64 server.i
# Add `#cgo LDFLAGS: -L/usr/local/lib -lladder` to `server/server.go`
go build
```
