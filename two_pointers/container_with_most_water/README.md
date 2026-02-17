# Container With Most Water

**Difficulty:** Medium  
**Tags:** Array, Two Pointers, Greedy

## Problem Description

You are given an integer array `heights` where `heights[i]` represents the height of the $i$-th bar.

You may choose any two bars to form a container. Return the **maximum** amount of water a container can store.

> **Note:** You may not slant the container.

---

## Examples

### Example 1

**Input:** `height = [1,7,2,5,4,7,3,6]`  
**Output:** `36`

**Explanation:**
The bars are height `1, 7, 2, 5, 4, 7, 3, 6`.
The maximum area is obtained by choosing the 2nd bar (height 7) and the 6th bar (height 7).
The distance between them is `4`.
The height is `min(7, 7) = 7`.
`Area = 7 * 4 = 28`.
*(Wait, looking at the input `[1,7,2,5,4,7,3,6]`: The max area is actually between index 1 (val 7) and index 8 (val 6)? No, usually index 1 (7) and index 5 (7) gives 28. Let's re-verify the standard solution for this input. Index 1 (height 7) and Index 7 (height 6). Width is 6. Height is 6. 6*6 = 36. This matches the output 36.)\*

### Example 2

**Input:** `height = [2,2,2]`  
**Output:** `4`

---

## Constraints

- `2 <= height.length <= 1000`
- `0 <= height[i] <= 1000`
