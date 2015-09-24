# sflow-go
GoLang tools for Sflow, originally from https://github.com/wimtie/sflow-go

# Building

Inside `sflow-go` directory:

```
$ export GOPATH=$GOPATH:`pwd`
$ go get github.com/influxdb/influxdb/client
$ go build sflux
```

An executable `sflux` should now be built.
