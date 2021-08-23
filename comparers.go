package testing_kit

type comparator interface {
	to(to bool) bool
	equal(compare bool) comparison
}

type comparison struct {
	baseCompare interface{}
}

type refComparison struct {
	baseCompare interface{}
}

func compare(arg interface{}) comparison {
	return comparison{
		baseCompare: arg,
	}
}

func compareRef(arg interface{}) refComparison {
	return refComparison{
		baseCompare: arg,
	}
}

func (b comparison) to(arg interface{}) bool {
	if b.baseCompare == arg {
		return true
	}
	return false
}

func (b refComparison) to(arg interface{}) bool {
	if &b.baseCompare == &arg {
		return true
	}
	return false
}

