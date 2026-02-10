def longestConsecutive(nums):
    numSet = set(nums)

    longest = 0
    for num in numSet:
        length = 1
        while num + length in numSet:
            length += 1
        if length > longest:
            longest = length
    return longest
