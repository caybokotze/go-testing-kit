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
	var person = person{
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
	var bob1 = person{
		name: "Bob",
		surname: "Williams",
		age: 24,
	}
	var bob2 = person{
		name: "Bob",
		surname: "Williams",
		age: 24,
	}
	// act
	var expected = compare(bob1).to(bob2)
	// assert
	if expected == false {
		t.Log("A type should be equal to the same type")
		t.Fail()
	}
}

func Test_Comparer_StructComparedToDifferentValues_ShouldBeFalse(t *testing.T) {
	var person1 = person{
		name: "Bob",
		surname: "Seebert",
		age: 54,
	}
	var person2 = person{
		name: "John",
		surname: "Mellow",
		age: 39,
	}
	// act
	// arrange
	if compare(person1).to(person2) {
		t.Log("A struct should not be equal to a different struct with different values")
		t.Fail()
	}
}

func Test_Comparer_StructEqualStruct_ShouldNotBeEqual(t *testing.T) {
	// arrange
	var bob1 = person{
		name:    "Bob",
		surname: "Williams",
		age:     24,
	}
	var bob2 = person{
		name: "Bob",
		surname: "Williams",
		age: 24,
	}
	// act
	// assert
	if !compare(bob1).to(bob2) {
		t.Log("Compare should match the same structs")
		t.Fail()
	}
	if compareRef(bob1).to(bob2) {
		t.Log("2 structs should not be equal via compare ref")
		t.Fail()
	}
}

type person struct {
	name string
	surname string
	age int
}