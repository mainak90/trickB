package trickB

type TBool int

const (
	True TBool = 2
	False TBool = 0
	UnSet TBool = 1
)

var values = [3]TBool{False, UnSet, True}

var strings = [3]string{"False", "Unset", "True"}

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

func (a TBool) And(b TBool) TBool {
	return values[min(a, b)]
}

func (a TBool) Or(b TBool) TBool {
	return values[max(a, b)]
}

func (a TBool) Nand(b TBool) TBool {
	return values[2-min(a, b)]
}

func (a TBool) Not() TBool {
	return values[2-a]
}

func (a TBool) Nor(b TBool) TBool {
	return values[2-max(a, b)]
}

func (a TBool) Xor(b TBool) TBool {
	return a.Or(b).And(a.Nand(b))
}

func (a TBool) Out() string {
	return strings[a]
}

func (a TBool) IsSet() bool {
	if a == UnSet {
		return false
	}
	return true
}





