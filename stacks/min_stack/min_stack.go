package minstack

// MinStack struct is used to construct a MinStack.
type MinStack struct {
	stack []int
	min   int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		min:   0,
	}

}

func (m *MinStack) Push(val int) {
	if len(m.stack) == 0 {
		// We only track differences in the stack value.
		m.stack = append(m.stack, 0)
		m.min = val
		return
	}
	m.stack = append(m.stack, val-m.min)
	// If the new value is less than the current min, update min.
	if val < m.min {
		m.min = val
	}
}

func (m *MinStack) Pop() {
	if len(m.stack) == 0 {
		return
	}
	// Get the last value in the stack.
	pop := m.stack[len(m.stack)-1]
	m.stack = m.stack[:len(m.stack)-1]
	// If the popped value is less than zero.
	if pop < 0 {
		m.min = m.min - pop
	}
}

func (m *MinStack) Top() int {
	top := m.stack[len(m.stack)-1]
	if top > 0 {
		return top + m.min
	}
	return m.min
}

func (m *MinStack) GetMin() int {
	return m.min
}
