package stack

import "errors"

type Stack struct {
	elements [5]int
	count    int
}

func (this *Stack) IsEmpty() bool {
	return this.count == 0
}

func (this *Stack) push(number int) error {
	if this.count >= len(this.elements) {
		return errors.New("Stack Overflow : stack is full, can't add another element")
	}
	this.elements[this.count] = number
	this.count++
	return nil
}

func (this *Stack) size() int {
	return this.count
}

func (this *Stack) pop() (int, error) {
	if this.count == 0 {
		return 0, errors.New("Stack is empty")
	}

	topIndex := this.count - 1
	element := this.elements[topIndex]

	this.elements[topIndex] = 0

	this.count--

	return element, nil
}

func (this *Stack) peak() (int, error) {
	topIndex := this.count - 1
	element := this.elements[topIndex]

	return element, nil
}
