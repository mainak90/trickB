package trickB

import "github.com/iancoleman/strcase"

type TBool int

// Trick Boolean types --> True/False/unSet.
const (
	True TBool = 2
	False TBool = 0
	UnSet TBool = 1
)

var values = [3]TBool{False, UnSet, True}

var strings = [3]string{"False", "UnSet", "True"}

func max(a, b TBool) TBool {
	if a > b {
		return a
	}
	return b
}

func min(a, b TBool) TBool {
	if a < b {
		return a
	}
	return b
}

/*
Implements logical And.
	False False | False
	False UnSet | False
	True  UnSet | UnSet
 */
func (a TBool) And(b TBool) TBool {
	return values[min(a, b)]
}

/*
Implements logical Or
 */
func (a TBool) Or(b TBool) TBool {
	return values[max(a, b)]
}

// Implements logical Nand.
func (a TBool) Nand(b TBool) TBool {
	return values[2-min(a, b)]
}

// Implements logical Not.
func (a TBool) Not() TBool {
	return values[2-a]
}

// Implements logical Nor.
func (a TBool) Nor(b TBool) TBool {
	return values[2-max(a, b)]
}

// Implements logical Xor.
func (a TBool) Xor(b TBool) TBool {
	return a.Or(b).And(a.Nand(b))
}

// Out outputs a TrickBool state as a string representation.
func (a TBool) Out() string {
	return strings[a]
}

// Checks if a Tbool instance isSet or Unset. Returns True if set, else Unset.
func (a TBool) IsSet() bool {
	if a == UnSet {
		return false
	}
	return true
}

/*
Outputs a TrickB state type representation from a basic string input.
Outputs Unset if string does not match bool/unset.
 */
func TrickBFromString(s string) TBool {
	switch strcase.ToCamel(s) {
	case "UnSet":
		return UnSet
	case "Unset":
		return UnSet
	case "True":
		return True
	case "TRUE":
		return True
	case "False":
		return False
	case "FALSE":
		return False
	default:
		return UnSet
	}
}





