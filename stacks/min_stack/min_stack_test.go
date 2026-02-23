package minstack

import (
	"testing"
)

func TestMinStack_Example1(t *testing.T) {
	// Replicating Example 1 from the prompt
	minStack := Constructor()

	minStack.Push(1)
	minStack.Push(2)
	minStack.Push(0)

	if got := minStack.GetMin(); got != 0 {
		t.Errorf("GetMin() = %d; want 0", got)
	}

	minStack.Pop()

	if got := minStack.Top(); got != 2 {
		t.Errorf("Top() = %d; want 2", got)
	}

	if got := minStack.GetMin(); got != 1 {
		t.Errorf("GetMin() = %d; want 1", got)
	}
}

func TestMinStack_NegativeNumbers(t *testing.T) {
	minStack := Constructor()

	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	if got := minStack.GetMin(); got != -3 {
		t.Errorf("GetMin() = %d; want -3", got)
	}

	minStack.Pop()

	if got := minStack.Top(); got != 0 {
		t.Errorf("Top() = %d; want 0", got)
	}

	if got := minStack.GetMin(); got != -2 {
		t.Errorf("GetMin() = %d; want -2", got)
	}
}
