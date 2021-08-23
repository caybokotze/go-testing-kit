package testing_kit

import "testing"

func Test_Comparer_TrueComparedToFalse_ShouldBeFalse(t *testing.T) {
	// arrange
	// act
	var expected = compare(true).to(false)
	// assert
	if expected == true {
		t.Log("True should not be equal to false")
		t.Fail()
	}
}

func Test_Comparer_FalseComparedToTrue_ShouldBeFalse(t *testing.T) {
	// arrange
	// act
	var expected = compare(false).to(true)
	// assert
	if expected == true {
		t.Log("False should not be equal to true")
		t.Fail()
	}
}

func Test_Comparer_TrueComparedToTrue_ShouldBeTrue(t *testing.T) {
	// arrange
	// act
	var expected = compare(true).to(true)
	// assert
	if expected == false {
		t.Log("True should be equal to true")
		t.Fail()
	}
}

func Test_Comparer_FalseComparedToFalse_ShouldBeTrue(t *testing.T) {
	// arrange
	// act
	var expected = compare(false).to(false)
	// assert
	if expected == false {
		t.Log("False should be equal to false")
		t.Fail()
	}
}

func Test_Comparer_StringComparedToBoolean_ShouldBeFalse(t *testing.T) {
	// arrange
	// act
	var expected = compare("random").to(false)
	// assert
	if expected == true {
		t.Log("String should not be equal to boolean")
		t.Fail()
	}
}

func Test_Comparer_StructShouldEqualStruct_ShouldBeTrue(t *testing.T) {
	// arrange
	var person = test {
		name: "Bob",
		surname: "Williams",
		age: 24,
	}
	// act
	var expected = compare(person).to(person)
	// assert
	if expected == false {
		t.Log("A type should be equal to the same type")
		t.Fail()
	}
}


func Test_Comparer_StructShouldNotEqualDifferentStruct_ShouldBeFalse(t *testing.T) {
	// arrange
	var person = test {
		name: "Bob",
		surname: "Williams",
		age: 24,
	}
	var person2 = test {
		name: "Bob",
		surname: "Williams",
		age: 24,
	}
	// act
	var expected = compare(person).to(person2)
	// assert
	if expected == false {
		t.Log("A type should be equal to the same type")
		t.Fail()
	}
}

type test struct {
	name string
	surname string
	age int
}