def three_sum(nums):
    result = []
    nums.sort()

    for i in range(len(nums)):
        first_value = nums[i]
        if first_value > 0:
            break
        if i > 0 and first_value == nums[i - 1]:
            continue
        left_pointer, right_pointer = i + 1, len(nums) - 1
        while left_pointer < right_pointer:
            left_value = nums[left_pointer]
            right_value = nums[right_pointer]
            total_sum = left_value + right_value + first_value
            if total_sum > 0:
                right_pointer -= 1
                continue
            if total_sum < 0:
                left_pointer += 1
                continue
            result.append([first_value, left_value, right_value])
            right_pointer -= 1
            left_pointer += 1
            while (
                left_pointer < right_pointer
                and nums[left_pointer] == nums[left_pointer - 1]
            ):
                left_pointer += 1

    return result
