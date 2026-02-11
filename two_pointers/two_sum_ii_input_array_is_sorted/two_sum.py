def two_sum(nums, target):
    l, r = 0, len(nums) - 1

    while l < r:
        l_value = nums[l]
        r_value = nums[r]
        if l_value + r_value == target:
            return [l + 1, r + 1]
        if l_value + r_value > target:
            r -= 1
            continue
        l += 1
    return []
