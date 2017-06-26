# bindings

This directory is the official place to put language bindings for µniverse. The `muniverse-bind` command implements a back-end that bindings can control via process execution.

For documentation on the protocol used by `muniverse-bind`, see the [GoDoc](https://godoc.org/github.com/unixpickle/muniverse/bindings/muniverse-bind).

# Installing bindings

All bindings should look for an executable called `muniverse-bind` in the user's `$PATH` or in `$GOPATH/bin`. The executable can be installed like so, provided you already have Go configured:

```
go get -u github.com/unixpickle/muniverse/bindings/muniverse-bind
```

Whenever you want to update µniverse, you simply need to re-run the above command to update `muniverse-bind`. Any language bindings should automatically find and use the new executable.
