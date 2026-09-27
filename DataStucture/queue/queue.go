package queue

import "errors"

type Queue struct {
	elements [5]int
	count    int
}

func (this *Queue) IsEmpty() bool {
	return this.count == 0
}

func (this *Queue) Size() int {
	return this.count
}

func (this *Queue) Enqueue(element int) error {
	if this.count >= len(this.elements) {
		return errors.New("Queue Overflow : queue is full")
	}
	this.elements[this.count] = element
	this.count++

	return nil

}

func (this *Queue) dequeue() (int, error) {

	if this.IsEmpty() == true {
		return 0, errors.New("Queue is empty")
	}
	front := this.elements[0]
	this.elements[0] = 0

	for index := 0; index < len(this.elements)-1; index++ {
		if this.elements[index] == 0 {
			temp := this.elements[index+1]
			this.elements[index+1] = this.elements[index]
			this.elements[index] = temp
		}
	}

	this.count--

	return front, nil

}

func (this *Queue) peak() (int, error) {
	if this.IsEmpty() == true {
		return 0, errors.New("Queue is empty")
	}
	return this.elements[0], nil
}

func (this *Queue) Empty() {
	this.elements = [5]int{}
	this.count = 0

}
