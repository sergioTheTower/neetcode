# Valid Sudoku

**Difficulty:** Medium | **Topics:** Hash Table, Matrix

## Problem Description

Determine if a `9 x 9` Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

1.  Each row must contain the digits `1-9` without repetition.
2.  Each column must contain the digits `1-9` without repetition.
3.  Each of the nine `3 x 3` sub-boxes of the grid must contain the digits `1-9` without repetition.

**Note:**

- A Sudoku board (partially filled) could be valid but is not necessarily solvable.
- Only the filled cells need to be validated according to the mentioned rules.

---

## Visualizing the 3x3 Grid Logic

The trickiest part of this problem is identifying which `3x3` sub-box a cell $(r, c)$ belongs to. You can map every cell to a box index using the formula:
$$\text{box\_index} = (r / 3) \times 3 + (c / 3)$$

[Image of Sudoku 3x3 grid indexing diagram]

---

### Example

**Input:**

```
[["5","3",".",".","7",".",".",".","."],
 ["6",".",".","1","9","5",".",".","."],
 [".","9","8",".",".",".",".","6","."],
 ["8",".",".",".","6",".",".",".","3"],
 ["4",".",".","8",".","3",".",".","1"],
 ["7",".",".",".","2",".",".",".","6"],
 [".","6",".",".",".",".","2","8","."],
 [".",".",".","4","1","9",".",".","5"],
 [".",".",".",".","8",".",".","7","9"]]

```

**Output:** False
