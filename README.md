# poc-pg-mirror

This is a proof of concept API build upon the equinox postgresql mirror.

Go types and accessor functions are genenrated based on the postgresql tables and associated indexes.  These are mapped to example structs which are then serislised to JSON in an http handler.

Installation:
```
go get github.com/utilitywarehouse/poc-pg-mirror
```

Usage:
```
pog-pg-mirror --help
```

Example call once running:
```
curl localhost:8080/customers/0000001
```

