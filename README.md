# poc-pg-mirror

This is a proof of concept API build upon the equinox postgresql mirror.

Go types and accessor functions are genenrated based on the postgresql tables and associated indexes.  These are mapped to example structs which are then serislised to JSON in an http handler.

Installation:
```
go get github.com/utilitywarehouse/poc-pg-mirror
```
(Note that initial installation is unusually slow for a Go application due to the huge amount of generated code included.)


Usage:
```
pog-pg-mirror --help
```

Example call once running:
```
curl localhost:8080/customers/0000001
```



This code is not production ready, is untested, and should not be used for anything.  It serves only as an example and to illustrate the effort and complexity that might be required for creating such APIs for real in Go.
