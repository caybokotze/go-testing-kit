package testing_kit

type comparator interface {
	to(to bool) bool
	equal(compare bool) Comparison
}

type Comparison struct {
	baseCompare interface{}
}

type RefComparison struct {
	baseCompare interface{}
}

func Compare(arg interface{}) Comparison {
	return Comparison{
		baseCompare: arg,
	}
}

func CompareRef(arg interface{}) RefComparison {
	return RefComparison{
		baseCompare: arg,
	}
}

func (b Comparison) To(arg interface{}) bool {
	if b.baseCompare == arg {
		return true
	}
	return false
}

func (b RefComparison) To(arg interface{}) bool {
	if &b.baseCompare == &arg {
		return true
	}
	return false
}

