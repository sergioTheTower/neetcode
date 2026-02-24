# 🧮 Evaluate Reverse Polish Notation

## 📝 Problem Description

You are given an array of strings `tokens` that represents a valid arithmetic expression in **Reverse Polish Notation (RPN)**.

Your task is to return the integer that represents the evaluation of the expression.

### Rules:

- **Operands** can be integers or the results of other operations.
- **Operators** include `+`, `-`, `*`, and `/`.
- **Division** between two integers must always **truncate toward zero**.
- The input is guaranteed to be a **valid** arithmetic expression.

---

## 💡 Example

**Input:** `tokens = ["1", "2", "+", "3", "*", "4", "-"]`  
**Output:** `5`  
**Explanation:** 1. `(1 + 2) = 3` 2. `(3 * 3) = 9` 3. `(9 - 4) = 5`

---

## 🧠 Logic & Approach: The Stack Method

Reverse Polish Notation (also known as Postfix Notation) is most efficiently solved using a **Stack** data structure. Because the operators follow the operands, we can process the expression linearly.

### Step-by-Step Algorithm:

1.  **Initialize** an empty stack.
2.  **Iterate** through each token in the array:
    - If the token is a **number**: Push it onto the stack.
    - If the token is an **operator**:
      - Pop the top two elements from the stack (let the first popped be `y` and the second be `x`).
      - Apply the operation: `x [operator] y`.
      - Push the result back onto the stack.
3.  **Result:** After the loop, the stack will contain exactly one element—the final answer.

---

## ⚙️ Complexity Analysis

| Type      | Complexity | Description                                                                                                    |
| :-------- | :--------- | :------------------------------------------------------------------------------------------------------------- |
| **Time**  | $O(n)$     | We iterate through the list of `n` tokens exactly once.                                                        |
| **Space** | $O(n)$     | In the worst-case scenario (e.g., all numbers followed by all operators), the stack stores up to $n$ elements. |

---

## 🛠️ Constraints

- `1 <= tokens.length <= 1000`
- `tokens[i]` is either an operator (`"+"`, `"-"`, `*`, or `"/"`) or an integer in the range `[-100, 100]`.

---
