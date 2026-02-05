def two_sum(nums, target):
    track = dict()
    for i, num in enumerate(nums):
        diff = target - num
        if diff in track:
            return [track[diff], i]
        track[num] = i
    return []
