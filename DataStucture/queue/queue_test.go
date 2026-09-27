package queue

import "testing"

func Test_newQueue_IsEmpty_returnTrue(t *testing.T) {
	newQueue := Queue{}
	expected := true
	if newQueue.IsEmpty() != expected {
		t.Error("Expected : ", expected, "\nActuall : ", false)
	}
}

func Test_newQueue_size_return0(t *testing.T) {
	newQueue := Queue{}
	expected := 0
	if newQueue.Size() != expected {
		t.Error("Expected : ", expected, "\nActuall : ", newQueue.Size())
	}
}

func Test_newQueue_enqueueOnce_isemptyIsFalse_sizeIsOne(t *testing.T) {
	newQueue := Queue{}

	newQueue.Enqueue(10)
	expected := false
	if newQueue.IsEmpty() != expected {
		t.Error("Expected : ", expected, "\nActuall : ", newQueue.IsEmpty())
	}

	expectedSize := 1
	if newQueue.Size() != expectedSize {
		t.Error("Expected : ", expectedSize, "\nActuall : ", newQueue.Size())
	}
}

func Test_newQueue_enqueueFiveTimes_SizeIs5_enqueue_RaiseError(t *testing.T) {
	newQueue := Queue{}

	newQueue.Enqueue(10)
	newQueue.Enqueue(20)
	newQueue.Enqueue(30)
	newQueue.Enqueue(40)
	newQueue.Enqueue(50)

	expectedSize := 5
	if newQueue.Size() != expectedSize {
		t.Error("Expected : ", expectedSize, "\nActuall : ", newQueue.Size())
	}

	exception := newQueue.Enqueue(50)

	expectedMessage := "Queue Overflow : queue is full"

	if exception.Error() != expectedMessage {
		t.Error("Expected Error Message >> ", expectedMessage, "\nActual Error Message >> ", exception.Error())
	}

}

func Test_newQueue_enqueueThreeTimes_sizeIs3_dequeue_sizeIs2(t *testing.T) {
	newQueue := Queue{}

	newQueue.Enqueue(10)
	newQueue.Enqueue(20)
	newQueue.Enqueue(30)

	expectedSize := 3
	if newQueue.Size() != expectedSize {
		t.Error("Expected : ", expectedSize, "\nActuall : ", newQueue.Size())
	}

	number, _ := newQueue.dequeue()

	expectedNumber := 10

	expectedSize = 2

	if number != expectedNumber {
		t.Error("Expected : ", expectedNumber, "\nActuall : ", number)

	}

}

func Test_newQueue_dequeue_raiseError(t *testing.T) {
	newQueue := Queue{}

	_, actual := newQueue.dequeue()

	expectedMessage := "Queue is empty"
	if actual.Error() != expectedMessage {
		t.Error("Expected : ", expectedMessage, "\nActuall : ", actual.Error())

	}

}

func Test_newQueue_push3Elements_peak_returns10_dequeue_peak_return20(t *testing.T) {
	newQueue := Queue{}

	newQueue.Enqueue(10)
	newQueue.Enqueue(20)
	newQueue.Enqueue(30)

	actual, _ := newQueue.peak()

	expected := 10

	if actual != expected {
		t.Error("Expected >> ", expected, "\nActual : ", actual)
	}

	newQueue.dequeue()

	actual, _ = newQueue.peak()

	expected = 20

	if actual != expected {
		t.Error("Expected >> ", expected, "\nActual : ", actual)
	}

}

func Test_newQueue_push3Elements_sizeIs3_empty_sizeIsZero_peak_returnError(t *testing.T) {
	newQueue := Queue{}

	newQueue.Enqueue(10)
	newQueue.Enqueue(20)
	newQueue.Enqueue(30)

	expected := 3

	if newQueue.Size() != expected {
		t.Error("Expected >> ", expected, "\nActual : ", newQueue.Size())
	}

	newQueue.Empty()

	expected = 0

	if newQueue.Size() != expected {
		t.Error("Expected >> ", expected, "\nActual : ", newQueue.Size())
	}

	_, actual := newQueue.peak()

	expectedMessage := "Queue is empty"

	if actual.Error() != expectedMessage {
		t.Error("Expected >> ", expected, "\nActual : ", actual)
	}

}
