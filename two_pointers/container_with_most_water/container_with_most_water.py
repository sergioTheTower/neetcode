def contains_water(heights):
    result = 0
    left, right = 0, len(heights) - 1
    while left < right:
        # right - left is the difference to determine the width.
        area = (right - left) * min(heights[left], heights[right])
        result = max(result, area)
        if heights[left] > heights[right]:
            right -= 1
            continue
        left += 1
    return result
