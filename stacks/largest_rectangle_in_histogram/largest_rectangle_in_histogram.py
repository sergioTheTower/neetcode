def largestRectangleArea(heights):
    maxArea = 0
    stack = []
    for idx, h in enumerate(heights):
        start = idx
        while stack and stack[-1][1] > h:
            index, height = stack.pop()
            maxArea = max(maxArea, height * (idx - index))
            start = index
        stack.append((start, h))

    for i, h in stack:
        maxArea = max(maxArea, h * (len(heights) - i))
    return maxArea
