# Trapping Rain Water

**Difficulty:** Hard  
**Tags:** Array, Two Pointers, Dynamic Programming, Stack

## Problem Description

You are given an array of non-negative integers `height` which represent an elevation map. Each value `height[i]` represents the height of a bar, which has a width of `1`.

Return the **maximum area** of water that can be trapped between the bars.

---

## Visual Representation

_In the visual above, the black bars represent the elevation `heights`, and the blue areas represent the trapped water._

---

## Examples

### Example 1

**Input:** `height = [0,2,0,3,1,0,1,3,2,1]`  
**Output:** `9`

**Explanation:**
We can calculate the water trapped at each index by finding the minimum of the maximum heights to the left and right, minus the current height.

|   Index   | Height | Max Left | Max Right | Calculation `min(L,R) - h` | Water Trapped |
| :-------: | :----: | :------: | :-------: | :------------------------: | :-----------: |
|     0     |   0    |    0     |     3     |           0 - 0            |       0       |
|     1     |   2    |    2     |     3     |           2 - 2            |       0       |
|     2     |   0    |    2     |     3     |           2 - 0            |     **2**     |
|     3     |   3    |    3     |     3     |           3 - 3            |       0       |
|     4     |   1    |    3     |     3     |           3 - 1            |     **2**     |
|     5     |   0    |    3     |     3     |           3 - 0            |     **3**     |
|     6     |   1    |    3     |     3     |           3 - 1            |     **2**     |
|     7     |   3    |    3     |     3     |           3 - 3            |       0       |
|     8     |   2    |    3     |     2     |           2 - 2            |       0       |
|     9     |   1    |    3     |     1     |           1 - 1            |       0       |
| **Total** |        |          |           |                            |     **9**     |

---

## Constraints

- `1 <= height.length <= 1000`
- `0 <= height[i] <= 1000`

---
