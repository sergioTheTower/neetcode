# Min Stack

**Difficulty:** Medium  
**Topics:** Stack, Design

## Description

Design a stack class that supports the `push`, `pop`, `top`, and `getMin` operations. Each function must run in **`O(1)`** time.

### Methods

- `MinStack()`: Initializes the stack object.
- `void push(int val)`: Pushes the element `val` onto the stack.
- `void pop()`: Removes the element on the top of the stack.
- `int top()`: Gets the top element of the stack.
- `int getMin()`: Retrieves the minimum element in the stack.

---

## Example 1

**Input:** `["MinStack", "push", 1, "push", 2, "push", 0, "getMin", "pop", "top", "getMin"]`

**Output:** `[null, null, null, null, 0, null, 2, 1]`

**Explanation:**

```java
MinStack minStack = new MinStack();
minStack.push(1);
minStack.push(2);
minStack.push(0);
minStack.getMin(); // return 0
minStack.pop();
minStack.top();    // return 2
minStack.getMin(); // return 1
```
