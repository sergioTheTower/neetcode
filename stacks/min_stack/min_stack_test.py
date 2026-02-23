from min_stack import MinStack


def test_min_stack_example_1():
    """Replicates Example 1 from the original problem description."""
    minStack = MinStack()

    minStack.push(1)
    minStack.push(2)
    minStack.push(0)

    assert minStack.getMin() == 0

    minStack.pop()

    assert minStack.top() == 2
    assert minStack.getMin() == 1


def test_min_stack_negative_numbers():
    """Tests the stack's behavior with negative integers."""
    minStack = MinStack()

    minStack.push(-2)
    minStack.push(0)
    minStack.push(-3)

    assert minStack.getMin() == -3

    minStack.pop()

    assert minStack.top() == 0
    assert minStack.getMin() == -2


def test_min_stack_single_element():
    """Tests operations when the stack only has one element."""
    minStack = MinStack()

    minStack.push(42)

    assert minStack.top() == 42
    assert minStack.getMin() == 42


def test_min_stack_duplicate_minimums():
    """Tests pushing and popping duplicate minimum values."""
    minStack = MinStack()

    minStack.push(5)
    minStack.push(5)

    assert minStack.getMin() == 5

    minStack.pop()

    assert minStack.getMin() == 5
    assert minStack.top() == 5
