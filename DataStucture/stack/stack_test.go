package stack

import "testing"

func Test_newStack_stackIsEmpty(t *testing.T) {
	newStack := Stack{}
	if !newStack.IsEmpty() {
		t.Error("expected ", true, " actual ", newStack.IsEmpty())
	}
}

func Test_newStack_pushElement_stackIsEmpty_false_sizeIs1(t *testing.T) {
	newStack := Stack{}

	newStack.push(1)

	expected := newStack.IsEmpty()

	if newStack.IsEmpty() {
		t.Error("Expected: ", expected, "Actual: ", false)
	}

	if newStack.size() != 1 {
		t.Error("Expected :", 1, "Actaul : ", newStack.size())
	}

}

func Test_newStack_pushFiveElement_sizeIs5_pushAgain_stackOverflowError(t *testing.T) {
	newStack := Stack{}

	newStack.push(10)
	newStack.push(20)
	newStack.push(30)
	newStack.push(40)
	newStack.push(50)

	if newStack.size() != 5 {
		t.Error("Expected :", 5, "Actaul : ", newStack.size())
	}

	result := newStack.push(60)

	expected := "Stack Overflow : stack is full, can't add another element"

	if result.Error() != expected {
		t.Error("Expected Error Message:\n", expected, "Actual : ", result.Error())
	}

}

func Test_newStack_pop_stackIsEmptyError(t *testing.T) {
	newStack := Stack{}

	_, exception := newStack.pop()

	errorMessage := "Stack is empty"
	if exception.Error() != errorMessage {
		t.Error("Expected : ", errorMessage, "\nActual : ", exception.Error())
	}

}

func Test_newStack_push3Element_LIFO_popOnce_sizeIs2(t *testing.T) {
	newStack := Stack{}

	newStack.push(10)
	newStack.push(20)
	newStack.push(30)

	if newStack.size() != 3 {
		t.Error("Expected : 3\nActual : ", newStack.size())
	}

	number, _ := newStack.pop()

	if number != 30 {
		t.Error("Excepted : 30\nActual", number)
	}

	if newStack.size() != 2 {
		t.Error("Expected : 2\nActual : ", newStack.size())
	}

}

func Test_newStack_push3Element_LIFO_peak30_pop_peak20(t *testing.T) {
	newStack := Stack{}

	newStack.push(10)
	newStack.push(20)
	newStack.push(30)

	if newStack.size() != 3 {
		t.Error("Expected : 3\nActual : ", newStack.size())
	}

	number, _ := newStack.pop()

	if number != 30 {
		t.Error("Excepted : 30\nActual", number)
	}

	peaked, _ := newStack.peak()

	if peaked != 20 {
		t.Error("Excepted Peaked Number: 20\nActual : ", peaked)
	}

	number, _ = newStack.pop()

	peaked, _ = newStack.peak()

	expected := 10
	if peaked != expected {
		t.Error("Excepted : ", expected, "\nActual : ", peaked)
	}
}
