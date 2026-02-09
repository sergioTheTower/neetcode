def product_except_self(nums):
    result = [1] * len(nums)
    prefix = 1
    for i in range(len(nums)):
        result[i] = prefix
        prefix *= nums[i]
    postfix = 1
    for i in range(len(nums), 0, -1):
        result[i - 1] *= postfix
        postfix *= nums[i - 1]
    return result
