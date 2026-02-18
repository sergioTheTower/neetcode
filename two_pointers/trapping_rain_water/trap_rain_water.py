def trap(heights):
    if not heights:
        return 0

    left_pointer, right_pointer = 0, len(heights) - 1
    left_max, right_max = heights[left_pointer], heights[right_pointer]
    result = 0
    while left_pointer < right_pointer:
        if left_max < right_max:
            left_pointer += 1
            left_max = max(left_max, heights[left_pointer])
            # Do not worry about negative numbers because we set left_max first.
            result += left_max - heights[left_pointer]
        else:
            right_pointer -= 1
            right_max = max(right_max, heights[right_pointer])
            # Do not worry about negative numbers because we set right_max first.
            result += right_max - heights[right_pointer]
    return result
