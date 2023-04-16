# TrickB [![](https://github.com/mainak90/trickB/workflows/build/badge.svg)](https://github.com/mainak90/trickB/actions) [![PkgGoDev](https://pkg.go.dev/badge/github.com/mainak90/trickB)](https://pkg.go.dev/github.com/mainak90/trickB)


`TrickB` is a simple golang library that allows users to set an extra `unset` state for Boolean alike types.

## Implementation
TrickB implements the standard logical operators supported by boolean types.
* `And`
* `Or`
* `Nand`
* `Not`
* `Nor`
* `Xor`

It supports 3 states on a variable `False`, `True` and `UnSet`. 
`UnSet` denotes that a bool like variable has been initiated but not set.
Please keep in mind that the default state for TrickB custom type is `UnSet` to set it, you need to do a `TrickB.True/False` of the variable.

## Utilities
As of now as support for two utility functions.
* `Out` Gives out the string representation of the Boolean state(including unset).
* `IsSet` Just provides a boolean indicator to if the specific flag is set or unset. Note that this does not tell user if flag is set to True/False.
* `TrickBFromString` This one takes a string and returns a trickBoolean equivalent type(`True`, `False`, `Unset`). This can be coupled with `Out` or other logic gates.

## Usage
To get package locally
```
go get github.com/mainak90/trickB
```

Import in code
```
import "github.com/mainak90/trickB"
```

Examples
```
examples/...
```

Live example
```
a := tb.UnSet
b := tb.False
out := a.And(b).Out()
fmt.Println(out)
```

Chain different strings as boolean equivalent types against logic gates and get result
```
fmt.Println(tb.TrickBFromString("true").And(tb.TrickBFromString("unset")).Out())
```

